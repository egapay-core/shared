// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/payment/payment_service.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_eganow_api_payment_payment_service_proto protoreflect.FileDescriptor

var file_eganow_api_payment_payment_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x65, 0x67, 0x61, 0x6e,
	0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1f,
	0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x65,
	0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xc9, 0x08, 0x0a, 0x15, 0x45, 0x67, 0x61, 0x50,
	0x61, 0x79, 0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x41, 0x70, 0x69, 0x53, 0x76,
	0x63, 0x12, 0x7b, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x45, 0x67, 0x61, 0x50, 0x61, 0x79, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x27, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x95,
	0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x65, 0x67, 0x61, 0x6e,
	0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x79, 0x63, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0xa4, 0x01, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x9c, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f,
	0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x2f, 0x70, 0x61,
	0x79, 0x6f, 0x75, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x8f, 0x01, 0x0a,
	0x13, 0x50, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x65, 0x67, 0x61, 0x6e,
	0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x3a, 0x01, 0x2a, 0x22, 0x22, 0x2f, 0x65, 0x67, 0x61,
	0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x91,
	0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x62,
	0x69, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x65, 0x67, 0x61, 0x6e,
	0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x3a, 0x01, 0x2a, 0x22,
	0x21, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x64, 0x65, 0x62, 0x69, 0x74, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0xae, 0x01, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x2e,
	0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e,
	0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x3a, 0x01,
	0x2a, 0x22, 0x20, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0x74, 0x0a, 0x10, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x53, 0x76, 0x63, 0x12, 0x60, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x6f, 0x6e, 0x65, 0x79, 0x12, 0x28, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x79, 0x0a, 0x0e, 0x4e, 0x61, 0x6d,
	0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x53, 0x76, 0x63, 0x12, 0x67, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x26, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65, 0x67,
	0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eganow_api_payment_payment_service_proto_goTypes = []interface{}{
	(*EgaPayEmpty)(nil),                         // 0: eganow.api.payment.EgaPayEmpty
	(*GetAccountInfoRequest)(nil),               // 1: eganow.api.payment.GetAccountInfoRequest
	(*GetServiceListRequest)(nil),               // 2: eganow.api.payment.GetServiceListRequest
	(*TransferRequest)(nil),                     // 3: eganow.api.payment.TransferRequest
	(*QueryTransferStatusRequest)(nil),          // 4: eganow.api.payment.QueryTransferStatusRequest
	(*MoneyTransferRequest)(nil),                // 5: eganow.api.payment.MoneyTransferRequest
	(*NameEnquiryRequest)(nil),                  // 6: eganow.api.payment.NameEnquiryRequest
	(*AccessTokenResponse)(nil),                 // 7: eganow.api.payment.AccessTokenResponse
	(*GetAccountInfoResponse)(nil),              // 8: eganow.api.payment.GetAccountInfoResponse
	(*GetServiceListResponse)(nil),              // 9: eganow.api.payment.GetServiceListResponse
	(*TransferResponse)(nil),                    // 10: eganow.api.payment.TransferResponse
	(*QueryTransferStatusTransferResponse)(nil), // 11: eganow.api.payment.QueryTransferStatusTransferResponse
	(*MoneyTransferResponse)(nil),               // 12: eganow.api.payment.MoneyTransferResponse
	(*NameEnquiryResponse)(nil),                 // 13: eganow.api.payment.NameEnquiryResponse
}
var file_eganow_api_payment_payment_service_proto_depIdxs = []int32{
	0,  // 0: eganow.api.payment.EgaPayDeveloperApiSvc.GetAccessToken:input_type -> eganow.api.payment.EgaPayEmpty
	1,  // 1: eganow.api.payment.EgaPayDeveloperApiSvc.GetAccountHolderInfo:input_type -> eganow.api.payment.GetAccountInfoRequest
	2,  // 2: eganow.api.payment.EgaPayDeveloperApiSvc.GetCollectionServiceList:input_type -> eganow.api.payment.GetServiceListRequest
	2,  // 3: eganow.api.payment.EgaPayDeveloperApiSvc.GetPayoutServiceList:input_type -> eganow.api.payment.GetServiceListRequest
	3,  // 4: eganow.api.payment.EgaPayDeveloperApiSvc.PayoutCreditAccount:input_type -> eganow.api.payment.TransferRequest
	3,  // 5: eganow.api.payment.EgaPayDeveloperApiSvc.CollectionDebitAccount:input_type -> eganow.api.payment.TransferRequest
	4,  // 6: eganow.api.payment.EgaPayDeveloperApiSvc.TransactionStatusQuery:input_type -> eganow.api.payment.QueryTransferStatusRequest
	5,  // 7: eganow.api.payment.MoneyTransferSvc.SendMoney:input_type -> eganow.api.payment.MoneyTransferRequest
	6,  // 8: eganow.api.payment.NameEnquirySvc.GetAccountHolderName:input_type -> eganow.api.payment.NameEnquiryRequest
	7,  // 9: eganow.api.payment.EgaPayDeveloperApiSvc.GetAccessToken:output_type -> eganow.api.payment.AccessTokenResponse
	8,  // 10: eganow.api.payment.EgaPayDeveloperApiSvc.GetAccountHolderInfo:output_type -> eganow.api.payment.GetAccountInfoResponse
	9,  // 11: eganow.api.payment.EgaPayDeveloperApiSvc.GetCollectionServiceList:output_type -> eganow.api.payment.GetServiceListResponse
	9,  // 12: eganow.api.payment.EgaPayDeveloperApiSvc.GetPayoutServiceList:output_type -> eganow.api.payment.GetServiceListResponse
	10, // 13: eganow.api.payment.EgaPayDeveloperApiSvc.PayoutCreditAccount:output_type -> eganow.api.payment.TransferResponse
	10, // 14: eganow.api.payment.EgaPayDeveloperApiSvc.CollectionDebitAccount:output_type -> eganow.api.payment.TransferResponse
	11, // 15: eganow.api.payment.EgaPayDeveloperApiSvc.TransactionStatusQuery:output_type -> eganow.api.payment.QueryTransferStatusTransferResponse
	12, // 16: eganow.api.payment.MoneyTransferSvc.SendMoney:output_type -> eganow.api.payment.MoneyTransferResponse
	13, // 17: eganow.api.payment.NameEnquirySvc.GetAccountHolderName:output_type -> eganow.api.payment.NameEnquiryResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_eganow_api_payment_payment_service_proto_init() }
func file_eganow_api_payment_payment_service_proto_init() {
	if File_eganow_api_payment_payment_service_proto != nil {
		return
	}
	file_eganow_api_payment_common_proto_init()
	file_eganow_api_payment_payment_payload_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eganow_api_payment_payment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_eganow_api_payment_payment_service_proto_goTypes,
		DependencyIndexes: file_eganow_api_payment_payment_service_proto_depIdxs,
	}.Build()
	File_eganow_api_payment_payment_service_proto = out.File
	file_eganow_api_payment_payment_service_proto_rawDesc = nil
	file_eganow_api_payment_payment_service_proto_goTypes = nil
	file_eganow_api_payment_payment_service_proto_depIdxs = nil
}