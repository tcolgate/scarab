package main

import (
	"flag"
	"log"
	"net"

	"github.com/tcolgate/scarab"
	pb "github.com/tcolgate/scarab/pb"
	"google.golang.org/grpc"
)

func main() {
	addr := ":8080"
	flag.Parse()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterManagerServer(grpcServer, &scarab.Server{})
	grpcServer.Serve(lis)
}
