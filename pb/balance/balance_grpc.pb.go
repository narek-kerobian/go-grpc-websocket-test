// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: pb/balance/balance.proto

package pb

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

// BalanceServiceClient is the client API for BalanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalanceServiceClient interface {
	Balance(ctx context.Context, in *BalanceServiceRequest, opts ...grpc.CallOption) (*BalanceServiceReply, error)
}

type balanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBalanceServiceClient(cc grpc.ClientConnInterface) BalanceServiceClient {
	return &balanceServiceClient{cc}
}

func (c *balanceServiceClient) Balance(ctx context.Context, in *BalanceServiceRequest, opts ...grpc.CallOption) (*BalanceServiceReply, error) {
	out := new(BalanceServiceReply)
	err := c.cc.Invoke(ctx, "/BalanceService/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalanceServiceServer is the server API for BalanceService service.
// All implementations must embed UnimplementedBalanceServiceServer
// for forward compatibility
type BalanceServiceServer interface {
	Balance(context.Context, *BalanceServiceRequest) (*BalanceServiceReply, error)
	mustEmbedUnimplementedBalanceServiceServer()
}

// UnimplementedBalanceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBalanceServiceServer struct {
}

func (UnimplementedBalanceServiceServer) Balance(context.Context, *BalanceServiceRequest) (*BalanceServiceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedBalanceServiceServer) mustEmbedUnimplementedBalanceServiceServer() {}

// UnsafeBalanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalanceServiceServer will
// result in compilation errors.
type UnsafeBalanceServiceServer interface {
	mustEmbedUnimplementedBalanceServiceServer()
}

func RegisterBalanceServiceServer(s grpc.ServiceRegistrar, srv BalanceServiceServer) {
	s.RegisterService(&BalanceService_ServiceDesc, srv)
}

func _BalanceService_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BalanceServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalanceServiceServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BalanceService/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalanceServiceServer).Balance(ctx, req.(*BalanceServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BalanceService_ServiceDesc is the grpc.ServiceDesc for BalanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BalanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BalanceService",
	HandlerType: (*BalanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Balance",
			Handler:    _BalanceService_Balance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/balance/balance.proto",
}
