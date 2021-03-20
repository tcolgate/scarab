package ui

import (
	"embed"
	_ "embed"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	grpc "google.golang.org/grpc"
)

//go:generate protoc -I. --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative     ui.proto
//go:generate protoc -I. --js_out=import_style=commonjs:./src --grpc-web_out=import_style=typescript,mode=grpcwebtext:./src ui.proto

//go:embed dist
var staticFiles embed.FS

func server() {
	grpcServer := grpc.Server()
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	var staticFS = http.FS(staticFiles)
	fs := http.FileServer(staticFS)

	mux := http.NewServeMux()
	// Serve static files
	mux.Handle("/", fs)

	tlsHttpServer.Handler = http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(resp, req)
		}
		// Fall back to other servers.
		mux.ServeHTTP(resp, req)
	})
}
