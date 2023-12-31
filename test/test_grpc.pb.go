// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: test.proto

package test

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

// MyServiceClient is the client API for MyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyServiceClient interface {
	GetNews(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (MyService_GetNewsClient, error)
}

type myServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyServiceClient(cc grpc.ClientConnInterface) MyServiceClient {
	return &myServiceClient{cc}
}

func (c *myServiceClient) GetNews(ctx context.Context, in *NewsRequest, opts ...grpc.CallOption) (MyService_GetNewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MyService_ServiceDesc.Streams[0], "/MyService/GetNews", opts...)
	if err != nil {
		return nil, err
	}
	x := &myServiceGetNewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MyService_GetNewsClient interface {
	Recv() (*NewsArticle, error)
	grpc.ClientStream
}

type myServiceGetNewsClient struct {
	grpc.ClientStream
}

func (x *myServiceGetNewsClient) Recv() (*NewsArticle, error) {
	m := new(NewsArticle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyServiceServer is the server API for MyService service.
// All implementations must embed UnimplementedMyServiceServer
// for forward compatibility
type MyServiceServer interface {
	GetNews(*NewsRequest, MyService_GetNewsServer) error
	mustEmbedUnimplementedMyServiceServer()
}

// UnimplementedMyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyServiceServer struct {
}

func (UnimplementedMyServiceServer) GetNews(*NewsRequest, MyService_GetNewsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (UnimplementedMyServiceServer) mustEmbedUnimplementedMyServiceServer() {}

// UnsafeMyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyServiceServer will
// result in compilation errors.
type UnsafeMyServiceServer interface {
	mustEmbedUnimplementedMyServiceServer()
}

func RegisterMyServiceServer(s grpc.ServiceRegistrar, srv MyServiceServer) {
	s.RegisterService(&MyService_ServiceDesc, srv)
}

func _MyService_GetNews_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NewsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MyServiceServer).GetNews(m, &myServiceGetNewsServer{stream})
}

type MyService_GetNewsServer interface {
	Send(*NewsArticle) error
	grpc.ServerStream
}

type myServiceGetNewsServer struct {
	grpc.ServerStream
}

func (x *myServiceGetNewsServer) Send(m *NewsArticle) error {
	return x.ServerStream.SendMsg(m)
}

// MyService_ServiceDesc is the grpc.ServiceDesc for MyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MyService",
	HandlerType: (*MyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNews",
			Handler:       _MyService_GetNews_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "test.proto",
}

// MyService2Client is the client API for MyService2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyService2Client interface {
	GetNews2(ctx context.Context, opts ...grpc.CallOption) (MyService2_GetNews2Client, error)
}

type myService2Client struct {
	cc grpc.ClientConnInterface
}

func NewMyService2Client(cc grpc.ClientConnInterface) MyService2Client {
	return &myService2Client{cc}
}

func (c *myService2Client) GetNews2(ctx context.Context, opts ...grpc.CallOption) (MyService2_GetNews2Client, error) {
	stream, err := c.cc.NewStream(ctx, &MyService2_ServiceDesc.Streams[0], "/MyService2/GetNews2", opts...)
	if err != nil {
		return nil, err
	}
	x := &myService2GetNews2Client{stream}
	return x, nil
}

type MyService2_GetNews2Client interface {
	Send(*NewsRequest) error
	CloseAndRecv() (*NewsArticle, error)
	grpc.ClientStream
}

type myService2GetNews2Client struct {
	grpc.ClientStream
}

func (x *myService2GetNews2Client) Send(m *NewsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *myService2GetNews2Client) CloseAndRecv() (*NewsArticle, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(NewsArticle)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyService2Server is the server API for MyService2 service.
// All implementations must embed UnimplementedMyService2Server
// for forward compatibility
type MyService2Server interface {
	GetNews2(MyService2_GetNews2Server) error
	mustEmbedUnimplementedMyService2Server()
}

// UnimplementedMyService2Server must be embedded to have forward compatible implementations.
type UnimplementedMyService2Server struct {
}

func (UnimplementedMyService2Server) GetNews2(MyService2_GetNews2Server) error {
	return status.Errorf(codes.Unimplemented, "method GetNews2 not implemented")
}
func (UnimplementedMyService2Server) mustEmbedUnimplementedMyService2Server() {}

// UnsafeMyService2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyService2Server will
// result in compilation errors.
type UnsafeMyService2Server interface {
	mustEmbedUnimplementedMyService2Server()
}

func RegisterMyService2Server(s grpc.ServiceRegistrar, srv MyService2Server) {
	s.RegisterService(&MyService2_ServiceDesc, srv)
}

func _MyService2_GetNews2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MyService2Server).GetNews2(&myService2GetNews2Server{stream})
}

type MyService2_GetNews2Server interface {
	SendAndClose(*NewsArticle) error
	Recv() (*NewsRequest, error)
	grpc.ServerStream
}

type myService2GetNews2Server struct {
	grpc.ServerStream
}

func (x *myService2GetNews2Server) SendAndClose(m *NewsArticle) error {
	return x.ServerStream.SendMsg(m)
}

func (x *myService2GetNews2Server) Recv() (*NewsRequest, error) {
	m := new(NewsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MyService2_ServiceDesc is the grpc.ServiceDesc for MyService2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyService2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "MyService2",
	HandlerType: (*MyService2Server)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNews2",
			Handler:       _MyService2_GetNews2_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}
