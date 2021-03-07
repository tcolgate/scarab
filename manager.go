package scarab

import (
	"log"
	"sync"
	"time"

	pb "github.com/tcolgate/scarab/pb"
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

	profiles ProfileRegistry

	done chan struct{}
}

func (s *Manager) RegisterProfile(req *pb.RegisterProfileRequest, stream pb.Manager_RegisterProfileServer) error {
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
	return nil
}
