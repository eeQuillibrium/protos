// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: kitchen.proto

package nikita_kitchen1

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

// KitchenClient is the client API for Kitchen service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KitchenClient interface {
	SendOrder(ctx context.Context, in *SendOrderReq, opts ...grpc.CallOption) (*EmptyOrderResp, error)
}

type kitchenClient struct {
	cc grpc.ClientConnInterface
}

func NewKitchenClient(cc grpc.ClientConnInterface) KitchenClient {
	return &kitchenClient{cc}
}

func (c *kitchenClient) SendOrder(ctx context.Context, in *SendOrderReq, opts ...grpc.CallOption) (*EmptyOrderResp, error) {
	out := new(EmptyOrderResp)
	err := c.cc.Invoke(ctx, "/kitchen.Kitchen/SendOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KitchenServer is the server API for Kitchen service.
// All implementations must embed UnimplementedKitchenServer
// for forward compatibility
type KitchenServer interface {
	SendOrder(context.Context, *SendOrderReq) (*EmptyOrderResp, error)
	mustEmbedUnimplementedKitchenServer()
}

// UnimplementedKitchenServer must be embedded to have forward compatible implementations.
type UnimplementedKitchenServer struct {
}

func (UnimplementedKitchenServer) SendOrder(context.Context, *SendOrderReq) (*EmptyOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOrder not implemented")
}
func (UnimplementedKitchenServer) mustEmbedUnimplementedKitchenServer() {}

// UnsafeKitchenServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KitchenServer will
// result in compilation errors.
type UnsafeKitchenServer interface {
	mustEmbedUnimplementedKitchenServer()
}

func RegisterKitchenServer(s grpc.ServiceRegistrar, srv KitchenServer) {
	s.RegisterService(&Kitchen_ServiceDesc, srv)
}

func _Kitchen_SendOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KitchenServer).SendOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kitchen.Kitchen/SendOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KitchenServer).SendOrder(ctx, req.(*SendOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Kitchen_ServiceDesc is the grpc.ServiceDesc for Kitchen service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Kitchen_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kitchen.Kitchen",
	HandlerType: (*KitchenServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOrder",
			Handler:    _Kitchen_SendOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kitchen.proto",
}
