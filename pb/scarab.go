package scarab

//go:generate protoc -I.:./include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative     scarab.proto scarab-ui.proto scarab-common.proto
//go:generate protoc -I.:./include --plugin=protoc-gen-ts=../ui/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:../ui/src --ts_out=service=grpc-web:../ui/src scarab-ui.proto scarab-common.proto
