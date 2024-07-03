// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/merchant/common_payload.proto

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

type MerchantRegulator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MerchantRegulator) Reset() {
	*x = MerchantRegulator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantRegulator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantRegulator) ProtoMessage() {}

func (x *MerchantRegulator) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantRegulator.ProtoReflect.Descriptor instead.
func (*MerchantRegulator) Descriptor() ([]byte, []int) {
	return file_eganow_api_merchant_common_payload_proto_rawDescGZIP(), []int{0}
}

func (x *MerchantRegulator) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MerchantRegulator) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MerchantIndustry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MerchantIndustry) Reset() {
	*x = MerchantIndustry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantIndustry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantIndustry) ProtoMessage() {}

func (x *MerchantIndustry) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantIndustry.ProtoReflect.Descriptor instead.
func (*MerchantIndustry) Descriptor() ([]byte, []int) {
	return file_eganow_api_merchant_common_payload_proto_rawDescGZIP(), []int{1}
}

func (x *MerchantIndustry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MerchantIndustry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MerchantBusinessSector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MerchantBusinessSector) Reset() {
	*x = MerchantBusinessSector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantBusinessSector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantBusinessSector) ProtoMessage() {}

func (x *MerchantBusinessSector) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantBusinessSector.ProtoReflect.Descriptor instead.
func (*MerchantBusinessSector) Descriptor() ([]byte, []int) {
	return file_eganow_api_merchant_common_payload_proto_rawDescGZIP(), []int{2}
}

func (x *MerchantBusinessSector) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MerchantBusinessSector) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MerchantRegulatorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regulators []*MerchantRegulator `protobuf:"bytes,1,rep,name=regulators,proto3" json:"regulators,omitempty"`
}

func (x *MerchantRegulatorList) Reset() {
	*x = MerchantRegulatorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantRegulatorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantRegulatorList) ProtoMessage() {}

func (x *MerchantRegulatorList) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantRegulatorList.ProtoReflect.Descriptor instead.
func (*MerchantRegulatorList) Descriptor() ([]byte, []int) {
	return file_eganow_api_merchant_common_payload_proto_rawDescGZIP(), []int{3}
}

func (x *MerchantRegulatorList) GetRegulators() []*MerchantRegulator {
	if x != nil {
		return x.Regulators
	}
	return nil
}

type MerchantBusinessSectorList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sectors []*MerchantBusinessSector `protobuf:"bytes,1,rep,name=sectors,proto3" json:"sectors,omitempty"`
}

func (x *MerchantBusinessSectorList) Reset() {
	*x = MerchantBusinessSectorList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantBusinessSectorList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantBusinessSectorList) ProtoMessage() {}

func (x *MerchantBusinessSectorList) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantBusinessSectorList.ProtoReflect.Descriptor instead.
func (*MerchantBusinessSectorList) Descriptor() ([]byte, []int) {
	return file_eganow_api_merchant_common_payload_proto_rawDescGZIP(), []int{4}
}

func (x *MerchantBusinessSectorList) GetSectors() []*MerchantBusinessSector {
	if x != nil {
		return x.Sectors
	}
	return nil
}

type MerchantIndustryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Industries []*MerchantIndustry `protobuf:"bytes,1,rep,name=industries,proto3" json:"industries,omitempty"`
}

func (x *MerchantIndustryList) Reset() {
	*x = MerchantIndustryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MerchantIndustryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MerchantIndustryList) ProtoMessage() {}

func (x *MerchantIndustryList) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_merchant_common_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MerchantIndustryList.ProtoReflect.Descriptor instead.
func (*MerchantIndustryList) Descriptor() ([]byte, []int) {
	return file_eganow_api_merchant_common_payload_proto_rawDescGZIP(), []int{5}
}

func (x *MerchantIndustryList) GetIndustries() []*MerchantIndustry {
	if x != nil {
		return x.Industries
	}
	return nil
}

var File_eganow_api_merchant_common_payload_proto protoreflect.FileDescriptor

var file_eganow_api_merchant_common_payload_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x65, 0x67, 0x61, 0x6e,
	0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x22,
	0x37, 0x0a, 0x11, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x10, 0x4d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x3c, 0x0a, 0x16, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5f,
	0x0a, 0x15, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x0a, 0x72, 0x65, 0x67, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x65, 0x67,
	0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e,
	0x74, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x22,
	0x63, 0x0a, 0x1a, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a,
	0x07, 0x73, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x42, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x22, 0x5d, 0x0a, 0x14, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x0a,
	0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x25, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x49,
	0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eganow_api_merchant_common_payload_proto_rawDescOnce sync.Once
	file_eganow_api_merchant_common_payload_proto_rawDescData = file_eganow_api_merchant_common_payload_proto_rawDesc
)

func file_eganow_api_merchant_common_payload_proto_rawDescGZIP() []byte {
	file_eganow_api_merchant_common_payload_proto_rawDescOnce.Do(func() {
		file_eganow_api_merchant_common_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_merchant_common_payload_proto_rawDescData)
	})
	return file_eganow_api_merchant_common_payload_proto_rawDescData
}

var file_eganow_api_merchant_common_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eganow_api_merchant_common_payload_proto_goTypes = []interface{}{
	(*MerchantRegulator)(nil),          // 0: eganow.api.merchant.MerchantRegulator
	(*MerchantIndustry)(nil),           // 1: eganow.api.merchant.MerchantIndustry
	(*MerchantBusinessSector)(nil),     // 2: eganow.api.merchant.MerchantBusinessSector
	(*MerchantRegulatorList)(nil),      // 3: eganow.api.merchant.MerchantRegulatorList
	(*MerchantBusinessSectorList)(nil), // 4: eganow.api.merchant.MerchantBusinessSectorList
	(*MerchantIndustryList)(nil),       // 5: eganow.api.merchant.MerchantIndustryList
}
var file_eganow_api_merchant_common_payload_proto_depIdxs = []int32{
	0, // 0: eganow.api.merchant.MerchantRegulatorList.regulators:type_name -> eganow.api.merchant.MerchantRegulator
	2, // 1: eganow.api.merchant.MerchantBusinessSectorList.sectors:type_name -> eganow.api.merchant.MerchantBusinessSector
	1, // 2: eganow.api.merchant.MerchantIndustryList.industries:type_name -> eganow.api.merchant.MerchantIndustry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eganow_api_merchant_common_payload_proto_init() }
func file_eganow_api_merchant_common_payload_proto_init() {
	if File_eganow_api_merchant_common_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_merchant_common_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantRegulator); i {
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
		file_eganow_api_merchant_common_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantIndustry); i {
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
		file_eganow_api_merchant_common_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantBusinessSector); i {
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
		file_eganow_api_merchant_common_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantRegulatorList); i {
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
		file_eganow_api_merchant_common_payload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantBusinessSectorList); i {
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
		file_eganow_api_merchant_common_payload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MerchantIndustryList); i {
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
			RawDescriptor: file_eganow_api_merchant_common_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eganow_api_merchant_common_payload_proto_goTypes,
		DependencyIndexes: file_eganow_api_merchant_common_payload_proto_depIdxs,
		MessageInfos:      file_eganow_api_merchant_common_payload_proto_msgTypes,
	}.Build()
	File_eganow_api_merchant_common_payload_proto = out.File
	file_eganow_api_merchant_common_payload_proto_rawDesc = nil
	file_eganow_api_merchant_common_payload_proto_goTypes = nil
	file_eganow_api_merchant_common_payload_proto_depIdxs = nil
}
