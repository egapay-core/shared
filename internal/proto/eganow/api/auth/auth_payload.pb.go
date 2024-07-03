// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/auth/auth_payload.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserGuid string `protobuf:"bytes,1,opt,name=user_guid,json=userGuid,proto3" json:"user_guid,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"` // email or mobile number
}

func (x *TokenRequest) Reset() {
	*x = TokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_auth_auth_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenRequest) ProtoMessage() {}

func (x *TokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_auth_auth_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenRequest.ProtoReflect.Descriptor instead.
func (*TokenRequest) Descriptor() ([]byte, []int) {
	return file_eganow_api_auth_auth_payload_proto_rawDescGZIP(), []int{0}
}

func (x *TokenRequest) GetUserGuid() string {
	if x != nil {
		return x.UserGuid
	}
	return ""
}

func (x *TokenRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type ValidateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserGuid     string `protobuf:"bytes,1,opt,name=user_guid,json=userGuid,proto3" json:"user_guid,omitempty"`
	MobileNumber string `protobuf:"bytes,2,opt,name=mobile_number,json=mobileNumber,proto3" json:"mobile_number,omitempty"`
	Email        string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *ValidateTokenResponse) Reset() {
	*x = ValidateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_auth_auth_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResponse) ProtoMessage() {}

func (x *ValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_auth_auth_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_eganow_api_auth_auth_payload_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateTokenResponse) GetUserGuid() string {
	if x != nil {
		return x.UserGuid
	}
	return ""
}

func (x *ValidateTokenResponse) GetMobileNumber() string {
	if x != nil {
		return x.MobileNumber
	}
	return ""
}

func (x *ValidateTokenResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_eganow_api_auth_auth_payload_proto protoreflect.FileDescriptor

var file_eganow_api_auth_auth_payload_proto_rawDesc = []byte{
	0x0a, 0x22, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6f, 0x0a, 0x15, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x47, 0x75, 0x69, 0x64, 0x12,
	0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_eganow_api_auth_auth_payload_proto_rawDescOnce sync.Once
	file_eganow_api_auth_auth_payload_proto_rawDescData = file_eganow_api_auth_auth_payload_proto_rawDesc
)

func file_eganow_api_auth_auth_payload_proto_rawDescGZIP() []byte {
	file_eganow_api_auth_auth_payload_proto_rawDescOnce.Do(func() {
		file_eganow_api_auth_auth_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_auth_auth_payload_proto_rawDescData)
	})
	return file_eganow_api_auth_auth_payload_proto_rawDescData
}

var file_eganow_api_auth_auth_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eganow_api_auth_auth_payload_proto_goTypes = []interface{}{
	(*TokenRequest)(nil),          // 0: eganow.api.auth.TokenRequest
	(*ValidateTokenResponse)(nil), // 1: eganow.api.auth.ValidateTokenResponse
}
var file_eganow_api_auth_auth_payload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eganow_api_auth_auth_payload_proto_init() }
func file_eganow_api_auth_auth_payload_proto_init() {
	if File_eganow_api_auth_auth_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_auth_auth_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenRequest); i {
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
		file_eganow_api_auth_auth_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenResponse); i {
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
			RawDescriptor: file_eganow_api_auth_auth_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eganow_api_auth_auth_payload_proto_goTypes,
		DependencyIndexes: file_eganow_api_auth_auth_payload_proto_depIdxs,
		MessageInfos:      file_eganow_api_auth_auth_payload_proto_msgTypes,
	}.Build()
	File_eganow_api_auth_auth_payload_proto = out.File
	file_eganow_api_auth_auth_payload_proto_rawDesc = nil
	file_eganow_api_auth_auth_payload_proto_goTypes = nil
	file_eganow_api_auth_auth_payload_proto_depIdxs = nil
}
