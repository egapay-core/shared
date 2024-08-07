// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/common/common_data_payload.proto

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

type CountryFilter int32

const (
	CountryFilter_COUNTRY_FILTER_UNSPECIFIED CountryFilter = 0
	CountryFilter_COUNTRY_FILTER_SIGNUP      CountryFilter = 1 // signup countries
	CountryFilter_COUNTRY_FILTER_OPERATING   CountryFilter = 2 // operating countries
)

// Enum value maps for CountryFilter.
var (
	CountryFilter_name = map[int32]string{
		0: "COUNTRY_FILTER_UNSPECIFIED",
		1: "COUNTRY_FILTER_SIGNUP",
		2: "COUNTRY_FILTER_OPERATING",
	}
	CountryFilter_value = map[string]int32{
		"COUNTRY_FILTER_UNSPECIFIED": 0,
		"COUNTRY_FILTER_SIGNUP":      1,
		"COUNTRY_FILTER_OPERATING":   2,
	}
)

func (x CountryFilter) Enum() *CountryFilter {
	p := new(CountryFilter)
	*p = x
	return p
}

func (x CountryFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CountryFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_eganow_api_common_common_data_payload_proto_enumTypes[0].Descriptor()
}

func (CountryFilter) Type() protoreflect.EnumType {
	return &file_eganow_api_common_common_data_payload_proto_enumTypes[0]
}

func (x CountryFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CountryFilter.Descriptor instead.
func (CountryFilter) EnumDescriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{0}
}

type PaymentAccountListType int32

const (
	PaymentAccountListType_PAYMENT_ACCOUNT_LIST_TYPE_UNSPECIFIED PaymentAccountListType = 0
	PaymentAccountListType_PAYMENT_ACCOUNT_LIST_TYPE_SOURCE      PaymentAccountListType = 1
	PaymentAccountListType_PAYMENT_ACCOUNT_LIST_TYPE_BENEFICIARY PaymentAccountListType = 2
)

// Enum value maps for PaymentAccountListType.
var (
	PaymentAccountListType_name = map[int32]string{
		0: "PAYMENT_ACCOUNT_LIST_TYPE_UNSPECIFIED",
		1: "PAYMENT_ACCOUNT_LIST_TYPE_SOURCE",
		2: "PAYMENT_ACCOUNT_LIST_TYPE_BENEFICIARY",
	}
	PaymentAccountListType_value = map[string]int32{
		"PAYMENT_ACCOUNT_LIST_TYPE_UNSPECIFIED": 0,
		"PAYMENT_ACCOUNT_LIST_TYPE_SOURCE":      1,
		"PAYMENT_ACCOUNT_LIST_TYPE_BENEFICIARY": 2,
	}
)

func (x PaymentAccountListType) Enum() *PaymentAccountListType {
	p := new(PaymentAccountListType)
	*p = x
	return p
}

func (x PaymentAccountListType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentAccountListType) Descriptor() protoreflect.EnumDescriptor {
	return file_eganow_api_common_common_data_payload_proto_enumTypes[1].Descriptor()
}

func (PaymentAccountListType) Type() protoreflect.EnumType {
	return &file_eganow_api_common_common_data_payload_proto_enumTypes[1]
}

func (x PaymentAccountListType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentAccountListType.Descriptor instead.
func (PaymentAccountListType) EnumDescriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{1}
}

type GetCountriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter CountryFilter `protobuf:"varint,1,opt,name=filter,proto3,enum=eganow.api.common.CountryFilter" json:"filter,omitempty"`
}

func (x *GetCountriesRequest) Reset() {
	*x = GetCountriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountriesRequest) ProtoMessage() {}

func (x *GetCountriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountriesRequest.ProtoReflect.Descriptor instead.
func (*GetCountriesRequest) Descriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{0}
}

func (x *GetCountriesRequest) GetFilter() CountryFilter {
	if x != nil {
		return x.Filter
	}
	return CountryFilter_COUNTRY_FILTER_UNSPECIFIED
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode           string `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	CountryName           string `protobuf:"bytes,2,opt,name=country_name,json=countryName,proto3" json:"country_name,omitempty"`
	CurrencyIso           string `protobuf:"bytes,3,opt,name=currency_iso,json=currencyIso,proto3" json:"currency_iso,omitempty"`
	CountryFlagUrl        string `protobuf:"bytes,4,opt,name=country_flag_url,json=countryFlagUrl,proto3" json:"country_flag_url,omitempty"`
	MinMobileNumberLength int32  `protobuf:"varint,5,opt,name=min_mobile_number_length,json=minMobileNumberLength,proto3" json:"min_mobile_number_length,omitempty"`
	MaxMobileNumberLength int32  `protobuf:"varint,6,opt,name=max_mobile_number_length,json=maxMobileNumberLength,proto3" json:"max_mobile_number_length,omitempty"`
	DialCode              string `protobuf:"bytes,7,opt,name=dial_code,json=dialCode,proto3" json:"dial_code,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{1}
}

