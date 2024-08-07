// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/pay_partner/fund_transfer_entity.proto

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

type TransactionStatus int32

const (
	TransactionStatus_PENDING    TransactionStatus = 0
	TransactionStatus_SUCCESSFUL TransactionStatus = 1
	TransactionStatus_FAILED     TransactionStatus = 2
)

// Enum value maps for TransactionStatus.
var (
	TransactionStatus_name = map[int32]string{
		0: "PENDING",
		1: "SUCCESSFUL",
		2: "FAILED",
	}
	TransactionStatus_value = map[string]int32{
		"PENDING":    0,
		"SUCCESSFUL": 1,
		"FAILED":     2,
	}
)

func (x TransactionStatus) Enum() *TransactionStatus {
	p := new(TransactionStatus)
	*p = x
	return p
}

func (x TransactionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes[0].Descriptor()
}

func (TransactionStatus) Type() protoreflect.EnumType {
	return &file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes[0]
}

func (x TransactionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionStatus.Descriptor instead.
func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescGZIP(), []int{0}
}

type TransactionType int32

const (
	TransactionType_COLLECTION TransactionType = 0
	TransactionType_PAYOUT     TransactionType = 1
	TransactionType_REMITTANCE TransactionType = 2 // for remittance (GHIPSS)
	TransactionType_QR_CREDIT  TransactionType = 3 // for QR code credit transactions
)

// Enum value maps for TransactionType.
var (
	TransactionType_name = map[int32]string{
		0: "COLLECTION",
		1: "PAYOUT",
		2: "REMITTANCE",
		3: "QR_CREDIT",
	}
	TransactionType_value = map[string]int32{
		"COLLECTION": 0,
		"PAYOUT":     1,
		"REMITTANCE": 2,
		"QR_CREDIT":  3,
	}
)

func (x TransactionType) Enum() *TransactionType {
	p := new(TransactionType)
	*p = x
	return p
}

func (x TransactionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransactionType) Descriptor() protoreflect.EnumDescriptor {
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes[1].Descriptor()
}

func (TransactionType) Type() protoreflect.EnumType {
	return &file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes[1]
}

func (x TransactionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransactionType.Descriptor instead.
func (TransactionType) EnumDescriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescGZIP(), []int{1}
}

type TransLogMomoBankMerchant_UpdateType int32

const (
	TransLogMomoBankMerchant_AUTO   TransLogMomoBankMerchant_UpdateType = 0
	TransLogMomoBankMerchant_MANUAL TransLogMomoBankMerchant_UpdateType = 1
)

// Enum value maps for TransLogMomoBankMerchant_UpdateType.
var (
	TransLogMomoBankMerchant_UpdateType_name = map[int32]string{
		0: "AUTO",
		1: "MANUAL",
	}
	TransLogMomoBankMerchant_UpdateType_value = map[string]int32{
		"AUTO":   0,
		"MANUAL": 1,
	}
)

func (x TransLogMomoBankMerchant_UpdateType) Enum() *TransLogMomoBankMerchant_UpdateType {
	p := new(TransLogMomoBankMerchant_UpdateType)
	*p = x
	return p
}

func (x TransLogMomoBankMerchant_UpdateType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransLogMomoBankMerchant_UpdateType) Descriptor() protoreflect.EnumDescriptor {
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes[2].Descriptor()
}

func (TransLogMomoBankMerchant_UpdateType) Type() protoreflect.EnumType {
	return &file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes[2]
}

func (x TransLogMomoBankMerchant_UpdateType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransLogMomoBankMerchant_UpdateType.Descriptor instead.
func (TransLogMomoBankMerchant_UpdateType) EnumDescriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescGZIP(), []int{0, 0}
}

