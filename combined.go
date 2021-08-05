package scarab

import (
	"context"
	"log"
	"net"

	pb "github.com/tcolgate/scarab/pb"
	grpc "google.golang.org/grpc"
)

type Workload struct {
	Spec *pb.ProfileSpec
}

func NewCombined(ctx context.Context, addr, uiAddr, serverAddr string, wrk Workload, user User) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	mngr := NewManager()
	pb.RegisterManagerServer(grpcServer, mngr)

	wrkr, err := NewWorker(ctx, addr, serverAddr, wrk, user)
	if err != nil {
		log.Fatalf("failed to create worker: %v", err)
	}
	defer wrkr.Close()
	pb.RegisterWorkerServer(grpcServer, wrkr)

	go func() {
		err := grpcServer.Serve(lis)
		log.Printf("listen err, %v", err)
	}()

	UIServer(uiAddr, mngr)
}