func (x *Country) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Country) GetCountryName() string {
	if x != nil {
		return x.CountryName
	}
	return ""
}

func (x *Country) GetCurrencyIso() string {
	if x != nil {
		return x.CurrencyIso
	}
	return ""
}

func (x *Country) GetCountryFlagUrl() string {
	if x != nil {
		return x.CountryFlagUrl
	}
	return ""
}

func (x *Country) GetMinMobileNumberLength() int32 {
	if x != nil {
		return x.MinMobileNumberLength
	}
	return 0
}

func (x *Country) GetMaxMobileNumberLength() int32 {
	if x != nil {
		return x.MaxMobileNumberLength
	}
	return 0
}

func (x *Country) GetDialCode() string {
	if x != nil {
		return x.DialCode
	}
	return ""
}

type GetCountriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Countries []*Country `protobuf:"bytes,1,rep,name=countries,proto3" json:"countries,omitempty"`
}

func (x *GetCountriesResponse) Reset() {
	*x = GetCountriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountriesResponse) ProtoMessage() {}

func (x *GetCountriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountriesResponse.ProtoReflect.Descriptor instead.
func (*GetCountriesResponse) Descriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{2}
}

func (x *GetCountriesResponse) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

type GetPaymentTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CountryCode string                 `protobuf:"bytes,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	ListType    PaymentAccountListType `protobuf:"varint,2,opt,name=list_type,json=listType,proto3,enum=eganow.api.common.PaymentAccountListType" json:"list_type,omitempty"`
}

func (x *GetPaymentTypesRequest) Reset() {
	*x = GetPaymentTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentTypesRequest) ProtoMessage() {}

func (x *GetPaymentTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentTypesRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentTypesRequest) Descriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentTypesRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *GetPaymentTypesRequest) GetListType() PaymentAccountListType {
	if x != nil {
		return x.ListType
	}
	return PaymentAccountListType_PAYMENT_ACCOUNT_LIST_TYPE_UNSPECIFIED
}

type PaymentType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PayPartnerTypeId      string `protobuf:"bytes,1,opt,name=pay_partner_type_id,json=payPartnerTypeId,proto3" json:"pay_partner_type_id,omitempty"`
	PayPartnerTypeDesc    string `protobuf:"bytes,2,opt,name=pay_partner_type_desc,json=payPartnerTypeDesc,proto3" json:"pay_partner_type_desc,omitempty"`
	AppName               string `protobuf:"bytes,3,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	PayPartnerServiceId   string `protobuf:"bytes,4,opt,name=pay_partner_service_id,json=payPartnerServiceId,proto3" json:"pay_partner_service_id,omitempty"`
	PayPartnerServiceName string `protobuf:"bytes,5,opt,name=pay_partner_service_name,json=payPartnerServiceName,proto3" json:"pay_partner_service_name,omitempty"`
	PayPartnerGroupId     string `protobuf:"bytes,6,opt,name=pay_partner_group_id,json=payPartnerGroupId,proto3" json:"pay_partner_group_id,omitempty"`
	PartnerGatewayCode    string `protobuf:"bytes,7,opt,name=partner_gateway_code,json=partnerGatewayCode,proto3" json:"partner_gateway_code,omitempty"`
	LogoUrl               string `protobuf:"bytes,8,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
}

func (x *PaymentType) Reset() {
	*x = PaymentType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentType) ProtoMessage() {}

func (x *PaymentType) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentType.ProtoReflect.Descriptor instead.
func (*PaymentType) Descriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{4}
}

func (x *PaymentType) GetPayPartnerTypeId() string {
	if x != nil {
		return x.PayPartnerTypeId
	}
	return ""
}

func (x *PaymentType) GetPayPartnerTypeDesc() string {
	if x != nil {
		return x.PayPartnerTypeDesc
	}
	return ""
}

func (x *PaymentType) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *PaymentType) GetPayPartnerServiceId() string {
	if x != nil {
		return x.PayPartnerServiceId
	}
	return ""
}

func (x *PaymentType) GetPayPartnerServiceName() string {
	if x != nil {
		return x.PayPartnerServiceName
	}
	return ""
}

func (x *PaymentType) GetPayPartnerGroupId() string {
	if x != nil {
		return x.PayPartnerGroupId
	}
	return ""
}

func (x *PaymentType) GetPartnerGatewayCode() string {
	if x != nil {
		return x.PartnerGatewayCode
	}
	return ""
}

func (x *PaymentType) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

type GetPaymentTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentTypes []*PaymentType `protobuf:"bytes,1,rep,name=payment_types,json=paymentTypes,proto3" json:"payment_types,omitempty"`
}

func (x *GetPaymentTypesResponse) Reset() {
	*x = GetPaymentTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentTypesResponse) ProtoMessage() {}

func (x *GetPaymentTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_common_common_data_payload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentTypesResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentTypesResponse) Descriptor() ([]byte, []int) {
	return file_eganow_api_common_common_data_payload_proto_rawDescGZIP(), []int{5}
}

func (x *GetPaymentTypesResponse) GetPaymentTypes() []*PaymentType {
	if x != nil {
		return x.PaymentTypes
	}
	return nil
}

var File_eganow_api_common_common_data_payload_proto protoreflect.FileDescriptor

var file_eganow_api_common_common_data_payload_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x65,
	0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x22, 0x4f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x22, 0xab, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f,
	0x69, 0x73, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x49, 0x73, 0x6f, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x46, 0x6c, 0x61, 0x67, 0x55, 0x72, 0x6c,
	0x12, 0x37, 0x0a, 0x18, 0x6d, 0x69, 0x6e, 0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x6d, 0x69, 0x6e, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x37, 0x0a, 0x18, 0x6d, 0x61, 0x78,
	0x5f, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x15, 0x6d, 0x61, 0x78,
	0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x6e, 0x67,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x22,
	0x50, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x67, 0x61,
	0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x22, 0x83, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x46, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x29, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c,
	0x69, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xf6, 0x02, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2d, 0x0a, 0x13, 0x70, 0x61, 0x79, 0x5f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x61, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x73, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x70, 0x70,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x70, 0x61, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x37, 0x0a, 0x18, 0x70, 0x61, 0x79,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x70, 0x61, 0x79,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x70, 0x61, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x5f, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x6f, 0x55, 0x72, 0x6c,
	0x22, 0x5e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2a, 0x68, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x55, 0x50, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x52, 0x59, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x2a, 0x94, 0x01, 0x0a, 0x16, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x24, 0x0a, 0x20, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f,
	0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x10, 0x01, 0x12, 0x29, 0x0a, 0x25, 0x50, 0x41, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x45, 0x4e, 0x45, 0x46, 0x49, 0x43, 0x49, 0x41, 0x52, 0x59, 0x10,
	0x02, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eganow_api_common_common_data_payload_proto_rawDescOnce sync.Once
	file_eganow_api_common_common_data_payload_proto_rawDescData = file_eganow_api_common_common_data_payload_proto_rawDesc
)

func file_eganow_api_common_common_data_payload_proto_rawDescGZIP() []byte {
	file_eganow_api_common_common_data_payload_proto_rawDescOnce.Do(func() {
		file_eganow_api_common_common_data_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_common_common_data_payload_proto_rawDescData)
	})
	return file_eganow_api_common_common_data_payload_proto_rawDescData
}

var file_eganow_api_common_common_data_payload_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_eganow_api_common_common_data_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_eganow_api_common_common_data_payload_proto_goTypes = []interface{}{
	(CountryFilter)(0),              // 0: eganow.api.common.CountryFilter
	(PaymentAccountListType)(0),     // 1: eganow.api.common.PaymentAccountListType
	(*GetCountriesRequest)(nil),     // 2: eganow.api.common.GetCountriesRequest
	(*Country)(nil),                 // 3: eganow.api.common.Country
	(*GetCountriesResponse)(nil),    // 4: eganow.api.common.GetCountriesResponse
	(*GetPaymentTypesRequest)(nil),  // 5: eganow.api.common.GetPaymentTypesRequest
	(*PaymentType)(nil),             // 6: eganow.api.common.PaymentType
	(*GetPaymentTypesResponse)(nil), // 7: eganow.api.common.GetPaymentTypesResponse
}
var file_eganow_api_common_common_data_payload_proto_depIdxs = []int32{
	0, // 0: eganow.api.common.GetCountriesRequest.filter:type_name -> eganow.api.common.CountryFilter
	3, // 1: eganow.api.common.GetCountriesResponse.countries:type_name -> eganow.api.common.Country
	1, // 2: eganow.api.common.GetPaymentTypesRequest.list_type:type_name -> eganow.api.common.PaymentAccountListType
	6, // 3: eganow.api.common.GetPaymentTypesResponse.payment_types:type_name -> eganow.api.common.PaymentType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_eganow_api_common_common_data_payload_proto_init() }
func file_eganow_api_common_common_data_payload_proto_init() {
	if File_eganow_api_common_common_data_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_common_common_data_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountriesRequest); i {
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
		file_eganow_api_common_common_data_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_eganow_api_common_common_data_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountriesResponse); i {
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
		file_eganow_api_common_common_data_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentTypesRequest); i {
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
		file_eganow_api_common_common_data_payload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentType); i {
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
		file_eganow_api_common_common_data_payload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPaymentTypesResponse); i {
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
			RawDescriptor: file_eganow_api_common_common_data_payload_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eganow_api_common_common_data_payload_proto_goTypes,
		DependencyIndexes: file_eganow_api_common_common_data_payload_proto_depIdxs,
		EnumInfos:         file_eganow_api_common_common_data_payload_proto_enumTypes,
		MessageInfos:      file_eganow_api_common_common_data_payload_proto_msgTypes,
	}.Build()
	File_eganow_api_common_common_data_payload_proto = out.File
	file_eganow_api_common_common_data_payload_proto_rawDesc = nil
	file_eganow_api_common_common_data_payload_proto_goTypes = nil
	file_eganow_api_common_common_data_payload_proto_depIdxs = nil
}
