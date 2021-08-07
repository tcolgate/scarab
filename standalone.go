package scarab

import (
	"context"
	"flag"
	"fmt"
	"net"

	pb "github.com/tcolgate/scarab/pb"
	grpc "google.golang.org/grpc"
)

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
