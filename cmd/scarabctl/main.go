package main

import (
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

type ManagerCmd struct {
	UIAddr string `kong:"default=':8080',help='address for the UI to listen on.'"`
	Addr   string `kong:"default=':8081',help='address for gRPC to listen on.'"`
}

func (cmd *ManagerCmd) Run(ctx *Context) error {
	log.Printf("cmd %#v", CLI)
	log.Printf("cmd %#v", cmd)
	log.Printf("ctx %#v", ctx)
	log.Printf("grpc listening on %s", cmd.Addr)
	lis, err := net.Listen("tcp", cmd.Addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	mngr := scarab.NewManager()
	pb.RegisterManagerServer(grpcServer, mngr)

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
