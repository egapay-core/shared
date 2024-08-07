// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/customer/common.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CustomerEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CustomerEmpty) Reset() {
	*x = CustomerEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_customer_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerEmpty) ProtoMessage() {}

func (x *CustomerEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_customer_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerEmpty.ProtoReflect.Descriptor instead.
func (*CustomerEmpty) Descriptor() ([]byte, []int) {
	return file_eganow_api_customer_common_proto_rawDescGZIP(), []int{0}
}

type CustomerStringValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CustomerStringValue) Reset() {
	*x = CustomerStringValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_customer_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerStringValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerStringValue) ProtoMessage() {}

func (x *CustomerStringValue) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_customer_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerStringValue.ProtoReflect.Descriptor instead.
func (*CustomerStringValue) Descriptor() ([]byte, []int) {
	return file_eganow_api_customer_common_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerStringValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_eganow_api_customer_common_proto protoreflect.FileDescriptor

var file_eganow_api_customer_common_proto_rawDesc = []byte{
	0x0a, 0x20, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eganow_api_customer_common_proto_rawDescOnce sync.Once
	file_eganow_api_customer_common_proto_rawDescData = file_eganow_api_customer_common_proto_rawDesc
)

func file_eganow_api_customer_common_proto_rawDescGZIP() []byte {
	file_eganow_api_customer_common_proto_rawDescOnce.Do(func() {
		file_eganow_api_customer_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_customer_common_proto_rawDescData)
	})
	return file_eganow_api_customer_common_proto_rawDescData
}

var file_eganow_api_customer_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eganow_api_customer_common_proto_goTypes = []interface{}{
	(*CustomerEmpty)(nil),       // 0: eganow.api.customer.CustomerEmpty
	(*CustomerStringValue)(nil), // 1: eganow.api.customer.CustomerStringValue
}
var file_eganow_api_customer_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eganow_api_customer_common_proto_init() }
func file_eganow_api_customer_common_proto_init() {
	if File_eganow_api_customer_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_customer_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerEmpty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_eganow_api_customer_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerStringValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eganow_api_customer_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eganow_api_customer_common_proto_goTypes,
		DependencyIndexes: file_eganow_api_customer_common_proto_depIdxs,
		MessageInfos:      file_eganow_api_customer_common_proto_msgTypes,
	}.Build()
	File_eganow_api_customer_common_proto = out.File
	file_eganow_api_customer_common_proto_rawDesc = nil
	file_eganow_api_customer_common_proto_goTypes = nil
	file_eganow_api_customer_common_proto_depIdxs = nil
}
