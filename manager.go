package scarab

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/tcolgate/scarab/pb"
	grpc "google.golang.org/grpc"
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

	ctx := stream.Context()

	log.Printf("registered worker for profile %s", req.Spec)

	// run keepalives
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
			case <-s.done:
			case <-ticker.C:
				stream.Send(&pb.KeepAlive{})
			}
		}
	}()

	clientOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	addr := req.Worker.Addr
	srvconn, err := grpc.Dial(addr, clientOpts...)
	if err != nil {
		return err
	}
	defer srvconn.Close()
	wcli := pb.NewWorkerClient(srvconn)
	loadrep, err := wcli.ReportLoad(ctx, &pb.ReportLoadRequest{})

	// consume load reports
	go func() {
		// throw away the keepalives.
		for {
			msg, err := loadrep.Recv()
			if err != nil {
				return
			}
			for range msg.Metrics {
			}
		}
	}()

	return nil
}

// RunProfile implements the UI StartJob method.
func (m *Manager) RunProfile(ctx context.Context, j *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	clientOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	wrkAddrs := m.profiles.GetWorkers(j.Profile, j.Version)
	if len(wrkAddrs) == 0 {
		return nil, errors.New("no worker registered for requested profile")
	}

	addr := wrkAddrs[0].Addr // needs to come
	srvconn, err := grpc.Dial(addr, clientOpts...)
	if err != nil {
		return nil, err
	}
	defer srvconn.Close()
	wcli := pb.NewWorkerClient(srvconn)
	rjsrc, err := wcli.RunJob(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		// throw away the keepalives.
		for {
			_, err := rjsrc.Recv()
			if err != nil {
				log.Printf("error getting loadreport, %#v", err)
				return
			}
		}
	}()

	err = rjsrc.Send(&pb.RunJobRequest{
		Profile: "myprofile",
		Args:    []*pb.JobArg{},
	})
	if err != nil {
		log.Panicf("runjub error, %v", err)
		return nil, err
	}

	return nil, nil
}

// StartJob implements the UI StartJob method.
func (s *Manager) StartJob(ctx context.Context, j *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	log.Printf("Start Job called")
	return s.RunProfile(ctx, j)
}

func (s *Manager) ListProfiles(context.Context, *pb.ListProfilesRequest) (*pb.ListProfilesResponse, error) {
	log.Printf("List Profiles called")
	return &pb.ListProfilesResponse{}, nil
}

func (s *Manager) WatchActiveJobs(*pb.WatchActiveJobsRequest, pb.ManagerUI_WatchActiveJobsServer) error {
	log.Printf("Watch active jobs")

	return nil
}
