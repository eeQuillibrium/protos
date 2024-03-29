// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: orders/orders.proto

package grpc_orders

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

// OrderingClient is the client API for Ordering service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderingClient interface {
	SendOrder(ctx context.Context, in *SendOrderReq, opts ...grpc.CallOption) (*EmptyOrderResp, error)
}

type orderingClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderingClient(cc grpc.ClientConnInterface) OrderingClient {
	return &orderingClient{cc}
}

func (c *orderingClient) SendOrder(ctx context.Context, in *SendOrderReq, opts ...grpc.CallOption) (*EmptyOrderResp, error) {
	out := new(EmptyOrderResp)
	err := c.cc.Invoke(ctx, "/orders.Ordering/SendOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderingServer is the server API for Ordering service.
// All implementations must embed UnimplementedOrderingServer
// for forward compatibility
type OrderingServer interface {
	SendOrder(context.Context, *SendOrderReq) (*EmptyOrderResp, error)
	mustEmbedUnimplementedOrderingServer()
}

// UnimplementedOrderingServer must be embedded to have forward compatible implementations.
type UnimplementedOrderingServer struct {
}

func (UnimplementedOrderingServer) SendOrder(context.Context, *SendOrderReq) (*EmptyOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrder not implemented")
}
func (UnimplementedOrderingServer) mustEmbedUnimplementedOrderingServer() {}

// UnsafeOrderingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderingServer will
// result in compilation errors.
type UnsafeOrderingServer interface {
	mustEmbedUnimplementedOrderingServer()
}

func RegisterOrderingServer(s grpc.ServiceRegistrar, srv OrderingServer) {
	s.RegisterService(&Ordering_ServiceDesc, srv)
}

func _Ordering_SendOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderingServer).SendOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orders.Ordering/SendOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderingServer).SendOrder(ctx, req.(*SendOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ordering_ServiceDesc is the grpc.ServiceDesc for Ordering service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ordering_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "orders.Ordering",
	HandlerType: (*OrderingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOrder",
			Handler:    _Ordering_SendOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders/orders.proto",
}
