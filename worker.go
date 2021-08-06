package scarab

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	prom "github.com/prometheus/client_golang/prometheus"
	pb "github.com/tcolgate/scarab/pb"
	grpc "google.golang.org/grpc"
)

type Runner interface {
	Run(ctx context.Context, args []*pb.JobArg)
}

type RunnerFunc func(ctx context.Context, args []*pb.JobArg)

func (f RunnerFunc) Run(ctx context.Context, args []*pb.JobArg) {
	f(ctx, args)
}

type User interface {
	Setup(reg prom.Registerer) Runner
}

type UserFunc func(reg prom.Registerer) Runner

func (u UserFunc) Setup(reg prom.Registerer) Runner {
	return u(reg)
}

type Worker struct {
	pb.UnimplementedWorkerServer

	// This is the registry for the entire worker.
	// Mostly only useful for the process metrics.
	reg *prom.Registry

	// user TODO(tcm) crap name
	user User

	conn *grpc.ClientConn
	done chan struct{}
}

func registerOnce(ctx context.Context, addr string, conn *grpc.ClientConn, wrk Workload) error {
	mcli := pb.NewManagerClient(conn)
	hn, _ := os.Hostname()
	req := pb.RegisterProfileRequest{
		Spec: wrk.Spec,
		Worker: &pb.WorkerDetails{
			Name: hn,
			Addr: hn + addr,
		},
	}
	cli, err := mcli.RegisterProfile(ctx, &req)
	if err != nil {
		return err
	}

	// throw away the keepalives.
	for {
		_, err := cli.Recv()
		if err != nil {
			return err
		}
	}
}

func NewWorker(ctx context.Context, addr, serverAddr string, wrk Workload, user User) (*Worker, error) {
	reg := prom.NewRegistry()
	pc := prom.NewProcessCollector(prom.ProcessCollectorOpts{})
	gocol := prom.NewGoCollector()

	err := reg.Register(pc)
	if err != nil {
		return nil, err
	}
	err = reg.Register(gocol)
	if err != nil {
		return nil, err
	}

	clientOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(serverAddr, clientOpts...)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			if err := registerOnce(ctx, addr, conn, wrk); err != nil {
				log.Printf("worker registration failed, %v", err)
				time.Sleep(5 * time.Second)
				continue
			}
			return
		}
	}()

	return &Worker{
		reg:  reg,
		conn: conn,
		user: user,
	}, nil
}

func (s *Worker) Close() {
	s.conn.Close()
}

func (s *Worker) ReportLoad(req *pb.ReportLoadRequest, stream pb.Worker_ReportLoadServer) error {
	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-s.done:
			return nil
		case <-ticker.C:
			ms, err := s.reg.Gather()
			if err != nil {
				continue
			}
			stream.Send(&pb.LoadMetrics{Metrics: ms})
		}
	}
}

func (s *Worker) runJob(req *pb.RunJobRequest) error {
	// setup metrics

	return nil
}

func (s *Worker) RunJob(stream pb.Worker_RunJobServer) error {
	// get the worker details
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	args := req.GetArgs()

	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	jreg := prom.NewRegistry()
	u := s.user.Setup(jreg)

	send := func() {
		ms, err := jreg.Gather()
		if err != nil {
			log.Printf("gather failed, %v", err)
			return
		}

		stream.Send(&pb.JobMetrics{Metrics: ms})
	}

	// TODO(tcm): can't just use contexts here for stopping
	// We should have a proper coordinated and confirmed
	// stop to avoid missing/cancelling any in flight requests
	// (doing so would skew results).
	activeChan := make(chan uint64)
	go func() {
		var cancels []func()
	loop:
		for {
			select {
			case <-ctx.Done():
				// not sure here, can we still send metrics?
				break loop
			case <-s.done:
				break loop
			case <-ticker.C:
				send()
			case setActive := <-activeChan:
				if setActive == 0 {
					break loop
				} else if uint64(len(cancels)) == setActive {
					// nothing to do
				} else if uint64(len(cancels)) > setActive {
					extra := cancels[setActive:]
					cancels = cancels[:setActive]
					for _, c := range extra {
						c()
					}
				} else {
					needed := setActive - uint64(len(cancels))
					us := make([]func(), needed)
					for i := range us {
						uctx, cancel := context.WithCancel(ctx)
						us[i] = cancel
						go func(ctx context.Context) {
							for {
								select {
								case <-ctx.Done():
									return
								default:
									// if the func quits early, we'll just re run it
									u.Run(ctx, args)
								}
							}
						}(uctx)
					}
					cancels = append(cancels, us...)
				}
			}
		}

		send()
		for _, c := range cancels {
			c()
		}
	}()

	activeChan <- req.Users

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			send()
			break
		}
		err = s.runJob(req)
		if err != nil {
			log.Printf("error running job, %v", err)
		}
	}

	return nil
}

func RunStandaloneWorker(profile, ver, desc string, args []*pb.ProfileArg, user User) error {
	ctx := context.Background()
	addr := flag.String("addr", ":8083", "address to listen on")
	serverAddr := flag.String("manager.addr", "localhost:8081", "address of the manager")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	wrk := Workload{
		Spec: &pb.ProfileSpec{
			Profile:     profile,
			Version:     ver,
			Description: desc,
			Args:        args,
		},
	}

	wrkr, err := NewWorker(ctx, *addr, *serverAddr, wrk, user)
	if err != nil {
		return fmt.Errorf("failed to create worker, %w", err)
	}
	defer wrkr.Close()
	pb.RegisterWorkerServer(grpcServer, wrkr)

	err = grpcServer.Serve(lis)
	if err != nil {
		return fmt.Errorf("worker listen err, %w", err)
	}
	return nil
}
