package scarab

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	model "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"github.com/prometheus/prometheus/pkg/labels"
	tsdb "github.com/prometheus/prometheus/tsdb"
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

func getValue(t *model.MetricType, m *model.Metric) float64 {
	switch *t {
	case model.MetricType_GAUGE:
		return *m.Gauge.Value
	case model.MetricType_COUNTER:
		return *m.Counter.Value
	case model.MetricType_UNTYPED:
		return *m.Untyped.Value
		//	case model.MetricType_HISTOGRAM:
		//		return *m.Histogram.SampleCount)q
		//	case model.MetricType_SUMMARY:
		//		return *m.Summary.Quantile
	default:
		return 0
	}
	/*
				Gauge                *Gauge       `protobuf:"bytes,2,opt,name=gauge" json:"gauge,omitempty"`
				Counter              *Counter     `protobuf:"bytes,3,opt,name=counter" json:"counter,omitempty"`
		 		Summary              *Summary     `protobuf:"bytes,4,opt,name=summary" json:"summary,omitempty"`
		 		Untyped              *Untyped     `protobuf:"bytes,5,opt,name=untyped" json:"untyped,omitempty"`
		 		Histogram            *Histogram   `protobuf:"bytes,7,opt,name=histogram" json:"histogram,omitempty"`
	*/

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

	workerAddrLabel := "worker_addr"
	for {
		lresp, err := loadrep.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for _, mf := range lresp.Metrics {
			for i := range mf.Metric {
				mf.Metric[i].Label = append(mf.Metric[i].Label, &model.LabelPair{Name: &workerAddrLabel, Value: &addr})
			}

			expfmt.MetricFamilyToText(os.Stdout, mf)
			// These should be sent off to a bunch of interested
			// subscribers
		}
	}
}

func splitUserCounts(workers, users uint) []uint {
	if workers == 0 {
		return nil
	}
	ws := make([]uint, workers)
	base := users / workers
	extra := users % workers
	for i := range ws {
		ws[i] = base
		if extra > 0 {
			ws[i]++
			extra--
		}
	}
	return ws
}

type metricsWriter struct {
	db *tsdb.DB
}

/*
	db, err := tsdb.Open("tsdb/", nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
*/

func (mw *metricsWriter) WriteMetrics(ctx context.Context, ms []*model.MetricFamily, ls []labels.Label) error {
	/* this needs to do the whole AppendFast thing */
	log.Printf("got metrics: %#v", ms)
	app := mw.db.Appender(ctx)
	for _, mf := range ms {
		mt := mf.Type
		for i := range mf.Metric {
			m := mf.Metric[i]
			m.Label = append(m.Label, &model.LabelPair{Name: &workerAddrLabel, Value: &addr})
			v := getValue(mt, m)
			app.Append(0, m.Label, m.GetTimestampMs(), v)
		}
	}
	app.Commit()

	return nil
}

func (m *Manager) RunJob(ctx context.Context, j *pb.StartJobRequest) (*pb.StartJobResponse, error) {
	clientOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	mw := &metricsWriter{}

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
		rjsrc, err := wcli.RunProfile(ctx)
		if err != nil {
			return nil, err
		}
		go func() {
			ls := []labels.Label{{"worker_addr", addr}}
			for {
				ms, err := rjsrc.Recv()
				if err != nil {
					log.Printf("error getting loadreport, %#v", err)
					return
				}
				mw.WriteMetrics(ctx, ms, ls)
			}
		}()
		err = rjsrc.Send(&pb.RunProfileRequest{
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
	return s.RunJob(ctx, j)
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
