package main

import (
	"context"
	"log"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tcolgate/scarab"
	pb "github.com/tcolgate/scarab/pb"
)

func setup(r prometheus.Registerer) scarab.Runner {
	h := prometheus.NewHistogram(
		prometheus.HistogramOpts{},
	)
	r.MustRegister(h)
	return scarab.RunnerFunc(func(ctx context.Context, args []*pb.JobArg) {
		f := newFake(
			2*time.Millisecond,
			1*time.Millisecond,
			10*time.Millisecond,
			2*time.Millisecond,
			0.70)
		for {
			select {
			case <-ctx.Done():
			default:
				func() {
					tmr := prometheus.NewTimer(h)
					defer tmr.ObserveDuration()
					d, _ := f.Sample()

					time.Sleep(d)
				}()
			}
		}
	})
}

func main() {
	args := []*pb.ProfileArg{{
		Name:        "host",
		Description: "host to target",
		Default: &pb.JobArgValue{
			Value: &pb.JobArgValue_String_{
				String_: "http://localhost:8089",
			},
		},
	}}

	err := scarab.RunStandaloneWorker(
		"my-other-workload",
		"1",
		"",
		args,
		scarab.UserFunc(setup),
	)

	if err != nil {
		log.Fatalf("failed to create worker: %v", err)
	}
}
