package scarab

import (
	"log"
	"time"

	prom "github.com/prometheus/client_golang/prometheus"
	promdto "github.com/prometheus/client_model/go"
	pb "github.com/tcolgate/scarab/pb"
)

type Worker struct {
	pb.UnimplementedWorkerServer

	reg *prom.Registry

	done chan struct{}
}

func promToProto(in []*promdto.MetricFamily) []*pb.MetricFamily {
	res := make([]*pb.MetricFamily, len(in))
	for i := range in {
		ms := make([]*pb.Metric, len(in[i].Metric))
		for j := range in[i].Metric {
			m := &pb.Metric{}
			ls := make(map[string]string)
			for _, l := range in[i].Metric[j].Label {
				ls[*l.Name] = *l.Value
			}
			m.Labels = ls
			ms[j] = m
		}
		res[i] = &pb.MetricFamily{
			Name:    *in[i].Name,
			Help:    *in[i].Help,
			Metrics: ms,
		}
	}
	return res
}

func NewWorker() (*Worker, error) {
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

	return &Worker{
		reg: reg,
	}, nil
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
			stream.Send(&pb.LoadMetrics{Metrics: promToProto(ms)})
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
