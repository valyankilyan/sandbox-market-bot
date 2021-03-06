// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: api/myinvestapi/account.proto

package myinvestapi

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MyInvestServiceClient is the client API for MyInvestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyInvestServiceClient interface {
	PayIn(ctx context.Context, in *PayinRequest, opts ...grpc.CallOption) (*Quotation, error)
	GetCurrencies(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*CurrencyList, error)
	GetAllCurrencies(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CurrencyList, error)
}

type myInvestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMyInvestServiceClient(cc grpc.ClientConnInterface) MyInvestServiceClient {
	return &myInvestServiceClient{cc}
}

func (c *myInvestServiceClient) PayIn(ctx context.Context, in *PayinRequest, opts ...grpc.CallOption) (*Quotation, error) {
	out := new(Quotation)
	err := c.cc.Invoke(ctx, "/myinvestapi.MyInvestService/PayIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myInvestServiceClient) GetCurrencies(ctx context.Context, in *CurrencyRequest, opts ...grpc.CallOption) (*CurrencyList, error) {
	out := new(CurrencyList)
	err := c.cc.Invoke(ctx, "/myinvestapi.MyInvestService/GetCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myInvestServiceClient) GetAllCurrencies(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CurrencyList, error) {
	out := new(CurrencyList)
	err := c.cc.Invoke(ctx, "/myinvestapi.MyInvestService/GetAllCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyInvestServiceServer is the server API for MyInvestService service.
// All implementations must embed UnimplementedMyInvestServiceServer
// for forward compatibility
type MyInvestServiceServer interface {
	PayIn(context.Context, *PayinRequest) (*Quotation, error)
	GetCurrencies(context.Context, *CurrencyRequest) (*CurrencyList, error)
	GetAllCurrencies(context.Context, *empty.Empty) (*CurrencyList, error)
	mustEmbedUnimplementedMyInvestServiceServer()
}

// UnimplementedMyInvestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMyInvestServiceServer struct {
}

func (UnimplementedMyInvestServiceServer) PayIn(context.Context, *PayinRequest) (*Quotation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayIn not implemented")
}
func (UnimplementedMyInvestServiceServer) GetCurrencies(context.Context, *CurrencyRequest) (*CurrencyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrencies not implemented")
}
func (UnimplementedMyInvestServiceServer) GetAllCurrencies(context.Context, *empty.Empty) (*CurrencyList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCurrencies not implemented")
}
func (UnimplementedMyInvestServiceServer) mustEmbedUnimplementedMyInvestServiceServer() {}

// UnsafeMyInvestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyInvestServiceServer will
// result in compilation errors.
type UnsafeMyInvestServiceServer interface {
	mustEmbedUnimplementedMyInvestServiceServer()
}

func RegisterMyInvestServiceServer(s grpc.ServiceRegistrar, srv MyInvestServiceServer) {
	s.RegisterService(&MyInvestService_ServiceDesc, srv)
}

func _MyInvestService_PayIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyInvestServiceServer).PayIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myinvestapi.MyInvestService/PayIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyInvestServiceServer).PayIn(ctx, req.(*PayinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyInvestService_GetCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyInvestServiceServer).GetCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myinvestapi.MyInvestService/GetCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyInvestServiceServer).GetCurrencies(ctx, req.(*CurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyInvestService_GetAllCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyInvestServiceServer).GetAllCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myinvestapi.MyInvestService/GetAllCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyInvestServiceServer).GetAllCurrencies(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MyInvestService_ServiceDesc is the grpc.ServiceDesc for MyInvestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyInvestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "myinvestapi.MyInvestService",
	HandlerType: (*MyInvestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PayIn",
			Handler:    _MyInvestService_PayIn_Handler,
		},
		{
			MethodName: "GetCurrencies",
			Handler:    _MyInvestService_GetCurrencies_Handler,
		},
		{
			MethodName: "GetAllCurrencies",
			Handler:    _MyInvestService_GetAllCurrencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/myinvestapi/account.proto",
}
