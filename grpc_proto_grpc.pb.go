// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: grpc_proto.proto

package grpc_proto

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

// SSHServiceClient is the client API for SSHService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SSHServiceClient interface {
	ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (SSHService_ExecuteCommandClient, error)
}

type sSHServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSSHServiceClient(cc grpc.ClientConnInterface) SSHServiceClient {
	return &sSHServiceClient{cc}
}

func (c *sSHServiceClient) ExecuteCommand(ctx context.Context, opts ...grpc.CallOption) (SSHService_ExecuteCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &SSHService_ServiceDesc.Streams[0], "/grpc_proto.SSHService/ExecuteCommand", opts...)
	if err != nil {
		return nil, err
	}
	x := &sSHServiceExecuteCommandClient{stream}
	return x, nil
}

type SSHService_ExecuteCommandClient interface {
	Send(*CommandRequest) error
	Recv() (*CommandResponse, error)
	grpc.ClientStream
}

type sSHServiceExecuteCommandClient struct {
	grpc.ClientStream
}

func (x *sSHServiceExecuteCommandClient) Send(m *CommandRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sSHServiceExecuteCommandClient) Recv() (*CommandResponse, error) {
	m := new(CommandResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SSHServiceServer is the server API for SSHService service.
// All implementations must embed UnimplementedSSHServiceServer
// for forward compatibility
type SSHServiceServer interface {
	ExecuteCommand(SSHService_ExecuteCommandServer) error
	mustEmbedUnimplementedSSHServiceServer()
}

// UnimplementedSSHServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSSHServiceServer struct {
}

func (UnimplementedSSHServiceServer) ExecuteCommand(SSHService_ExecuteCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedSSHServiceServer) mustEmbedUnimplementedSSHServiceServer() {}

// UnsafeSSHServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SSHServiceServer will
// result in compilation errors.
type UnsafeSSHServiceServer interface {
	mustEmbedUnimplementedSSHServiceServer()
}

func RegisterSSHServiceServer(s grpc.ServiceRegistrar, srv SSHServiceServer) {
	s.RegisterService(&SSHService_ServiceDesc, srv)
}

func _SSHService_ExecuteCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SSHServiceServer).ExecuteCommand(&sSHServiceExecuteCommandServer{stream})
}

type SSHService_ExecuteCommandServer interface {
	Send(*CommandResponse) error
	Recv() (*CommandRequest, error)
	grpc.ServerStream
}

type sSHServiceExecuteCommandServer struct {
	grpc.ServerStream
}

func (x *sSHServiceExecuteCommandServer) Send(m *CommandResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sSHServiceExecuteCommandServer) Recv() (*CommandRequest, error) {
	m := new(CommandRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SSHService_ServiceDesc is the grpc.ServiceDesc for SSHService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SSHService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_proto.SSHService",
	HandlerType: (*SSHServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteCommand",
			Handler:       _SSHService_ExecuteCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc_proto.proto",
}
