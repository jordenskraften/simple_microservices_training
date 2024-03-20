// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0--rc3
// source: service.proto

package generated

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

// SimpleMessageServiceClient is the client API for SimpleMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleMessageServiceClient interface {
	SimpleMessageMethod(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error)
}

type simpleMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleMessageServiceClient(cc grpc.ClientConnInterface) SimpleMessageServiceClient {
	return &simpleMessageServiceClient{cc}
}

func (c *simpleMessageServiceClient) SimpleMessageMethod(ctx context.Context, in *SimpleMessage, opts ...grpc.CallOption) (*SimpleMessage, error) {
	out := new(SimpleMessage)
	err := c.cc.Invoke(ctx, "/SimpleMessageService/SimpleMessageMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleMessageServiceServer is the server API for SimpleMessageService service.
// All implementations must embed UnimplementedSimpleMessageServiceServer
// for forward compatibility
type SimpleMessageServiceServer interface {
	SimpleMessageMethod(context.Context, *SimpleMessage) (*SimpleMessage, error)
	//mustEmbedUnimplementedSimpleMessageServiceServer()
}

// UnimplementedSimpleMessageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleMessageServiceServer struct {
}

func (UnimplementedSimpleMessageServiceServer) SimpleMessageMethod(context.Context, *SimpleMessage) (*SimpleMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SimpleMessageMethod not implemented")
}
func (UnimplementedSimpleMessageServiceServer) mustEmbedUnimplementedSimpleMessageServiceServer() {}

// UnsafeSimpleMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleMessageServiceServer will
// result in compilation errors.
type UnsafeSimpleMessageServiceServer interface {
	mustEmbedUnimplementedSimpleMessageServiceServer()
}

func RegisterSimpleMessageServiceServer(s grpc.ServiceRegistrar, srv SimpleMessageServiceServer) {
	s.RegisterService(&SimpleMessageService_ServiceDesc, srv)
}

func _SimpleMessageService_SimpleMessageMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleMessageServiceServer).SimpleMessageMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimpleMessageService/SimpleMessageMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleMessageServiceServer).SimpleMessageMethod(ctx, req.(*SimpleMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleMessageService_ServiceDesc is the grpc.ServiceDesc for SimpleMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SimpleMessageService",
	HandlerType: (*SimpleMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SimpleMessageMethod",
			Handler:    _SimpleMessageService_SimpleMessageMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}