package scarab

import (
	"log"
	"time"

	pb "github.com/tcolgate/scarab/pb"
)

// Server implments the leaders service.
type Server struct {
	done chan struct{}
}

func (s *Server) Register(req *pb.RegisterRequest, stream pb.Leader_RegisterServer) error {
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

func (s *Server) RunJob(req *pb.RunJobRequest, stream pb.Worker_RunJobServer) error {
	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	log.Printf("start running profile %s", req.Profile)

	for {
		select {
		case <-ctx.Done():
		case <-s.done:
		case <-ticker.C:
			stream.Send(&pb.JobMetrics{})
		}
	}
	return nil
}
