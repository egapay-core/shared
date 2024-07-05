// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: eganow/api/payment/payment_service.proto

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

const (
	CoreMoneyTransferSvc_CollectMoney_FullMethodName = "/eganow.api.payment.CoreMoneyTransferSvc/CollectMoney"
	CoreMoneyTransferSvc_PayoutMoney_FullMethodName  = "/eganow.api.payment.CoreMoneyTransferSvc/PayoutMoney"
)

// CoreMoneyTransferSvcClient is the client API for CoreMoneyTransferSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreMoneyTransferSvcClient interface {
	CollectMoney(ctx context.Context, in *PaymentMoneyTransferRequest, opts ...grpc.CallOption) (*PaymentMoneyTransferResponse, error)
	PayoutMoney(ctx context.Context, in *PaymentMoneyTransferRequest, opts ...grpc.CallOption) (*PaymentMoneyTransferResponse, error)
}

type coreMoneyTransferSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreMoneyTransferSvcClient(cc grpc.ClientConnInterface) CoreMoneyTransferSvcClient {
	return &coreMoneyTransferSvcClient{cc}
}

func (c *coreMoneyTransferSvcClient) CollectMoney(ctx context.Context, in *PaymentMoneyTransferRequest, opts ...grpc.CallOption) (*PaymentMoneyTransferResponse, error) {
	out := new(PaymentMoneyTransferResponse)
	err := c.cc.Invoke(ctx, CoreMoneyTransferSvc_CollectMoney_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coreMoneyTransferSvcClient) PayoutMoney(ctx context.Context, in *PaymentMoneyTransferRequest, opts ...grpc.CallOption) (*PaymentMoneyTransferResponse, error) {
	out := new(PaymentMoneyTransferResponse)
	err := c.cc.Invoke(ctx, CoreMoneyTransferSvc_PayoutMoney_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreMoneyTransferSvcServer is the server API for CoreMoneyTransferSvc service.
// All implementations must embed UnimplementedCoreMoneyTransferSvcServer
// for forward compatibility
type CoreMoneyTransferSvcServer interface {
	CollectMoney(context.Context, *PaymentMoneyTransferRequest) (*PaymentMoneyTransferResponse, error)
	PayoutMoney(context.Context, *PaymentMoneyTransferRequest) (*PaymentMoneyTransferResponse, error)
	mustEmbedUnimplementedCoreMoneyTransferSvcServer()
}

// UnimplementedCoreMoneyTransferSvcServer must be embedded to have forward compatible implementations.
type UnimplementedCoreMoneyTransferSvcServer struct {
}

func (UnimplementedCoreMoneyTransferSvcServer) CollectMoney(context.Context, *PaymentMoneyTransferRequest) (*PaymentMoneyTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectMoney not implemented")
}
func (UnimplementedCoreMoneyTransferSvcServer) PayoutMoney(context.Context, *PaymentMoneyTransferRequest) (*PaymentMoneyTransferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayoutMoney not implemented")
}
func (UnimplementedCoreMoneyTransferSvcServer) mustEmbedUnimplementedCoreMoneyTransferSvcServer() {}

// UnsafeCoreMoneyTransferSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreMoneyTransferSvcServer will
// result in compilation errors.
type UnsafeCoreMoneyTransferSvcServer interface {
	mustEmbedUnimplementedCoreMoneyTransferSvcServer()
}

func RegisterCoreMoneyTransferSvcServer(s grpc.ServiceRegistrar, srv CoreMoneyTransferSvcServer) {
	s.RegisterService(&CoreMoneyTransferSvc_ServiceDesc, srv)
}

func _CoreMoneyTransferSvc_CollectMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMoneyTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreMoneyTransferSvcServer).CollectMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoreMoneyTransferSvc_CollectMoney_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreMoneyTransferSvcServer).CollectMoney(ctx, req.(*PaymentMoneyTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoreMoneyTransferSvc_PayoutMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentMoneyTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreMoneyTransferSvcServer).PayoutMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoreMoneyTransferSvc_PayoutMoney_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreMoneyTransferSvcServer).PayoutMoney(ctx, req.(*PaymentMoneyTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoreMoneyTransferSvc_ServiceDesc is the grpc.ServiceDesc for CoreMoneyTransferSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreMoneyTransferSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eganow.api.payment.CoreMoneyTransferSvc",
	HandlerType: (*CoreMoneyTransferSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CollectMoney",
			Handler:    _CoreMoneyTransferSvc_CollectMoney_Handler,
		},
		{
			MethodName: "PayoutMoney",
			Handler:    _CoreMoneyTransferSvc_PayoutMoney_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eganow/api/payment/payment_service.proto",
}

const (
	CoreNameEnquirySvc_GetAccountHolderName_FullMethodName = "/eganow.api.payment.CoreNameEnquirySvc/GetAccountHolderName"
)

// CoreNameEnquirySvcClient is the client API for CoreNameEnquirySvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoreNameEnquirySvcClient interface {
	GetAccountHolderName(ctx context.Context, in *PaymentNameEnquiryRequest, opts ...grpc.CallOption) (*PaymentNameEnquiryResponse, error)
}

type coreNameEnquirySvcClient struct {
	cc grpc.ClientConnInterface
}

func NewCoreNameEnquirySvcClient(cc grpc.ClientConnInterface) CoreNameEnquirySvcClient {
	return &coreNameEnquirySvcClient{cc}
}

func (c *coreNameEnquirySvcClient) GetAccountHolderName(ctx context.Context, in *PaymentNameEnquiryRequest, opts ...grpc.CallOption) (*PaymentNameEnquiryResponse, error) {
	out := new(PaymentNameEnquiryResponse)
	err := c.cc.Invoke(ctx, CoreNameEnquirySvc_GetAccountHolderName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoreNameEnquirySvcServer is the server API for CoreNameEnquirySvc service.
// All implementations must embed UnimplementedCoreNameEnquirySvcServer
// for forward compatibility
type CoreNameEnquirySvcServer interface {
	GetAccountHolderName(context.Context, *PaymentNameEnquiryRequest) (*PaymentNameEnquiryResponse, error)
	mustEmbedUnimplementedCoreNameEnquirySvcServer()
}

// UnimplementedCoreNameEnquirySvcServer must be embedded to have forward compatible implementations.
type UnimplementedCoreNameEnquirySvcServer struct {
}

func (UnimplementedCoreNameEnquirySvcServer) GetAccountHolderName(context.Context, *PaymentNameEnquiryRequest) (*PaymentNameEnquiryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountHolderName not implemented")
}
func (UnimplementedCoreNameEnquirySvcServer) mustEmbedUnimplementedCoreNameEnquirySvcServer() {}

// UnsafeCoreNameEnquirySvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoreNameEnquirySvcServer will
// result in compilation errors.
type UnsafeCoreNameEnquirySvcServer interface {
	mustEmbedUnimplementedCoreNameEnquirySvcServer()
}

func RegisterCoreNameEnquirySvcServer(s grpc.ServiceRegistrar, srv CoreNameEnquirySvcServer) {
	s.RegisterService(&CoreNameEnquirySvc_ServiceDesc, srv)
}

func _CoreNameEnquirySvc_GetAccountHolderName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaymentNameEnquiryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoreNameEnquirySvcServer).GetAccountHolderName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CoreNameEnquirySvc_GetAccountHolderName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoreNameEnquirySvcServer).GetAccountHolderName(ctx, req.(*PaymentNameEnquiryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoreNameEnquirySvc_ServiceDesc is the grpc.ServiceDesc for CoreNameEnquirySvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoreNameEnquirySvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eganow.api.payment.CoreNameEnquirySvc",
	HandlerType: (*CoreNameEnquirySvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccountHolderName",
			Handler:    _CoreNameEnquirySvc_GetAccountHolderName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eganow/api/payment/payment_service.proto",
}
