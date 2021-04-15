package main

import (
	"context"
	"flag"
	"log"
	"net"

	"github.com/tcolgate/scarab"
	pb "github.com/tcolgate/scarab/pb"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	addr := flag.String("addr", ":8083", "address to listen on")
	serverAddr := flag.String("manager.addr", "localhost:8081", "address of the manager")
	flag.Parse()

	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	spec := &pb.ProfileSpec{
		Profile: "my-other-workload",
		Version: "1",
		Args: []*pb.ProfileArg{
			&pb.ProfileArg{
				Name:        "host",
				Description: "host to target",
				Default:     &pb.JobArgValue{Value: &pb.JobArgValue_String_{}},
			},
		},
	}

	wrk := scarab.Workload{
		Spec: spec,
	}

	wrkr, err := scarab.NewWorker(ctx, *addr, *serverAddr, wrk)
	if err != nil {
		log.Fatalf("failed to create worker: %v", err)
	}
	defer wrkr.Close()
	pb.RegisterWorkerServer(grpcServer, wrkr)

	err = grpcServer.Serve(lis)
	log.Printf("listen err, %v", err)

}
