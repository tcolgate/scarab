package scarab

import (
	"context"
	"math/rand"
	"time"

	"github.com/oklog/ulid"

	pb "github.com/tcolgate/scarab/pb"
)

type ProfileRegistration struct {
	ProfileKey
	ProfileWorker
}

type ProfileKey struct {
	Profile string
	Version string
}

type ProfileWorker struct {
	Addr string
	ID   string
	Name string
}

type Profile struct {
	FirstRegistration time.Time
	Args              []*pb.ProfileArg
	Workers           map[ProfileRegistration]struct{}
}

func (p *Profile) GetActiveWorkers() []*pb.WorkerDetails {
	var ws []*pb.WorkerDetails
	for w := range p.Workers {
		ws = append(ws, nil)
	}
	return ws
}

type profRegResp struct {
	err error
	reg ProfileRegistration
}

type profReg struct {
	*pb.RegisterProfileRequest
	resp chan profRegResp
}

type profSubResp struct {
	id string
	ch chan []*pb.WorkerDetails
}

type profSubReq struct {
	profile string
	version string

	resp chan ProfileSubscription
}

type proflistResp struct {
}

type proflistReq struct {
	resp chan proflistResp
}

type ProfileRegistry struct {
	now func() time.Time

	reg   chan profReg
	unreg chan ProfileRegistration

	subs   chan profSubReq
	unsubs chan ProfileSubscription

	list chan proflistReq
}

func (pr *ProfileRegistry) loop(ctx context.Context) {
	entries := map[ProfileKey]Profile{}
	subs := map[ProfileKey][]ProfileSubscription{}

	t := pr.now()
	r := rand.New(rand.NewSource(t.UnixNano()))
	entropy := ulid.Monotonic(r, 0)

	for {
		select {
		case <-ctx.Done():
		case reg := <-pr.reg:
			key := ProfileKey{
				Profile: reg.Spec.Profile,
				Version: reg.Spec.Version,
			}

			rt := pr.now()
			prof, ok := entries[key]
			if !ok {
				prof.FirstRegistration = rt
			}

			id := ulid.MustNew(ulid.Timestamp(rt), entropy).String()

			pw := ProfileWorker{
				Addr: reg.Worker.Addr,
				Name: reg.Worker.Name,
				ID:   id,
			}

			preg := ProfileRegistration{
				ProfileKey:    key,
				ProfileWorker: pw,
			}

			if prof.Workers == nil {
				prof.Workers = make(map[ProfileRegistration]struct{})
			}
			prof.Workers[preg] = struct{}{}
			entries[key] = prof
			reg.resp <- profRegResp{
				err: nil,
				reg: preg,
			}

			active := prof.GetActiveWorkers()
			for _, s := range subs[key] {
				s.workers <- active
			}

		case unreq := <-pr.unreg:
			key := unreq.ProfileKey
			prof := entries[key]
			if prof.Workers == nil {
				return
			}
			delete(prof.Workers, unreq)
			entries[key] = prof
			active := prof.GetActiveWorkers()
			for _, s := range subs[key] {
				s.workers <- active
			}

		case req := <-pr.subs:
			key := ProfileKey{
				Profile: req.profile,
				Version: req.version,
			}

			ch := make(chan []*pb.WorkerDetails)
			rt := pr.now()
			id := ulid.MustNew(ulid.Timestamp(rt), entropy).String()
			resp := ProfileSubscription{
				key:     key,
				id:      id,
				workers: ch,
			}
			subs[key] = append(subs[key], resp)
			req.resp <- resp

			p := entries[key]
			resp.workers <- p.GetActiveWorkers()

		case req := <-pr.unsubs:
			key := req.key
			psubs := subs[req.key]
			n := 0
			for i, s := range psubs {
				if s.id == req.id {
					close(s.workers)
					continue
				}
				psubs[i] = s
				n++
			}
			psubs = psubs[:n]
			subs[key] = psubs
		}
	}
}

func (pr *ProfileRegistry) Register(req *pb.RegisterProfileRequest) (ProfileRegistration, error) {
	ch := make(chan profRegResp)
	defer close(ch)
	pr.reg <- profReg{
		RegisterProfileRequest: req,
		resp:                   ch,
	}
	resp := <-ch

	return resp.reg, resp.err
}

func (pr *ProfileRegistry) Unregister(preg ProfileRegistration) {
	pr.unreg <- preg
}

func (pr *ProfileRegistry) ListProfiles() {
	ch := make(chan proflistResp)
	pr.list <- proflistReq{resp: ch}
	resp := <-ch
}

type ProfileSubscription struct {
	Args    []*pb.ProfileArg
	key     ProfileKey
	id      string
	workers chan []*pb.WorkerDetails
}

func (*ProfileSubscription) Close() {
}

// Subscrube subscribes to updates for the list of workers associated with a
// profile.
func (pr *ProfileRegistry) Subscribe(profile, version string) ProfileSubscription {
	ch := make(chan ProfileSubscription)

	req := profSubReq{
		resp: ch,
	}

	resp := <-ch

	return resp
}

// Unsubscribe unsubscribes a given subscription, shut down the worker
// updates
func (pr *ProfileRegistry) Unsubscribe(s ProfileSubscription) {
	pr.unsubs <- s
}
