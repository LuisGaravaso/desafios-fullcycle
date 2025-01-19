// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: internal/infra/grpc/protofiles/create_order.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CreateOrderService_CreateOrder_FullMethodName = "/pb.CreateOrderService/CreateOrder"
)

// CreateOrderServiceClient is the client API for CreateOrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreateOrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
}

type createOrderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreateOrderServiceClient(cc grpc.ClientConnInterface) CreateOrderServiceClient {
	return &createOrderServiceClient{cc}
}

func (c *createOrderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, CreateOrderService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreateOrderServiceServer is the server API for CreateOrderService service.
// All implementations must embed UnimplementedCreateOrderServiceServer
// for forward compatibility.
type CreateOrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	mustEmbedUnimplementedCreateOrderServiceServer()
}

// UnimplementedCreateOrderServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCreateOrderServiceServer struct{}

func (UnimplementedCreateOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedCreateOrderServiceServer) mustEmbedUnimplementedCreateOrderServiceServer() {}
func (UnimplementedCreateOrderServiceServer) testEmbeddedByValue()                            {}

// UnsafeCreateOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreateOrderServiceServer will
// result in compilation errors.
type UnsafeCreateOrderServiceServer interface {
	mustEmbedUnimplementedCreateOrderServiceServer()
}

func RegisterCreateOrderServiceServer(s grpc.ServiceRegistrar, srv CreateOrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedCreateOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CreateOrderService_ServiceDesc, srv)
}

func _CreateOrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreateOrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreateOrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreateOrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreateOrderService_ServiceDesc is the grpc.ServiceDesc for CreateOrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreateOrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CreateOrderService",
	HandlerType: (*CreateOrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _CreateOrderService_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/infra/grpc/protofiles/create_order.proto",
}
