package scarab

import (
	"context"
	"log"
	"time"

	pb "github.com/tcolgate/scarab/pb"
)

// Manager implments the leaders service.
type Manager struct {
	pb.UnimplementedManagerServer
	pb.UnimplementedManagerUIServer

	profiles ProfileRegistry

	done chan struct{}
}

func NewManager() *Manager {
	return &Manager{
		profiles: ProfileRegistry{
			now: time.Now,
		},
		done: make(chan struct{}),
	}
}

func (s *Manager) RegisterProfile(req *pb.RegisterProfileRequest, stream pb.Manager_RegisterProfileServer) error {
	log.Printf("RegisterProfile")
	reg, err := s.profiles.Register(req)
	if err != nil {
		return err
	}
	defer s.profiles.Unregister(reg)

	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	log.Printf("registered worker for profile %s", req.Spec)

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

func (*Manager) StartJob(context.Context, *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	log.Printf("Start Job called")
	return &pb.StartJobResponse{}, nil
}

func (*Manager) ListProfiles(context.Context, *pb.ListProfilesRequest) (*pb.ListProfilesResponse, error) {
	log.Printf("List Profiles called")
	return &pb.ListProfilesResponse{}, nil
}

func (*Manager) WatchActiveJobs(*pb.WatchActiveJobsRequest, pb.ManagerUI_WatchActiveJobsServer) error {
	log.Printf("Watch active jobs")
	return nil
}
