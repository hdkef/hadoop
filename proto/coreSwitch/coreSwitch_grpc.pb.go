// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.7
// source: proto/coreSwitch/coreSwitch.proto

package coreSwitch

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

const (
	CoreSwitch_Write_FullMethodName = "/CoreSwitch/Write"
)

// CoreSwitchClient is the client API for CoreSwitch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreSwitchClient interface {
	Write(ctx context.Context, in *WriteReq, opts ...grpc.CallOption) (CoreSwitch_WriteClient, error)
}

type coreSwitchClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreSwitchClient(cc grpc.ClientConnInterface) CoreSwitchClient {
	return &coreSwitchClient{cc}
}

func (c *coreSwitchClient) Write(ctx context.Context, in *WriteReq, opts ...grpc.CallOption) (CoreSwitch_WriteClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoreSwitch_ServiceDesc.Streams[0], CoreSwitch_Write_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &coreSwitchWriteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoreSwitch_WriteClient interface {
	Recv() (*WriteRes, error)
	grpc.ClientStream
}

type coreSwitchWriteClient struct {
	grpc.ClientStream
}

func (x *coreSwitchWriteClient) Recv() (*WriteRes, error) {
	m := new(WriteRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CoreSwitchServer is the server API for CoreSwitch service.
// All implementations must embed UnimplementedCoreSwitchServer
// for forward compatibility
type CoreSwitchServer interface {
	Write(*WriteReq, CoreSwitch_WriteServer) error
	mustEmbedUnimplementedCoreSwitchServer()
}

// UnimplementedCoreSwitchServer must be embedded to have forward compatible implementations.
type UnimplementedCoreSwitchServer struct {
}

func (UnimplementedCoreSwitchServer) Write(*WriteReq, CoreSwitch_WriteServer) error {
	return status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedCoreSwitchServer) mustEmbedUnimplementedCoreSwitchServer() {}

// UnsafeCoreSwitchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreSwitchServer will
// result in compilation errors.
type UnsafeCoreSwitchServer interface {
	mustEmbedUnimplementedCoreSwitchServer()
}

func RegisterCoreSwitchServer(s grpc.ServiceRegistrar, srv CoreSwitchServer) {
	s.RegisterService(&CoreSwitch_ServiceDesc, srv)
}

func _CoreSwitch_Write_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WriteReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoreSwitchServer).Write(m, &coreSwitchWriteServer{stream})
}

type CoreSwitch_WriteServer interface {
	Send(*WriteRes) error
	grpc.ServerStream
}

type coreSwitchWriteServer struct {
	grpc.ServerStream
}

func (x *coreSwitchWriteServer) Send(m *WriteRes) error {
	return x.ServerStream.SendMsg(m)
}

// CoreSwitch_ServiceDesc is the grpc.ServiceDesc for CoreSwitch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreSwitch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CoreSwitch",
	HandlerType: (*CoreSwitchServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Write",
			Handler:       _CoreSwitch_Write_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/coreSwitch/coreSwitch.proto",
}
