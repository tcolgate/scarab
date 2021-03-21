package main

import (
	"context"
	"flag"

	_ "github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/tcolgate/scarab"
	pb "github.com/tcolgate/scarab/pb"
)

func main() {
	addr := ":8080"
	uiAddr := ":8081"
	serverAddr := "127.0.0.1:8080"
	flag.Parse()
	ctx := context.Background()

	spec := &pb.ProfileSpec{
		Profile: "myworkload",
		Version: "1234",
		Args: []*pb.ProfileArg{
			&pb.ProfileArg{
				Name:        "host",
				Description: "host to target",
				Default:     &pb.ProfileArg_String_{String_: "http://myserver.default"},
			},
		},
	}

	wrk := scarab.Workload{
		Spec: spec,
	}

	scarab.NewCombined(ctx, addr, uiAddr, serverAddr, wrk)
}
