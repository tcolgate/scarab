package scarab

import (
	"context"
	"log"
	"os"
	"time"

	model "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	pb "github.com/tcolgate/scarab/pb"
	grpc "google.golang.org/grpc"
)

// Manager implments the leaders service.
type Manager struct {
	pb.UnimplementedManagerServer
	pb.UnimplementedManagerUIServer

	profiles *ProfileRegistry

	done chan struct{}
}

func NewManager() *Manager {
	mng := &Manager{
		profiles: NewProfileRegistr(),
		done:     make(chan struct{}),
	}
	return mng
}

func (s *Manager) RegisterProfile(req *pb.RegisterProfileRequest, stream pb.Manager_RegisterProfileServer) error {
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
				return
			case <-s.done:
				return
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
	if err != nil {
		return err
	}

	for {
		_, err := loadrep.Recv()
		if err != nil {
			return err
		}
	}
}

// RunProfile implements the manager's RunProfile method
func (m *Manager) RunProfile(ctx context.Context, j *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	clientOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	sub := m.profiles.Subscribe(j.Profile, j.Version)
	//args := sub.Args
	for sub.Update(ctx) {
		ws := sub.ActiveWorkers()
		if len(ws) == 0 {
			log.Printf("no active workers")
			continue
		}

		addr := ws[0].Addr // needs to come
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
			workerAddrLabel := "worker_addr"
			for {
				ms, err := rjsrc.Recv()
				if err != nil {
					log.Printf("error getting loadreport, %#v", err)
					return
				}
				log.Printf("got metrics: %#v", ms)
				for _, mf := range ms.Metrics {
					for i := range mf.Metric {
						mf.Metric[i].Label = append(mf.Metric[i].Label, &model.LabelPair{Name: &workerAddrLabel, Value: &addr})
					}
					expfmt.MetricFamilyToText(os.Stdout, mf)
				}
			}
		}()
		err = rjsrc.Send(&pb.RunJobRequest{
			Profile: j.Profile,
			Args:    []*pb.JobArg{},
			Users:   j.Users,
			MaxRate: j.MaxRate,
		})
		if err != nil {
			log.Printf("runjub error, %v", err)
			return nil, err
		}
	}

	return &pb.StartJobResponse{}, nil
}

// StartJob implements the UI StartJob method.
func (s *Manager) StartJob(ctx context.Context, j *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	log.Printf("Start Job called")
	return s.RunProfile(ctx, j)
}

func (s *Manager) StopJob(ctx context.Context, j *pb.StopJobRequest) (*pb.StopJobResponse, error) {
	log.Printf("Stop Job called")
	return &pb.StopJobResponse{}, nil
}

func (s *Manager) ListProfiles(ctx context.Context, req *pb.ListProfilesRequest) (*pb.ListProfilesResponse, error) {
	log.Printf("List Profiles called")
	return s.profiles.ListProfiles(ctx, req)
}

func (s *Manager) WatchActiveJobs(*pb.WatchActiveJobsRequest, pb.ManagerUI_WatchActiveJobsServer) error {
	log.Printf("Watch active jobs")

	return nil
}
