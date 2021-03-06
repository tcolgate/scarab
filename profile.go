package scarab

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
	"google.golang.org/protobuf/types/known/timestamppb"

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
	*pb.WorkerDetails
	ID string
}

type Profile struct {
	FirstRegistration time.Time
	Spec              *pb.ProfileSpec
	Workers           map[ProfileRegistration]struct{}
}

func (p *Profile) GetActiveWorkers() []*pb.WorkerDetails {
	ws := make([]*pb.WorkerDetails, 0, len(p.Workers))
	for w := range p.Workers {
		ws = append(ws, w.ProfileWorker.WorkerDetails)
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

type proflistReq struct {
	req  *pb.ListProfilesRequest
	resp chan *pb.ListProfilesResponse
}

type ProfileRegistry struct {
	now func() time.Time

	reg   chan profReg
	unreg chan ProfileRegistration

	subs   chan profSubReq
	unsubs chan ProfileSubscription

	list chan *proflistReq
}

func NewProfileRegistr() *ProfileRegistry {
	pr := &ProfileRegistry{
		now:   time.Now,
		reg:   make(chan profReg),
		unreg: make(chan ProfileRegistration),

		subs:   make(chan profSubReq),
		unsubs: make(chan ProfileSubscription),

		list: make(chan *proflistReq),
	}

	go pr.loop(context.TODO())
	return pr
}

func (pr *ProfileRegistry) loop(ctx context.Context) {
	entries := map[ProfileKey]Profile{}
	subs := map[ProfileKey][]ProfileSubscription{}

	t := pr.now()
	r := rand.New(rand.NewSource(t.UnixNano()))
	entropy := ulid.Monotonic(r, 0)

	for {
		log.Printf("profile loop")
		log.Printf("profile loop entries: %#v", entries)
		log.Printf("profile loop subs: %#v", subs)
		select {
		case <-ctx.Done():
		case reg := <-pr.reg:
			log.Printf("loop reg")
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
				WorkerDetails: reg.Worker,
				ID:            id,
			}

			preg := ProfileRegistration{
				ProfileKey:    key,
				ProfileWorker: pw,
			}

			if prof.Workers == nil {
				prof.Spec = reg.Spec
				prof.Workers = make(map[ProfileRegistration]struct{})
			}
			prof.Workers[preg] = struct{}{}
			entries[key] = prof
			reg.resp <- profRegResp{
				err: nil,
				reg: preg,
			}

			active := prof.GetActiveWorkers()
			log.Printf("active ws %#v", active)
			for _, s := range subs[key] {
				s.workers <- active
			}

		case unreq := <-pr.unreg:
			log.Printf("loop unreg")
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
			log.Printf("loop subs")
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
			log.Printf("loop unsubs")
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

		case req := <-pr.list:
			log.Printf("loop list")
			resp := &pb.ListProfilesResponse{}
			for _, e := range entries {
				rprof := &pb.RegisteredProfile{}
				rprof.Spec = e.Spec
				rprof.Workers = e.GetActiveWorkers()
				rprof.FirstRegistration = timestamppb.New(e.FirstRegistration)
				resp.Registered = append(resp.Registered, rprof)
			}
			req.resp <- resp
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

type ProfileSubscription struct {
	Args    []*pb.ProfileArg
	key     ProfileKey
	id      string
	workers chan []*pb.WorkerDetails

	lastWorkers []*pb.WorkerDetails
}

// Subscrube subscribes to updates for the list of workers associated with a
// profile.
func (pr *ProfileRegistry) Subscribe(profile, version string) ProfileSubscription {
	ch := make(chan ProfileSubscription)

	pr.subs <- profSubReq{
		profile: profile,
		version: version,
		resp:    ch,
	}

	resp := <-ch

	return resp
}

// Unsubscribe unsubscribes a given subscription, shut down the worker
// updates
func (pr *ProfileRegistry) Unsubscribe(s ProfileSubscription) {
	pr.unsubs <- s
}

func (pr *ProfileRegistry) ListProfiles(ctx context.Context, req *pb.ListProfilesRequest) (*pb.ListProfilesResponse, error) {
	log.Printf("internal List Profiles called")
	ch := make(chan *pb.ListProfilesResponse)
	pr.list <- &proflistReq{resp: ch, req: req}
	log.Printf("internal List message sent")
	resp := <-ch
	log.Printf("list profiles resp: %#v", resp)
	return resp, nil
}

func (s *ProfileSubscription) Update(ctx context.Context) bool {
	select {
	case ws, ok := <-s.workers:
		if !ok {
			log.Printf("update workers not ok")
			return false
		}
		log.Printf("update workers %#v", ws)
		s.lastWorkers = ws
		return true
	case <-ctx.Done():
		return false
	}
}

func (s *ProfileSubscription) ActiveWorkers() []*pb.WorkerDetails {
	return s.lastWorkers
}
