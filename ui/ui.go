package ui

import (
	"embed"
	_ "embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	grpc "google.golang.org/grpc"
)

//go:generate protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative     ui.proto
//go:generate protoc -I. --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:./src --ts_out=service=grpc-web:./src ui.proto

//go:embed dist
var staticFiles embed.FS

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(staticFiles, "dist")
	if err != nil {
		log.Fatal(err)
	}
	return http.FS(fsys)
}

func Server(addr string, uiSrvr ManagerUIServer) {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	RegisterManagerUIServer(grpcServer, uiSrvr)

	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	fs := http.FileServer(getFileSystem())

	mux := http.NewServeMux()
	// Serve static files
	mux.Handle("/", fs)

	h := http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
		}
		// Fall back to other servers.
		mux.ServeHTTP(resp, req)
	})

	http.ListenAndServe(":8081", h)
}
