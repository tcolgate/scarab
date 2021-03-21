package scarab

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/tcolgate/scarab/pb"
	grpc "google.golang.org/grpc"
)

type Workload struct {
	Spec *pb.ProfileSpec
}

func NewCombined(addr, uiAddr, serverAddr string, wrk Workload) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	mngr := NewManager()
	wrkr, err := NewWorker()
	if err != nil {
		log.Fatalf("failed to create worker: %v", err)
	}
	pb.RegisterManagerServer(grpcServer, mngr)
	pb.RegisterWorkerServer(grpcServer, wrkr)

	go func() {
		err := grpcServer.Serve(lis)
		log.Printf("listen err, %v", err)
	}()

	ctx := context.Background()
	clientOpts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(serverAddr, clientOpts...)
	if err != nil {
		log.Panicf("dial error, %v", err)
		return
	}
	defer conn.Close()

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
		log.Panicf("profile registration error, %v", err)
		return
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

	srvconn, err := grpc.Dial(serverAddr, clientOpts...)
	if err != nil {
		log.Panicf("dial error, %v", err)
		return
	}
	defer conn.Close()
	wcli := pb.NewWorkerClient(srvconn)
	loadrep, err := wcli.ReportLoad(ctx, &pb.ReportLoadRequest{})

	go func() {
		// throw away the keepalives.
		for {
			msg, err := loadrep.Recv()
			if err != nil {
				return
			}
			for range msg.Metrics {
			}
		}
	}()

	rjsrc, err := wcli.RunJob(ctx)
	if err != nil {
		log.Panicf("dial error, %v", err)
		return
	}

	go func() {
		// throw away the keepalives.
		for {
			_, err := rjsrc.Recv()
			if err != nil {
				log.Printf("error getting loadreport, %#v", err)
				return
			}
		}
	}()

	err = rjsrc.Send(&pb.RunJobRequest{
		Profile: "myprofile",
		Args:    []*pb.JobArg{},
	})
	if err != nil {
		log.Panicf("runjub error, %v", err)
		return
	}

	Server(uiAddr, mngr)
}
