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

	worker := pb.NewManagerClient(conn)

	hn, _ := os.Hostname()
	req := pb.RegisterProfileRequest{
		Spec: wrk.Spec,
		Worker: &pb.WorkerDetails{
			Name: hn,
			Addr: hn + addr,
		},
	}
	cli, err := worker.RegisterProfile(ctx, &req)
	if err != nil {
		return nil, err
	}

	go func() {
		// throw away the keepalives.
		for {
			_, err := cli.Recv()
			if err != nil {
				log.Printf("error getting keepalive, %#v", err)
				return
			}
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

	log.Printf("start rreporting load")

	for {
		select {
		case <-ctx.Done():
		case <-s.done:
		case <-ticker.C:
			ms, err := s.reg.Gather()
			if err != nil {
				continue
			}
			stream.Send(&pb.LoadMetrics{Metrics: ms})
		}
	}
	return nil
}

func (s *Worker) RunJob(stream pb.Worker_RunJobServer) error {
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
