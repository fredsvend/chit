// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Chit_Chat

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

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	Publish(ctx context.Context, opts ...grpc.CallOption) (Chat_PublishClient, error)
	Broadcast(ctx context.Context, opts ...grpc.CallOption) (Chat_BroadcastClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) Publish(ctx context.Context, opts ...grpc.CallOption) (Chat_PublishClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[0], "/chitChat.Chat/Publish", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatPublishClient{stream}
	return x, nil
}

type Chat_PublishClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatPublishClient struct {
	grpc.ClientStream
}

func (x *chatPublishClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatPublishClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) Broadcast(ctx context.Context, opts ...grpc.CallOption) (Chat_BroadcastClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chat_ServiceDesc.Streams[1], "/chitChat.Chat/Broadcast", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatBroadcastClient{stream}
	return x, nil
}

type Chat_BroadcastClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatBroadcastClient struct {
	grpc.ClientStream
}

func (x *chatBroadcastClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatBroadcastClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	Publish(Chat_PublishServer) error
	Broadcast(Chat_BroadcastServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) Publish(Chat_PublishServer) error {
	return status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedChatServer) Broadcast(Chat_BroadcastServer) error {
	return status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&Chat_ServiceDesc, srv)
}

func _Chat_Publish_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Publish(&chatPublishServer{stream})
}

type Chat_PublishServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatPublishServer struct {
	grpc.ServerStream
}

func (x *chatPublishServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatPublishServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Chat_Broadcast_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).Broadcast(&chatBroadcastServer{stream})
}

type Chat_BroadcastServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatBroadcastServer struct {
	grpc.ServerStream
}

func (x *chatBroadcastServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatBroadcastServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Chat_ServiceDesc is the grpc.ServiceDesc for Chat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chitChat.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Publish",
			Handler:       _Chat_Publish_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Broadcast",
			Handler:       _Chat_Broadcast_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "services.proto",
}