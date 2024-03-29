// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package notify

import (
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NotifyServiceClient is the client API for NotifyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotifyServiceClient interface {
}

type notifyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotifyServiceClient(cc grpc.ClientConnInterface) NotifyServiceClient {
	return &notifyServiceClient{cc}
}

// NotifyServiceServer is the server API for NotifyService service.
// All implementations must embed UnimplementedNotifyServiceServer
// for forward compatibility
type NotifyServiceServer interface {
	mustEmbedUnimplementedNotifyServiceServer()
}

// UnimplementedNotifyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotifyServiceServer struct {
}

func (UnimplementedNotifyServiceServer) mustEmbedUnimplementedNotifyServiceServer() {}

// UnsafeNotifyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotifyServiceServer will
// result in compilation errors.
type UnsafeNotifyServiceServer interface {
	mustEmbedUnimplementedNotifyServiceServer()
}

func RegisterNotifyServiceServer(s grpc.ServiceRegistrar, srv NotifyServiceServer) {
	s.RegisterService(&NotifyService_ServiceDesc, srv)
}

// NotifyService_ServiceDesc is the grpc.ServiceDesc for NotifyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotifyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notify.notifyService",
	HandlerType: (*NotifyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "protos/notify.proto",
}
