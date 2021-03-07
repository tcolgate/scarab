package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"

	"github.com/tcolgate/scarab"
	pb "github.com/tcolgate/scarab/pb"
	"google.golang.org/grpc"
)

func main() {
	addr := ":8080"
	serverAddr := "127.0.0.1:8080"
	flag.Parse()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterManagerServer(grpcServer, &scarab.Server{})
	pb.RegisterWorkerServer(grpcServer, &scarab.Server{})

	grpcServer.Serve(lis)

	ctx := context.Background()
	var clientOpts []grpc.DialOption
	conn, err := grpc.Dial(serverAddr, clientOpts...)
	if err != nil {
		log.Panicf("dial error, %v", err)
		return
	}
	defer conn.Close()
	worker := pb.NewManagerClient(conn)

	profileArgs := []*pb.ProfileArg{
		&pb.ProfileArg{
			Name:        "host",
			Description: "host to target",
			Default:     &pb.ProfileArg_String_{String_: "http://myserver.default"},
		},
	}
	hn, _ := os.Hostname()
	req := pb.RegisterProfileRequest{
		Profile: "myprofile",
		Name:    hn,
		Addr:    hn + addr,
		Args:    profileArgs,
	}
	cli, err := worker.RegisterProfile(ctx, &req)
	if err != nil {
		log.Panicf("profile registration error, %v", err)
		return
	}

	go func() {
		// throw away the keepalives.
		for {
			_, err := cli.Recv()
			if err != nil {
				return
			}
		}
	}()

}