type TransLogMomoBankMerchant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TID                         int32                               `protobuf:"varint,18,opt,name=TID,proto3" json:"TID,omitempty"` // transaction ID (auto increment)
	EganowTransRefNo            string                              `protobuf:"bytes,1,opt,name=EganowTransRefNo,proto3" json:"EganowTransRefNo,omitempty"`
	PartnerTransRefNo           string                              `protobuf:"bytes,6,opt,name=PartnerTransRefNo,proto3" json:"PartnerTransRefNo,omitempty"` // update with value received from pay partner
	AccountNumberOrMSISDN       string                              `protobuf:"bytes,2,opt,name=AccountNumberOrMSISDN,proto3" json:"AccountNumberOrMSISDN,omitempty"`
	EganowPartnerId             string                              `protobuf:"bytes,3,opt,name=EganowPartnerId,proto3" json:"EganowPartnerId,omitempty"` // MTN, Vodafone, AirtelTigo, GHIPSS
	TransAmount                 float32                             `protobuf:"fixed32,4,opt,name=TransAmount,proto3" json:"TransAmount,omitempty"`
	TransTypeCollectionOrPayout TransactionType                     `protobuf:"varint,5,opt,name=TransTypeCollectionOrPayout,proto3,enum=eganow.api.pay_partner.TransactionType" json:"TransTypeCollectionOrPayout,omitempty"`
	RequestID                   string                              `protobuf:"bytes,7,opt,name=RequestID,proto3" json:"RequestID,omitempty"`   // ??
	ResponseID                  string                              `protobuf:"bytes,8,opt,name=ResponseID,proto3" json:"ResponseID,omitempty"` // ??
	TransNarration              string                              `protobuf:"bytes,9,opt,name=TransNarration,proto3" json:"TransNarration,omitempty"`
	RequestDate                 string                              `protobuf:"bytes,10,opt,name=RequestDate,proto3" json:"RequestDate,omitempty"`
	AutoOrManualUpdate          TransLogMomoBankMerchant_UpdateType `protobuf:"varint,12,opt,name=AutoOrManualUpdate,proto3,enum=eganow.api.pay_partner.TransLogMomoBankMerchant_UpdateType" json:"AutoOrManualUpdate,omitempty"`
	InitialRequestPayload       string                              `protobuf:"bytes,14,opt,name=InitialRequestPayload,proto3" json:"InitialRequestPayload,omitempty"`   // JSON/string
	InitialResponsePayload      string                              `protobuf:"bytes,15,opt,name=InitialResponsePayload,proto3" json:"InitialResponsePayload,omitempty"` // update with value received from pay partner (JSON/string)
	EganowTransactionStatus     TransactionStatus                   `protobuf:"varint,16,opt,name=EganowTransactionStatus,proto3,enum=eganow.api.pay_partner.TransactionStatus" json:"EganowTransactionStatus,omitempty"`
	PayPartnerTransactionStatus string                              `protobuf:"bytes,17,opt,name=PayPartnerTransactionStatus,proto3" json:"PayPartnerTransactionStatus,omitempty"`
	TransactionCallbackResponse string                              `protobuf:"bytes,19,opt,name=TransactionCallbackResponse,proto3" json:"TransactionCallbackResponse,omitempty"`
	LastUpdated                 string                              `protobuf:"bytes,11,opt,name=LastUpdated,proto3" json:"LastUpdated,omitempty"`
	UpdatedBy                   string                              `protobuf:"bytes,13,opt,name=UpdatedBy,proto3" json:"UpdatedBy,omitempty"`
}

func (x *TransLogMomoBankMerchant) Reset() {
	*x = TransLogMomoBankMerchant{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_pay_partner_fund_transfer_entity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransLogMomoBankMerchant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransLogMomoBankMerchant) ProtoMessage() {}

func (x *TransLogMomoBankMerchant) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_pay_partner_fund_transfer_entity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransLogMomoBankMerchant.ProtoReflect.Descriptor instead.
func (*TransLogMomoBankMerchant) Descriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescGZIP(), []int{0}
}

func (x *TransLogMomoBankMerchant) GetTID() int32 {
	if x != nil {
		return x.TID
	}
	return 0
}

