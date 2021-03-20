package scarab

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid"
	pb "github.com/tcolgate/scarab/pb"
	"github.com/tcolgate/scarab/ui"
)

type ProfileRegistration struct {
	ProfileKey
	ProfileWorker
}

type ProfileKey struct {
	Name    string
	Version string
}

type ProfileWorker struct {
	Addr string
	ID   string
	Name string
}
type Profile struct {
	FirstRegistration time.Time
	Workers           map[ProfileRegistration]struct{}
}

type ProfileRegistry struct {
	sync.RWMutex
	now     func() time.Time
	entries map[ProfileKey]Profile
}

func (pr *ProfileRegistry) Register(name, version string, workerName, addr string) (ProfileRegistration, error) {
	pr.Lock()
	defer pr.Unlock()
	if pr.entries == nil {
		pr.entries = make(map[ProfileKey]Profile)
	}
	key := ProfileKey{Name: name, Version: version}
	// TODO check the the profile request is compatible/identical
	// with others registered for the same name/version

	t := pr.now()
	r := rand.New(rand.NewSource(t.UnixNano()))
	entropy := ulid.Monotonic(r, 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy).String()
	prof := pr.entries[key]

	pw := ProfileWorker{
		Addr: addr,
		ID:   id,
		Name: workerName,
	}

	preg := ProfileRegistration{
		ProfileKey:    key,
		ProfileWorker: pw,
	}

	if prof.Workers == nil {
		prof.Workers = make(map[ProfileRegistration]struct{})
	}
	prof.Workers[preg] = struct{}{}
	pr.entries[key] = prof

	return preg, nil
}

func (pr *ProfileRegistry) Unregister(preg ProfileRegistration) {
	pr.Lock()
	defer pr.Unlock()

	if pr.entries == nil {
		return
	}

	key := preg.ProfileKey

	prof := pr.entries[key]
	if prof.Workers == nil {
		return
	}
	delete(prof.Workers, preg)
	pr.entries[key] = prof
}

func (pr *ProfileRegistry) ListProfiles() {
	pr.RLock()
	defer pr.RUnlock()
}

// Manager implments the leaders service.
type Manager struct {
	pb.UnimplementedManagerServer
	ui.UnimplementedManagerUIServer

	profiles ProfileRegistry

	done chan struct{}
}

func (s *Manager) RegisterProfile(req *pb.RegisterProfileRequest, stream pb.Manager_RegisterProfileServer) error {
	log.Printf("RegisterProfile")
	reg, err := s.profiles.Register(req.Profile, req.Version, req.Name, req.Addr)
	if err != nil {
		return err
	}
	defer s.profiles.Unregister(reg)

	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	log.Printf("registered worker for profile %s", req.Profile)

	for {
		select {
		case <-ctx.Done():
		case <-s.done:
		case <-ticker.C:
			stream.Send(&pb.KeepAlive{})
		}
	}
	return nil
}

func (s *Manager) RunProfile(name, version string) error {
	log.Printf("RunProfile")
	return nil
}

func (*Manager) StartJob(context.Context, *ui.StartJobRequest) (*ui.StartJobResponse, error) {
	log.Printf("Start Job called")
	return &ui.StartJobResponse{}, nil
}

func (*Manager) ListProfiles(context.Context, *ui.ListProfilesRequest) (*ui.ListProfilesResponse, error) {
	log.Printf("List Profiles called")
	return &ui.ListProfilesResponse{}, nil
}

func (*Manager) WatchActiveJobs(*ui.WatchActiveJobsRequest, ui.ManagerUI_WatchActiveJobsServer) error {
	log.Printf("Watch active jobs")
	return nil
}
