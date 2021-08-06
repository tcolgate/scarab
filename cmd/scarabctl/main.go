package main

import (
	"context"
	"log"
	"net"

	"github.com/alecthomas/kong"
	"github.com/tcolgate/scarab"
	pb "github.com/tcolgate/scarab/pb"
	"google.golang.org/grpc"
)

type Context struct {
	Debug bool
}

type RunCmd struct {
	Addr    string `kong:"default=':8081',help='address for gRPC to connect to.'"`
	Profile string `kong:"arg,help='the profile to run.'"`
	Version string `kong:"arg,help='the profile version.'"`
	Name    string `kong:"help='a name for the job'"`
	Users   uint64 `kong:"default='1',help='how many users to run'"`
}

func (cmd *RunCmd) Run(cctx *Context) error {
	var err error
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(cmd.Addr, opts...)
	if err != nil {
		return err
	}

	client := pb.NewManagerUIClient(conn)
	ctx := context.Background()

	req := &pb.StartJobRequest{
		Profile: cmd.Profile,
		Version: cmd.Version,
		Users:   cmd.Users,
	}

	resp, err := client.StartJob(ctx, req)
	log.Printf("got id %v", resp.GetId())

	return err
}

type GetCmd struct {
	Addr string `kong:"default=':8081',help='address for gRPC to connect to.'"`
}

func (cmd *GetCmd) Run(cctx *Context) error {
	var err error
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(cmd.Addr, opts...)
	if err != nil {
		return err
	}

	client := pb.NewManagerUIClient(conn)

	ctx := context.Background()
	ps, err := client.ListProfiles(ctx, &pb.ListProfilesRequest{})
	if err != nil {
		return err
	}

	for _, p := range ps.GetRegistered() {
		log.Printf("p: %#v", p.String())
	}

	return err
}

type ManagerCmd struct {
	UIAddr string `kong:"default=':8080',help='address for the UI to listen on.'"`
	Addr   string `kong:"default=':8081',help='address for gRPC to listen on.'"`
}

func (cmd *ManagerCmd) Run(ctx *Context) error {
	lis, err := net.Listen("tcp", cmd.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	mngr := scarab.NewManager()
	pb.RegisterManagerServer(grpcServer, mngr)
	pb.RegisterManagerUIServer(grpcServer, mngr)

	go func() {
		err = grpcServer.Serve(lis)
		log.Printf("listen err, %v", err)
	}()

	log.Printf("UI listening on %s", cmd.UIAddr)
	scarab.UIServer(cmd.UIAddr, mngr)
	return err
}

var CLI = struct {
	Debug bool `kong:"cmd,help='enable debug.'"`

	Manager ManagerCmd `kong:"cmd,help='run manager.'"`
	Get     GetCmd     `kong:"cmd,help='query manager.'"`
	Run     RunCmd     `kong:"cmd,help='run a profile.'"`
}{
	Debug: false,
}

func main() {
	ctx := kong.Parse(&CLI)
	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&Context{
		Debug: CLI.Debug,
	})
	ctx.FatalIfErrorf(err)
}
