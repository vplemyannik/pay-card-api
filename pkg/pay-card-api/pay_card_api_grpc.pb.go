// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pay_card_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PayCardApiServiceClient is the client API for PayCardApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayCardApiServiceClient interface {
	CreateCard(ctx context.Context, in *CreateCardV1Request, opts ...grpc.CallOption) (*CreateCardV1Response, error)
	RemoveCard(ctx context.Context, in *RemoveCardV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DescribeCard(ctx context.Context, in *DescribeCardV1Request, opts ...grpc.CallOption) (*Card, error)
	ListCard(ctx context.Context, in *ListCardV1Request, opts ...grpc.CallOption) (*ListCardV1Response, error)
}

type payCardApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayCardApiServiceClient(cc grpc.ClientConnInterface) PayCardApiServiceClient {
	return &payCardApiServiceClient{cc}
}

func (c *payCardApiServiceClient) CreateCard(ctx context.Context, in *CreateCardV1Request, opts ...grpc.CallOption) (*CreateCardV1Response, error) {
	out := new(CreateCardV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.pay_card_api.v1.PayCardApiService/CreateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payCardApiServiceClient) RemoveCard(ctx context.Context, in *RemoveCardV1Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ozonmp.pay_card_api.v1.PayCardApiService/RemoveCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payCardApiServiceClient) DescribeCard(ctx context.Context, in *DescribeCardV1Request, opts ...grpc.CallOption) (*Card, error) {
	out := new(Card)
	err := c.cc.Invoke(ctx, "/ozonmp.pay_card_api.v1.PayCardApiService/DescribeCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payCardApiServiceClient) ListCard(ctx context.Context, in *ListCardV1Request, opts ...grpc.CallOption) (*ListCardV1Response, error) {
	out := new(ListCardV1Response)
	err := c.cc.Invoke(ctx, "/ozonmp.pay_card_api.v1.PayCardApiService/ListCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayCardApiServiceServer is the server API for PayCardApiService service.
// All implementations must embed UnimplementedPayCardApiServiceServer
// for forward compatibility
type PayCardApiServiceServer interface {
	CreateCard(context.Context, *CreateCardV1Request) (*CreateCardV1Response, error)
	RemoveCard(context.Context, *RemoveCardV1Request) (*emptypb.Empty, error)
	DescribeCard(context.Context, *DescribeCardV1Request) (*Card, error)
	ListCard(context.Context, *ListCardV1Request) (*ListCardV1Response, error)
	mustEmbedUnimplementedPayCardApiServiceServer()
}

// UnimplementedPayCardApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayCardApiServiceServer struct {
}

func (UnimplementedPayCardApiServiceServer) CreateCard(context.Context, *CreateCardV1Request) (*CreateCardV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedPayCardApiServiceServer) RemoveCard(context.Context, *RemoveCardV1Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCard not implemented")
}
func (UnimplementedPayCardApiServiceServer) DescribeCard(context.Context, *DescribeCardV1Request) (*Card, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeCard not implemented")
}
func (UnimplementedPayCardApiServiceServer) ListCard(context.Context, *ListCardV1Request) (*ListCardV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCard not implemented")
}
func (UnimplementedPayCardApiServiceServer) mustEmbedUnimplementedPayCardApiServiceServer() {}

// UnsafePayCardApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayCardApiServiceServer will
// result in compilation errors.
type UnsafePayCardApiServiceServer interface {
	mustEmbedUnimplementedPayCardApiServiceServer()
}

func RegisterPayCardApiServiceServer(s grpc.ServiceRegistrar, srv PayCardApiServiceServer) {
	s.RegisterService(&PayCardApiService_ServiceDesc, srv)
}

func _PayCardApiService_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayCardApiServiceServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.pay_card_api.v1.PayCardApiService/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayCardApiServiceServer).CreateCard(ctx, req.(*CreateCardV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayCardApiService_RemoveCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCardV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayCardApiServiceServer).RemoveCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.pay_card_api.v1.PayCardApiService/RemoveCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayCardApiServiceServer).RemoveCard(ctx, req.(*RemoveCardV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayCardApiService_DescribeCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeCardV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayCardApiServiceServer).DescribeCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.pay_card_api.v1.PayCardApiService/DescribeCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayCardApiServiceServer).DescribeCard(ctx, req.(*DescribeCardV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayCardApiService_ListCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCardV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayCardApiServiceServer).ListCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ozonmp.pay_card_api.v1.PayCardApiService/ListCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayCardApiServiceServer).ListCard(ctx, req.(*ListCardV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PayCardApiService_ServiceDesc is the grpc.ServiceDesc for PayCardApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayCardApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ozonmp.pay_card_api.v1.PayCardApiService",
	HandlerType: (*PayCardApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCard",
			Handler:    _PayCardApiService_CreateCard_Handler,
		},
		{
			MethodName: "RemoveCard",
			Handler:    _PayCardApiService_RemoveCard_Handler,
		},
		{
			MethodName: "DescribeCard",
			Handler:    _PayCardApiService_DescribeCard_Handler,
		},
		{
			MethodName: "ListCard",
			Handler:    _PayCardApiService_ListCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ozonmp/pay_card_api/v1/pay_card_api.proto",
}
