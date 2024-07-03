// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: eganow/api/common/otp_service.proto

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
	OtpSvc_SendOTP_FullMethodName   = "/eganow.api.common.OtpSvc/SendOTP"
	OtpSvc_VerifyOTP_FullMethodName = "/eganow.api.common.OtpSvc/VerifyOTP"
)

// OtpSvcClient is the client API for OtpSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OtpSvcClient interface {
	// send OTP to a given phone number or email
	SendOTP(ctx context.Context, in *SendOTPRequest, opts ...grpc.CallOption) (*CommonStringValue, error)
	// verify OTP for a given phone number or email
	VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*CommonStringValue, error)
}

type otpSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewOtpSvcClient(cc grpc.ClientConnInterface) OtpSvcClient {
	return &otpSvcClient{cc}
}

func (c *otpSvcClient) SendOTP(ctx context.Context, in *SendOTPRequest, opts ...grpc.CallOption) (*CommonStringValue, error) {
	out := new(CommonStringValue)
	err := c.cc.Invoke(ctx, OtpSvc_SendOTP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otpSvcClient) VerifyOTP(ctx context.Context, in *VerifyOTPRequest, opts ...grpc.CallOption) (*CommonStringValue, error) {
	out := new(CommonStringValue)
	err := c.cc.Invoke(ctx, OtpSvc_VerifyOTP_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OtpSvcServer is the server API for OtpSvc service.
// All implementations must embed UnimplementedOtpSvcServer
// for forward compatibility
type OtpSvcServer interface {
	// send OTP to a given phone number or email
	SendOTP(context.Context, *SendOTPRequest) (*CommonStringValue, error)
	// verify OTP for a given phone number or email
	VerifyOTP(context.Context, *VerifyOTPRequest) (*CommonStringValue, error)
	mustEmbedUnimplementedOtpSvcServer()
}

// UnimplementedOtpSvcServer must be embedded to have forward compatible implementations.
type UnimplementedOtpSvcServer struct {
}

func (UnimplementedOtpSvcServer) SendOTP(context.Context, *SendOTPRequest) (*CommonStringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOTP not implemented")
}
func (UnimplementedOtpSvcServer) VerifyOTP(context.Context, *VerifyOTPRequest) (*CommonStringValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyOTP not implemented")
}
func (UnimplementedOtpSvcServer) mustEmbedUnimplementedOtpSvcServer() {}

// UnsafeOtpSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OtpSvcServer will
// result in compilation errors.
type UnsafeOtpSvcServer interface {
	mustEmbedUnimplementedOtpSvcServer()
}

func RegisterOtpSvcServer(s grpc.ServiceRegistrar, srv OtpSvcServer) {
	s.RegisterService(&OtpSvc_ServiceDesc, srv)
}

func _OtpSvc_SendOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpSvcServer).SendOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpSvc_SendOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpSvcServer).SendOTP(ctx, req.(*SendOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OtpSvc_VerifyOTP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyOTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtpSvcServer).VerifyOTP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OtpSvc_VerifyOTP_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtpSvcServer).VerifyOTP(ctx, req.(*VerifyOTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OtpSvc_ServiceDesc is the grpc.ServiceDesc for OtpSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OtpSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eganow.api.common.OtpSvc",
	HandlerType: (*OtpSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendOTP",
			Handler:    _OtpSvc_SendOTP_Handler,
		},
		{
			MethodName: "VerifyOTP",
			Handler:    _OtpSvc_VerifyOTP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "eganow/api/common/otp_service.proto",
}