func (x *TransLogMomoBankMerchant) GetEganowTransRefNo() string {
	if x != nil {
		return x.EganowTransRefNo
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetPartnerTransRefNo() string {
	if x != nil {
		return x.PartnerTransRefNo
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetAccountNumberOrMSISDN() string {
	if x != nil {
		return x.AccountNumberOrMSISDN
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetEganowPartnerId() string {
	if x != nil {
		return x.EganowPartnerId
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetTransAmount() float32 {
	if x != nil {
		return x.TransAmount
	}
	return 0
}

func (x *TransLogMomoBankMerchant) GetTransTypeCollectionOrPayout() TransactionType {
	if x != nil {
		return x.TransTypeCollectionOrPayout
	}
	return TransactionType_COLLECTION
}

func (x *TransLogMomoBankMerchant) GetRequestID() string {
	if x != nil {
		return x.RequestID
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetResponseID() string {
	if x != nil {
		return x.ResponseID
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetTransNarration() string {
	if x != nil {
		return x.TransNarration
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetRequestDate() string {
	if x != nil {
		return x.RequestDate
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetAutoOrManualUpdate() TransLogMomoBankMerchant_UpdateType {
	if x != nil {
		return x.AutoOrManualUpdate
	}
	return TransLogMomoBankMerchant_AUTO
}

func (x *TransLogMomoBankMerchant) GetInitialRequestPayload() string {
	if x != nil {
		return x.InitialRequestPayload
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetInitialResponsePayload() string {
	if x != nil {
		return x.InitialResponsePayload
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetEganowTransactionStatus() TransactionStatus {
	if x != nil {
		return x.EganowTransactionStatus
	}
	return TransactionStatus_PENDING
}

func (x *TransLogMomoBankMerchant) GetPayPartnerTransactionStatus() string {
	if x != nil {
		return x.PayPartnerTransactionStatus
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetTransactionCallbackResponse() string {
	if x != nil {
		return x.TransactionCallbackResponse
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

func (x *TransLogMomoBankMerchant) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

var File_eganow_api_pay_partner_fund_transfer_entity_proto protoreflect.FileDescriptor

var file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x66, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0xa3, 0x08, 0x0a, 0x18,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x4c, 0x6f, 0x67, 0x4d, 0x6f, 0x6d, 0x6f, 0x42, 0x61, 0x6e, 0x6b,
	0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x49, 0x44, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x54, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x67,
	0x61, 0x6e, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52, 0x65, 0x66, 0x4e, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x11, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x52,
	0x65, 0x66, 0x4e, 0x6f, 0x12, 0x34, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x72, 0x4d, 0x53, 0x49, 0x53, 0x44, 0x4e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x72, 0x4d, 0x53, 0x49, 0x53, 0x44, 0x4e, 0x12, 0x28, 0x0a, 0x0f, 0x45, 0x67,
	0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x69, 0x0a, 0x1b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x50,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x65, 0x67,
	0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x1b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72, 0x50, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x49, 0x44, 0x12,
	0x26, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4e, 0x61, 0x72, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4e, 0x61,
	0x72, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x6b, 0x0a, 0x12, 0x41, 0x75, 0x74,
	0x6f, 0x4f, 0x72, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x4c, 0x6f, 0x67, 0x4d, 0x6f, 0x6d, 0x6f, 0x42, 0x61, 0x6e, 0x6b, 0x4d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x12, 0x41, 0x75, 0x74, 0x6f, 0x4f, 0x72, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x15, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x36, 0x0a, 0x16,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x49, 0x6e,
	0x69, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x63, 0x0a, 0x17, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x17, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x1b, 0x50, 0x61, 0x79,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b,
	0x50, 0x61, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x40, 0x0a, 0x1b, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x1b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x4c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0x22, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x41,
	0x55, 0x54, 0x4f, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10,
	0x01, 0x2a, 0x3c, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e,
	0x47, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x46, 0x55,
	0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x2a,
	0x4c, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4c, 0x4c, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x59, 0x4f, 0x55, 0x54, 0x10, 0x01, 0x12, 0x0e,
	0x0a, 0x0a, 0x52, 0x45, 0x4d, 0x49, 0x54, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0d,
	0x0a, 0x09, 0x51, 0x52, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54, 0x10, 0x03, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e,
	0x6f, 0x77, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescOnce sync.Once
	file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescData = file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDesc
)

func file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescGZIP() []byte {
	file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescOnce.Do(func() {
		file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescData)
	})
	return file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDescData
}

var file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_eganow_api_pay_partner_fund_transfer_entity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_eganow_api_pay_partner_fund_transfer_entity_proto_goTypes = []interface{}{
	(TransactionStatus)(0),                   // 0: eganow.api.pay_partner.TransactionStatus
	(TransactionType)(0),                     // 1: eganow.api.pay_partner.TransactionType
	(TransLogMomoBankMerchant_UpdateType)(0), // 2: eganow.api.pay_partner.TransLogMomoBankMerchant.UpdateType
	(*TransLogMomoBankMerchant)(nil),         // 3: eganow.api.pay_partner.TransLogMomoBankMerchant
}
var file_eganow_api_pay_partner_fund_transfer_entity_proto_depIdxs = []int32{
	1, // 0: eganow.api.pay_partner.TransLogMomoBankMerchant.TransTypeCollectionOrPayout:type_name -> eganow.api.pay_partner.TransactionType
	2, // 1: eganow.api.pay_partner.TransLogMomoBankMerchant.AutoOrManualUpdate:type_name -> eganow.api.pay_partner.TransLogMomoBankMerchant.UpdateType
	0, // 2: eganow.api.pay_partner.TransLogMomoBankMerchant.EganowTransactionStatus:type_name -> eganow.api.pay_partner.TransactionStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eganow_api_pay_partner_fund_transfer_entity_proto_init() }
func file_eganow_api_pay_partner_fund_transfer_entity_proto_init() {
	if File_eganow_api_pay_partner_fund_transfer_entity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_pay_partner_fund_transfer_entity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransLogMomoBankMerchant); i {
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
			RawDescriptor: file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eganow_api_pay_partner_fund_transfer_entity_proto_goTypes,
		DependencyIndexes: file_eganow_api_pay_partner_fund_transfer_entity_proto_depIdxs,
		EnumInfos:         file_eganow_api_pay_partner_fund_transfer_entity_proto_enumTypes,
		MessageInfos:      file_eganow_api_pay_partner_fund_transfer_entity_proto_msgTypes,
	}.Build()
	File_eganow_api_pay_partner_fund_transfer_entity_proto = out.File
	file_eganow_api_pay_partner_fund_transfer_entity_proto_rawDesc = nil
	file_eganow_api_pay_partner_fund_transfer_entity_proto_goTypes = nil
	file_eganow_api_pay_partner_fund_transfer_entity_proto_depIdxs = nil
}
