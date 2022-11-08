// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: streaming.proto

package example

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

// StreamingServiceClient is the client API for StreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingServiceClient interface {
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_ClientStreamClient, error)
	ResponseStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (StreamingService_ResponseStreamClient, error)
	BiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_BiDirectionalStreamClient, error)
}

type streamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingServiceClient(cc grpc.ClientConnInterface) StreamingServiceClient {
	return &streamingServiceClient{cc}
}

func (c *streamingServiceClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_ClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[0], "/tutorial.StreamingService/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceClientStreamClient{stream}
	return x, nil
}

type StreamingService_ClientStreamClient interface {
	Send(*GreetRequest) error
	CloseAndRecv() (*GreetResponse, error)
	grpc.ClientStream
}

type streamingServiceClientStreamClient struct {
	grpc.ClientStream
}

func (x *streamingServiceClientStreamClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingServiceClientStreamClient) CloseAndRecv() (*GreetResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingServiceClient) ResponseStream(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (StreamingService_ResponseStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[1], "/tutorial.StreamingService/ResponseStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceResponseStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingService_ResponseStreamClient interface {
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type streamingServiceResponseStreamClient struct {
	grpc.ClientStream
}

func (x *streamingServiceResponseStreamClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingServiceClient) BiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (StreamingService_BiDirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[2], "/tutorial.StreamingService/BiDirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingServiceBiDirectionalStreamClient{stream}
	return x, nil
}

type StreamingService_BiDirectionalStreamClient interface {
	Send(*GreetRequest) error
	Recv() (*GreetResponse, error)
	grpc.ClientStream
}

type streamingServiceBiDirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *streamingServiceBiDirectionalStreamClient) Send(m *GreetRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingServiceBiDirectionalStreamClient) Recv() (*GreetResponse, error) {
	m := new(GreetResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingServiceServer is the server API for StreamingService service.
// All implementations must embed UnimplementedStreamingServiceServer
// for forward compatibility
type StreamingServiceServer interface {
	ClientStream(StreamingService_ClientStreamServer) error
	ResponseStream(*GreetRequest, StreamingService_ResponseStreamServer) error
	BiDirectionalStream(StreamingService_BiDirectionalStreamServer) error
	mustEmbedUnimplementedStreamingServiceServer()
}

// UnimplementedStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingServiceServer struct {
}

func (UnimplementedStreamingServiceServer) ClientStream(StreamingService_ClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStream not implemented")
}
func (UnimplementedStreamingServiceServer) ResponseStream(*GreetRequest, StreamingService_ResponseStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ResponseStream not implemented")
}
func (UnimplementedStreamingServiceServer) BiDirectionalStream(StreamingService_BiDirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalStream not implemented")
}
func (UnimplementedStreamingServiceServer) mustEmbedUnimplementedStreamingServiceServer() {}

// UnsafeStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServiceServer will
// result in compilation errors.
type UnsafeStreamingServiceServer interface {
	mustEmbedUnimplementedStreamingServiceServer()
}

func RegisterStreamingServiceServer(s grpc.ServiceRegistrar, srv StreamingServiceServer) {
	s.RegisterService(&StreamingService_ServiceDesc, srv)
}

func _StreamingService_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServiceServer).ClientStream(&streamingServiceClientStreamServer{stream})
}

type StreamingService_ClientStreamServer interface {
	SendAndClose(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type streamingServiceClientStreamServer struct {
	grpc.ServerStream
}

func (x *streamingServiceClientStreamServer) SendAndClose(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingServiceClientStreamServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamingService_ResponseStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServiceServer).ResponseStream(m, &streamingServiceResponseStreamServer{stream})
}

type StreamingService_ResponseStreamServer interface {
	Send(*GreetResponse) error
	grpc.ServerStream
}

type streamingServiceResponseStreamServer struct {
	grpc.ServerStream
}

func (x *streamingServiceResponseStreamServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamingService_BiDirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingServiceServer).BiDirectionalStream(&streamingServiceBiDirectionalStreamServer{stream})
}

type StreamingService_BiDirectionalStreamServer interface {
	Send(*GreetResponse) error
	Recv() (*GreetRequest, error)
	grpc.ServerStream
}

type streamingServiceBiDirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *streamingServiceBiDirectionalStreamServer) Send(m *GreetResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingServiceBiDirectionalStreamServer) Recv() (*GreetRequest, error) {
	m := new(GreetRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingService_ServiceDesc is the grpc.ServiceDesc for StreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tutorial.StreamingService",
	HandlerType: (*StreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _StreamingService_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ResponseStream",
			Handler:       _StreamingService_ResponseStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BiDirectionalStream",
			Handler:       _StreamingService_BiDirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "streaming.proto",
}
