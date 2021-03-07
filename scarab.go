package scarab

import (
	"log"
	"time"

	pb "github.com/tcolgate/scarab/pb"
)

// Server implments the leaders service.
type Server struct {
	pb.UnimplementedManagerServer
	pb.UnimplementedWorkerServer

	done chan struct{}
}

func (s *Server) RegisterProfile(req *pb.RegisterProfileRequest, stream pb.Manager_RegisterProfileServer) error {
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

func (s *Server) ReportLoad(req *pb.ReportLoadRequest, stream pb.Worker_ReportLoadServer) error {
	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	log.Printf("start rreporting load")

	for {
		select {
		case <-ctx.Done():
		case <-s.done:
		case <-ticker.C:
			stream.Send(&pb.LoadMetrics{})
		}
	}
	return nil
}

func (s *Server) RunJob(stream pb.Worker_RunJobServer) error {
	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
			case <-s.done:
			case <-ticker.C:
				stream.Send(&pb.JobMetrics{})
			}
		}
	}()
	for {
		stream.Recv()
	}

	return nil
}
