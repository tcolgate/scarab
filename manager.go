package scarab

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/tcolgate/scarab/pb"
	"github.com/tcolgate/scarab/ui"
)

type ProfileKey struct {
	Name    string
	Version string
}

type Profile struct {
	WorkerAddrs []string
}

type ProfileRegistry struct {
	sync.RWMutex
	entries map[ProfileKey]Profile
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
