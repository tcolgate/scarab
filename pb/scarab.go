package scarab

//go:generate protoc -I.:./include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative     scarab.proto
