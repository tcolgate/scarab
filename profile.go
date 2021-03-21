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

type profUnSubReq struct {
	id string
}

type profSubReq struct {
	*pb.RegisterProfileRequest
	resp chan profRegResp
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
	unsubs chan profUnSubReq

	list chan proflistReq
}

type subscription struct {
	id string
	ch []*pb.WorkerDetails
}

func (pr *ProfileRegistry) loop(ctx context.Context) {
	entries := map[ProfileKey]Profile{}
	subs := map[ProfileKey][]subscription{}

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
				prof.FirstRegistration = t
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

		case unreq := <-pr.unreg:
			key := unreq.ProfileKey
			prof := entries[key]
			if prof.Workers == nil {
				return
			}
			delete(prof.Workers, unreq)
			entries[key] = prof

		case req := <-pr.subs:
		/*
			key := ProfileKey{
				Profile: profile,
				Version: version,
			}
		*/

		case req := <-pr.unsubs:

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

// GetProfile should return a subscription to a list of worker addresses
func (pr *ProfileRegistry) GetProfile(profile, version string) ([]*pb.ProfileArg, []*pb.WorkerDetails) {
	var dets []*pb.WorkerDetails
	for w := range prof.Workers {
		dets = append(dets, &pb.WorkerDetails{
			Name: w.Name,
			Addr: w.Addr,
		})
	}
	return prof.Args, dets
}
