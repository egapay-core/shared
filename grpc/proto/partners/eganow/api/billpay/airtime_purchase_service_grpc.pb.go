// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: eganow/api/billpay/airtime_purchase_service.proto

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
	AirtimePurchaseSvc_Purchase_FullMethodName              = "/eganow.api.billpay.AirtimePurchaseSvc/Purchase"
	AirtimePurchaseSvc_HubtelAirtimeCallback_FullMethodName = "/eganow.api.billpay.AirtimePurchaseSvc/HubtelAirtimeCallback"
)

// AirtimePurchaseSvcClient is the client API for AirtimePurchaseSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AirtimePurchaseSvcClient interface {
	// Purchase is called to initiate a airtime purchase transaction
	Purchase(ctx context.Context, in *AirtimePurchaseRequest, opts ...grpc.CallOption) (*BillPayStringValue, error)
	// Callback is called by the partner to notify us of the status of the transaction
	HubtelAirtimeCallback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*BillPayEmpty, error)
}

type airtimePurchaseSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewAirtimePurchaseSvcClient(cc grpc.ClientConnInterface) AirtimePurchaseSvcClient {
	return &airtimePurchaseSvcClient{cc}
}

func (c *airtimePurchaseSvcClient) Purchase(ctx context.Context, in *AirtimePurchaseRequest, opts ...grpc.CallOption) (*BillPayStringValue, error) {
	out := new(BillPayStringValue)
	err := c.cc.Invoke(ctx, AirtimePurchaseSvc_Purchase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtimePurchaseSvcClient) HubtelAirtimeCallback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*BillPayEmpty, error) {
	out := new(BillPayEmpty)
	err := c.cc.Invoke(ctx, AirtimePurchaseSvc_HubtelAirtimeCallback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AirtimePurchaseSvcServer is the server API for AirtimePurchaseSvc service.
// All implementations must embed UnimplementedAirtimePurchaseSvcServer
// for forward compatibility
type AirtimePurchaseSvcServer interface {
	// Purchase is called to initiate a airtime purchase transaction
	Purchase(context.Context, *AirtimePurchaseRequest) (*BillPayStringValue, error)
	// Callback is called by the partner to notify us of the status of the transaction
	HubtelAirtimeCallback(context.Context, *CallbackRequest) (*BillPayEmpty, error)
	mustEmbedUnimplementedAirtimePurchaseSvcServer()
}

// UnimplementedAirtimePurchaseSvcServer must be embedded to have forward compatible implementations.
type UnimplementedAirtimePurchaseSvcServer struct {
}

func (UnimplementedAirtimePurchaseSvcServer) Purchase(context.Context, *AirtimePurchaseRequest) (*BillPayStringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purchase not implemented")
}
func (UnimplementedAirtimePurchaseSvcServer) HubtelAirtimeCallback(context.Context, *CallbackRequest) (*BillPayEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HubtelAirtimeCallback not implemented")
}
func (UnimplementedAirtimePurchaseSvcServer) mustEmbedUnimplementedAirtimePurchaseSvcServer() {}

// UnsafeAirtimePurchaseSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AirtimePurchaseSvcServer will
// result in compilation errors.
type UnsafeAirtimePurchaseSvcServer interface {
	mustEmbedUnimplementedAirtimePurchaseSvcServer()
}

func RegisterAirtimePurchaseSvcServer(s grpc.ServiceRegistrar, srv AirtimePurchaseSvcServer) {
	s.RegisterService(&AirtimePurchaseSvc_ServiceDesc, srv)
}

func _AirtimePurchaseSvc_Purchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirtimePurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirtimePurchaseSvcServer).Purchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirtimePurchaseSvc_Purchase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirtimePurchaseSvcServer).Purchase(ctx, req.(*AirtimePurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirtimePurchaseSvc_HubtelAirtimeCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirtimePurchaseSvcServer).HubtelAirtimeCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirtimePurchaseSvc_HubtelAirtimeCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirtimePurchaseSvcServer).HubtelAirtimeCallback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AirtimePurchaseSvc_ServiceDesc is the grpc.ServiceDesc for AirtimePurchaseSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AirtimePurchaseSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eganow.api.billpay.AirtimePurchaseSvc",
	HandlerType: (*AirtimePurchaseSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Purchase",
			Handler:    _AirtimePurchaseSvc_Purchase_Handler,
		},
		{
			MethodName: "HubtelAirtimeCallback",
			Handler:    _AirtimePurchaseSvc_HubtelAirtimeCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eganow/api/billpay/airtime_purchase_service.proto",
}
