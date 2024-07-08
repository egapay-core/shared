// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/core/EganowPaypartnerCallBack.proto

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

// ///messages
type EganowPaypartnerCallBackEmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EganowPaypartnerCallBackEmptyMessage) Reset() {
	*x = EganowPaypartnerCallBackEmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EganowPaypartnerCallBackEmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EganowPaypartnerCallBackEmptyMessage) ProtoMessage() {}

func (x *EganowPaypartnerCallBackEmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EganowPaypartnerCallBackEmptyMessage.ProtoReflect.Descriptor instead.
func (*EganowPaypartnerCallBackEmptyMessage) Descriptor() ([]byte, []int) {
	return file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescGZIP(), []int{0}
}

type EganowPaypartnerCallRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EganowTransRefNo             string `protobuf:"bytes,1,opt,name=eganowTransRefNo,proto3" json:"eganowTransRefNo,omitempty"`
	CountryCode                  string `protobuf:"bytes,2,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	TransactionStatusFromPartner string `protobuf:"bytes,3,opt,name=transactionStatusFromPartner,proto3" json:"transactionStatusFromPartner,omitempty"`
	PartnerSuccessorFailMessage  string `protobuf:"bytes,4,opt,name=partnerSuccessorFailMessage,proto3" json:"partnerSuccessorFailMessage,omitempty"`
}

func (x *EganowPaypartnerCallRequestMessage) Reset() {
	*x = EganowPaypartnerCallRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EganowPaypartnerCallRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EganowPaypartnerCallRequestMessage) ProtoMessage() {}

func (x *EganowPaypartnerCallRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EganowPaypartnerCallRequestMessage.ProtoReflect.Descriptor instead.
func (*EganowPaypartnerCallRequestMessage) Descriptor() ([]byte, []int) {
	return file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescGZIP(), []int{1}
}

func (x *EganowPaypartnerCallRequestMessage) GetEganowTransRefNo() string {
	if x != nil {
		return x.EganowTransRefNo
	}
	return ""
}

func (x *EganowPaypartnerCallRequestMessage) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *EganowPaypartnerCallRequestMessage) GetTransactionStatusFromPartner() string {
	if x != nil {
		return x.TransactionStatusFromPartner
	}
	return ""
}

func (x *EganowPaypartnerCallRequestMessage) GetPartnerSuccessorFailMessage() string {
	if x != nil {
		return x.PartnerSuccessorFailMessage
	}
	return ""
}

var File_eganow_api_core_EganowPaypartnerCallBack_proto protoreflect.FileDescriptor

var file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x22, 0x26, 0x0a, 0x24, 0x45, 0x67,
	0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61,
	0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x22, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79,
	0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x67, 0x61,
	0x6e, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x52, 0x65, 0x66, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x42, 0x0a, 0x1c, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x72, 0x6f, 0x6d,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x46, 0x72, 0x6f, 0x6d, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x1b, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6f, 0x72, 0x46,
	0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1b, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x6f, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xc8, 0x01,
	0x0a, 0x1b, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x53, 0x76, 0x63, 0x12, 0xa8, 0x01,
	0x0a, 0x28, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x50,
	0x61, 0x79, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x2e, 0x45, 0x67, 0x61,
	0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x6c,
	0x6c, 0x42, 0x61, 0x63, 0x6b, 0x2e, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x3e, 0x2e, 0x45, 0x67, 0x61, 0x6e, 0x6f,
	0x77, 0x50, 0x61, 0x79, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x42,
	0x61, 0x63, 0x6b, 0x2e, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x6c, 0x6c, 0x42, 0x61, 0x63, 0x6b, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x5c, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0xaa, 0x02, 0x36,
	0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x47, 0x72, 0x70, 0x63, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x4d, 0x54, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x41, 0x6e, 0x64, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x69, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescOnce sync.Once
	file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescData = file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDesc
)

func file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescGZIP() []byte {
	file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescOnce.Do(func() {
		file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescData)
	})
	return file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDescData
}

var file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eganow_api_core_EganowPaypartnerCallBack_proto_goTypes = []interface{}{
	(*EganowPaypartnerCallBackEmptyMessage)(nil), // 0: EganowPaypartnerCallBack.EganowPaypartnerCallBackEmptyMessage
	(*EganowPaypartnerCallRequestMessage)(nil),   // 1: EganowPaypartnerCallBack.EganowPaypartnerCallRequestMessage
}
var file_eganow_api_core_EganowPaypartnerCallBack_proto_depIdxs = []int32{
	1, // 0: EganowPaypartnerCallBack.EganowPaypartnerCallBackSvc.ConfirmSenderPaypartnerTransactionStatus:input_type -> EganowPaypartnerCallBack.EganowPaypartnerCallRequestMessage
	0, // 1: EganowPaypartnerCallBack.EganowPaypartnerCallBackSvc.ConfirmSenderPaypartnerTransactionStatus:output_type -> EganowPaypartnerCallBack.EganowPaypartnerCallBackEmptyMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eganow_api_core_EganowPaypartnerCallBack_proto_init() }
func file_eganow_api_core_EganowPaypartnerCallBack_proto_init() {
	if File_eganow_api_core_EganowPaypartnerCallBack_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EganowPaypartnerCallBackEmptyMessage); i {
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
		file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EganowPaypartnerCallRequestMessage); i {
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
			RawDescriptor: file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eganow_api_core_EganowPaypartnerCallBack_proto_goTypes,
		DependencyIndexes: file_eganow_api_core_EganowPaypartnerCallBack_proto_depIdxs,
		MessageInfos:      file_eganow_api_core_EganowPaypartnerCallBack_proto_msgTypes,
	}.Build()
	File_eganow_api_core_EganowPaypartnerCallBack_proto = out.File
	file_eganow_api_core_EganowPaypartnerCallBack_proto_rawDesc = nil
	file_eganow_api_core_EganowPaypartnerCallBack_proto_goTypes = nil
	file_eganow_api_core_EganowPaypartnerCallBack_proto_depIdxs = nil
}