package scarab

import (
	"context"
	"log"
	"os"
	"time"

	prom "github.com/prometheus/client_golang/prometheus"
	pb "github.com/tcolgate/scarab/pb"
	grpc "google.golang.org/grpc"
)

type Worker struct {
	pb.UnimplementedWorkerServer

	reg *prom.Registry

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

func NewWorker(ctx context.Context, addr, serverAddr string, wrk Workload) (*Worker, error) {
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
	}, nil
}

func (s *Worker) Close() {
	s.conn.Close()
}

func (s *Worker) ReportLoad(req *pb.ReportLoadRequest, stream pb.Worker_ReportLoadServer) error {
	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	log.Printf("start reporting load")

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

func (s *Worker) RunJob(stream pb.Worker_RunJobServer) error {
	ticker := time.NewTicker(1 * time.Second)
	ctx := stream.Context()
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-s.done:
				return
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
