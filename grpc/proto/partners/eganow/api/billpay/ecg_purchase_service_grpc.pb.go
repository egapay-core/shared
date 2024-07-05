// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: eganow/api/billpay/ecg_purchase_service.proto

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
	ElectricityCreditPurchaseSvc_QueryMeterDetailsWithPhoneNumber_FullMethodName = "/eganow.api.billpay.ElectricityCreditPurchaseSvc/QueryMeterDetailsWithPhoneNumber"
	ElectricityCreditPurchaseSvc_Purchase_FullMethodName                         = "/eganow.api.billpay.ElectricityCreditPurchaseSvc/Purchase"
	ElectricityCreditPurchaseSvc_HubtelUtilityCallback_FullMethodName            = "/eganow.api.billpay.ElectricityCreditPurchaseSvc/HubtelUtilityCallback"
)

// ElectricityCreditPurchaseSvcClient is the client API for ElectricityCreditPurchaseSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ElectricityCreditPurchaseSvcClient interface {
	QueryMeterDetailsWithPhoneNumber(ctx context.Context, in *BillPayStringValue, opts ...grpc.CallOption) (*QueryMeterDetailsResponse, error)
	Purchase(ctx context.Context, in *PurchaseElectricityCreditRequest, opts ...grpc.CallOption) (*BillPayStringValue, error)
	// Callback is called by the partner to notify us of the status of the transaction
	HubtelUtilityCallback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*BillPayEmpty, error)
}

type electricityCreditPurchaseSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewElectricityCreditPurchaseSvcClient(cc grpc.ClientConnInterface) ElectricityCreditPurchaseSvcClient {
	return &electricityCreditPurchaseSvcClient{cc}
}

func (c *electricityCreditPurchaseSvcClient) QueryMeterDetailsWithPhoneNumber(ctx context.Context, in *BillPayStringValue, opts ...grpc.CallOption) (*QueryMeterDetailsResponse, error) {
	out := new(QueryMeterDetailsResponse)
	err := c.cc.Invoke(ctx, ElectricityCreditPurchaseSvc_QueryMeterDetailsWithPhoneNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electricityCreditPurchaseSvcClient) Purchase(ctx context.Context, in *PurchaseElectricityCreditRequest, opts ...grpc.CallOption) (*BillPayStringValue, error) {
	out := new(BillPayStringValue)
	err := c.cc.Invoke(ctx, ElectricityCreditPurchaseSvc_Purchase_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *electricityCreditPurchaseSvcClient) HubtelUtilityCallback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*BillPayEmpty, error) {
	out := new(BillPayEmpty)
	err := c.cc.Invoke(ctx, ElectricityCreditPurchaseSvc_HubtelUtilityCallback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElectricityCreditPurchaseSvcServer is the server API for ElectricityCreditPurchaseSvc service.
// All implementations must embed UnimplementedElectricityCreditPurchaseSvcServer
// for forward compatibility
type ElectricityCreditPurchaseSvcServer interface {
	QueryMeterDetailsWithPhoneNumber(context.Context, *BillPayStringValue) (*QueryMeterDetailsResponse, error)
	Purchase(context.Context, *PurchaseElectricityCreditRequest) (*BillPayStringValue, error)
	// Callback is called by the partner to notify us of the status of the transaction
	HubtelUtilityCallback(context.Context, *CallbackRequest) (*BillPayEmpty, error)
	mustEmbedUnimplementedElectricityCreditPurchaseSvcServer()
}

// UnimplementedElectricityCreditPurchaseSvcServer must be embedded to have forward compatible implementations.
type UnimplementedElectricityCreditPurchaseSvcServer struct {
}

func (UnimplementedElectricityCreditPurchaseSvcServer) QueryMeterDetailsWithPhoneNumber(context.Context, *BillPayStringValue) (*QueryMeterDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMeterDetailsWithPhoneNumber not implemented")
}
func (UnimplementedElectricityCreditPurchaseSvcServer) Purchase(context.Context, *PurchaseElectricityCreditRequest) (*BillPayStringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Purchase not implemented")
}
func (UnimplementedElectricityCreditPurchaseSvcServer) HubtelUtilityCallback(context.Context, *CallbackRequest) (*BillPayEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HubtelUtilityCallback not implemented")
}
func (UnimplementedElectricityCreditPurchaseSvcServer) mustEmbedUnimplementedElectricityCreditPurchaseSvcServer() {
}

// UnsafeElectricityCreditPurchaseSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ElectricityCreditPurchaseSvcServer will
// result in compilation errors.
type UnsafeElectricityCreditPurchaseSvcServer interface {
	mustEmbedUnimplementedElectricityCreditPurchaseSvcServer()
}

func RegisterElectricityCreditPurchaseSvcServer(s grpc.ServiceRegistrar, srv ElectricityCreditPurchaseSvcServer) {
	s.RegisterService(&ElectricityCreditPurchaseSvc_ServiceDesc, srv)
}

func _ElectricityCreditPurchaseSvc_QueryMeterDetailsWithPhoneNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillPayStringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectricityCreditPurchaseSvcServer).QueryMeterDetailsWithPhoneNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectricityCreditPurchaseSvc_QueryMeterDetailsWithPhoneNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectricityCreditPurchaseSvcServer).QueryMeterDetailsWithPhoneNumber(ctx, req.(*BillPayStringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectricityCreditPurchaseSvc_Purchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseElectricityCreditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectricityCreditPurchaseSvcServer).Purchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectricityCreditPurchaseSvc_Purchase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectricityCreditPurchaseSvcServer).Purchase(ctx, req.(*PurchaseElectricityCreditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElectricityCreditPurchaseSvc_HubtelUtilityCallback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElectricityCreditPurchaseSvcServer).HubtelUtilityCallback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ElectricityCreditPurchaseSvc_HubtelUtilityCallback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElectricityCreditPurchaseSvcServer).HubtelUtilityCallback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ElectricityCreditPurchaseSvc_ServiceDesc is the grpc.ServiceDesc for ElectricityCreditPurchaseSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ElectricityCreditPurchaseSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eganow.api.billpay.ElectricityCreditPurchaseSvc",
	HandlerType: (*ElectricityCreditPurchaseSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryMeterDetailsWithPhoneNumber",
			Handler:    _ElectricityCreditPurchaseSvc_QueryMeterDetailsWithPhoneNumber_Handler,
		},
		{
			MethodName: "Purchase",
			Handler:    _ElectricityCreditPurchaseSvc_Purchase_Handler,
		},
		{
			MethodName: "HubtelUtilityCallback",
			Handler:    _ElectricityCreditPurchaseSvc_HubtelUtilityCallback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eganow/api/billpay/ecg_purchase_service.proto",
}
