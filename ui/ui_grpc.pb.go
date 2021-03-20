// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ui

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ManagerUIClient is the client API for ManagerUI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerUIClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	ListProfiles(ctx context.Context, in *ListProfilesRequest, opts ...grpc.CallOption) (*ListProfilesResponse, error)
}

type managerUIClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerUIClient(cc grpc.ClientConnInterface) ManagerUIClient {
	return &managerUIClient{cc}
}

func (c *managerUIClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	out := new(StartJobResponse)
	err := c.cc.Invoke(ctx, "/ui.ManagerUI/StartJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerUIClient) ListProfiles(ctx context.Context, in *ListProfilesRequest, opts ...grpc.CallOption) (*ListProfilesResponse, error) {
	out := new(ListProfilesResponse)
	err := c.cc.Invoke(ctx, "/ui.ManagerUI/ListProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerUIServer is the server API for ManagerUI service.
// All implementations must embed UnimplementedManagerUIServer
// for forward compatibility
type ManagerUIServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	ListProfiles(context.Context, *ListProfilesRequest) (*ListProfilesResponse, error)
	mustEmbedUnimplementedManagerUIServer()
}

// UnimplementedManagerUIServer must be embedded to have forward compatible implementations.
type UnimplementedManagerUIServer struct {
}

func (UnimplementedManagerUIServer) StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedManagerUIServer) ListProfiles(context.Context, *ListProfilesRequest) (*ListProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProfiles not implemented")
}
func (UnimplementedManagerUIServer) mustEmbedUnimplementedManagerUIServer() {}

// UnsafeManagerUIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerUIServer will
// result in compilation errors.
type UnsafeManagerUIServer interface {
	mustEmbedUnimplementedManagerUIServer()
}

func RegisterManagerUIServer(s grpc.ServiceRegistrar, srv ManagerUIServer) {
	s.RegisterService(&ManagerUI_ServiceDesc, srv)
}

func _ManagerUI_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerUIServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.ManagerUI/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerUIServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ManagerUI_ListProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerUIServer).ListProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ui.ManagerUI/ListProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerUIServer).ListProfiles(ctx, req.(*ListProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ManagerUI_ServiceDesc is the grpc.ServiceDesc for ManagerUI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ManagerUI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ui.ManagerUI",
	HandlerType: (*ManagerUIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _ManagerUI_StartJob_Handler,
		},
		{
			MethodName: "ListProfiles",
			Handler:    _ManagerUI_ListProfiles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ui.proto",
}
