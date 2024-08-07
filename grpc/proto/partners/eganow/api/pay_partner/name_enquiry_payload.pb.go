// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/pay_partner/name_enquiry_payload.proto

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

// NameEnquiryRequest is the request sent to the NameEnquiry service
type NameEnquiryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AccountNumber:
	//
	//	*NameEnquiryRequest_BankNameEnquiryRequest_
	//	*NameEnquiryRequest_PhoneNumber
	AccountNumber isNameEnquiryRequest_AccountNumber `protobuf_oneof:"account_number"`
}

func (x *NameEnquiryRequest) Reset() {
	*x = NameEnquiryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameEnquiryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameEnquiryRequest) ProtoMessage() {}

func (x *NameEnquiryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameEnquiryRequest.ProtoReflect.Descriptor instead.
func (*NameEnquiryRequest) Descriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescGZIP(), []int{0}
}

func (m *NameEnquiryRequest) GetAccountNumber() isNameEnquiryRequest_AccountNumber {
	if m != nil {
		return m.AccountNumber
	}
	return nil
}

func (x *NameEnquiryRequest) GetBankNameEnquiryRequest() *NameEnquiryRequest_BankNameEnquiryRequest {
	if x, ok := x.GetAccountNumber().(*NameEnquiryRequest_BankNameEnquiryRequest_); ok {
		return x.BankNameEnquiryRequest
	}
	return nil
}

func (x *NameEnquiryRequest) GetPhoneNumber() string {
	if x, ok := x.GetAccountNumber().(*NameEnquiryRequest_PhoneNumber); ok {
		return x.PhoneNumber
	}
	return ""
}

type isNameEnquiryRequest_AccountNumber interface {
	isNameEnquiryRequest_AccountNumber()
}

type NameEnquiryRequest_BankNameEnquiryRequest_ struct {
	BankNameEnquiryRequest *NameEnquiryRequest_BankNameEnquiryRequest `protobuf:"bytes,2,opt,name=bank_name_enquiry_request,json=bankNameEnquiryRequest,proto3,oneof"`
}

type NameEnquiryRequest_PhoneNumber struct {
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3,oneof"`
}

func (*NameEnquiryRequest_BankNameEnquiryRequest_) isNameEnquiryRequest_AccountNumber() {}

func (*NameEnquiryRequest_PhoneNumber) isNameEnquiryRequest_AccountNumber() {}

type GhQrNameEnquiryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientAccountNumber string              `protobuf:"bytes,1,opt,name=recipient_account_number,json=recipientAccountNumber,proto3" json:"recipient_account_number,omitempty"`
	TerminalId             string              `protobuf:"bytes,2,opt,name=terminal_id,json=terminalId,proto3" json:"terminal_id,omitempty"` // required for GHQR - GHIPSS
	BankCode               string              `protobuf:"bytes,3,opt,name=bank_code,json=bankCode,proto3" json:"bank_code,omitempty"`       // 3 digits
	SenderAccountName      string              `protobuf:"bytes,4,opt,name=sender_account_name,json=senderAccountName,proto3" json:"sender_account_name,omitempty"`
	SenderAccountNumber    string              `protobuf:"bytes,5,opt,name=sender_account_number,json=senderAccountNumber,proto3" json:"sender_account_number,omitempty"`
	RequestType            BankNameEnquiryType `protobuf:"varint,6,opt,name=request_type,json=requestType,proto3,enum=eganow.api.pay_partner.BankNameEnquiryType" json:"request_type,omitempty"`
}

func (x *GhQrNameEnquiryRequest) Reset() {
	*x = GhQrNameEnquiryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GhQrNameEnquiryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GhQrNameEnquiryRequest) ProtoMessage() {}

func (x *GhQrNameEnquiryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GhQrNameEnquiryRequest.ProtoReflect.Descriptor instead.
func (*GhQrNameEnquiryRequest) Descriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescGZIP(), []int{1}
}

func (x *GhQrNameEnquiryRequest) GetRecipientAccountNumber() string {
	if x != nil {
		return x.RecipientAccountNumber
	}
	return ""
}

func (x *GhQrNameEnquiryRequest) GetTerminalId() string {
	if x != nil {
		return x.TerminalId
	}
	return ""
}

func (x *GhQrNameEnquiryRequest) GetBankCode() string {
	if x != nil {
		return x.BankCode
	}
	return ""
}

func (x *GhQrNameEnquiryRequest) GetSenderAccountName() string {
	if x != nil {
		return x.SenderAccountName
	}
	return ""
}

func (x *GhQrNameEnquiryRequest) GetSenderAccountNumber() string {
	if x != nil {
		return x.SenderAccountNumber
	}
	return ""
}

func (x *GhQrNameEnquiryRequest) GetRequestType() BankNameEnquiryType {
	if x != nil {
		return x.RequestType
	}
	return BankNameEnquiryType_UNSPECIFIED
}

type GhQrNameEnquiryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AccountHolder:
	//
	//	*GhQrNameEnquiryResponse_Name
	//	*GhQrNameEnquiryResponse_QrCode
	AccountHolder isGhQrNameEnquiryResponse_AccountHolder `protobuf_oneof:"account_holder"`
}

func (x *GhQrNameEnquiryResponse) Reset() {
	*x = GhQrNameEnquiryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GhQrNameEnquiryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GhQrNameEnquiryResponse) ProtoMessage() {}

func (x *GhQrNameEnquiryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GhQrNameEnquiryResponse.ProtoReflect.Descriptor instead.
func (*GhQrNameEnquiryResponse) Descriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescGZIP(), []int{2}
}

func (m *GhQrNameEnquiryResponse) GetAccountHolder() isGhQrNameEnquiryResponse_AccountHolder {
	if m != nil {
		return m.AccountHolder
	}
	return nil
}

func (x *GhQrNameEnquiryResponse) GetName() string {
	if x, ok := x.GetAccountHolder().(*GhQrNameEnquiryResponse_Name); ok {
		return x.Name
	}
	return ""
}

func (x *GhQrNameEnquiryResponse) GetQrCode() string {
	if x, ok := x.GetAccountHolder().(*GhQrNameEnquiryResponse_QrCode); ok {
		return x.QrCode
	}
	return ""
}

type isGhQrNameEnquiryResponse_AccountHolder interface {
	isGhQrNameEnquiryResponse_AccountHolder()
}

type GhQrNameEnquiryResponse_Name struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3,oneof"`
}

type GhQrNameEnquiryResponse_QrCode struct {
	QrCode string `protobuf:"bytes,2,opt,name=qr_code,json=qrCode,proto3,oneof"`
}

func (*GhQrNameEnquiryResponse_Name) isGhQrNameEnquiryResponse_AccountHolder() {}

func (*GhQrNameEnquiryResponse_QrCode) isGhQrNameEnquiryResponse_AccountHolder() {}

type NameEnquiryRequest_BankNameEnquiryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecipientAccountNumber string              `protobuf:"bytes,1,opt,name=recipient_account_number,json=recipientAccountNumber,proto3" json:"recipient_account_number,omitempty"`
	TerminalId             string              `protobuf:"bytes,2,opt,name=terminal_id,json=terminalId,proto3" json:"terminal_id,omitempty"` // required for GHQR - GHIPSS
	BankCode               string              `protobuf:"bytes,3,opt,name=bank_code,json=bankCode,proto3" json:"bank_code,omitempty"`       // 3 digits
	SenderAccountName      string              `protobuf:"bytes,4,opt,name=sender_account_name,json=senderAccountName,proto3" json:"sender_account_name,omitempty"`
	SenderAccountNumber    string              `protobuf:"bytes,5,opt,name=sender_account_number,json=senderAccountNumber,proto3" json:"sender_account_number,omitempty"`
	RequestType            BankNameEnquiryType `protobuf:"varint,6,opt,name=request_type,json=requestType,proto3,enum=eganow.api.pay_partner.BankNameEnquiryType" json:"request_type,omitempty"`
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) Reset() {
	*x = NameEnquiryRequest_BankNameEnquiryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameEnquiryRequest_BankNameEnquiryRequest) ProtoMessage() {}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameEnquiryRequest_BankNameEnquiryRequest.ProtoReflect.Descriptor instead.
func (*NameEnquiryRequest_BankNameEnquiryRequest) Descriptor() ([]byte, []int) {
	return file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescGZIP(), []int{0, 0}
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) GetRecipientAccountNumber() string {
	if x != nil {
		return x.RecipientAccountNumber
	}
	return ""
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) GetTerminalId() string {
	if x != nil {
		return x.TerminalId
	}
	return ""
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) GetBankCode() string {
	if x != nil {
		return x.BankCode
	}
	return ""
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) GetSenderAccountName() string {
	if x != nil {
		return x.SenderAccountName
	}
	return ""
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) GetSenderAccountNumber() string {
	if x != nil {
		return x.SenderAccountNumber
	}
	return ""
}

func (x *NameEnquiryRequest_BankNameEnquiryRequest) GetRequestType() BankNameEnquiryType {
	if x != nil {
		return x.RequestType
	}
	return BankNameEnquiryType_UNSPECIFIED
}

var File_eganow_api_pay_partner_name_enquiry_payload_proto protoreflect.FileDescriptor

var file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDesc = []byte{
	0x0a, 0x31, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x65, 0x67,
	0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72,
	0x79, 0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0,
	0x04, 0x0a, 0x12, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x7e, 0x0a, 0x19, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f,
	0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71,
	0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x16, 0x62,
	0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x48, 0x00, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x1a,
	0xdd, 0x02, 0x0a, 0x16, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75,
	0x69, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x18, 0x72, 0x65,
	0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41,
	0x02, 0x52, 0x16, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x61,
	0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x13,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x11,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x37, 0x0a, 0x15, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x0c, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2b, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x10, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0xe2, 0x02, 0x0a, 0x16, 0x47, 0x68, 0x51, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e,
	0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x18,
	0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03,
	0xe0, 0x41, 0x02, 0x52, 0x16, 0x72, 0x65, 0x63, 0x69, 0x70, 0x69, 0x65, 0x6e, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0b, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x13, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x11, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x15, 0x73, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x13, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x53, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x2e, 0x42, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5c, 0x0a, 0x17, 0x47, 0x68, 0x51, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x45, 0x6e, 0x71, 0x75, 0x69, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x71, 0x72, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x71, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescOnce sync.Once
	file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescData = file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDesc
)

func file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescGZIP() []byte {
	file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescOnce.Do(func() {
		file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescData)
	})
	return file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDescData
}

var file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eganow_api_pay_partner_name_enquiry_payload_proto_goTypes = []interface{}{
	(*NameEnquiryRequest)(nil),                        // 0: eganow.api.pay_partner.NameEnquiryRequest
	(*GhQrNameEnquiryRequest)(nil),                    // 1: eganow.api.pay_partner.GhQrNameEnquiryRequest
	(*GhQrNameEnquiryResponse)(nil),                   // 2: eganow.api.pay_partner.GhQrNameEnquiryResponse
	(*NameEnquiryRequest_BankNameEnquiryRequest)(nil), // 3: eganow.api.pay_partner.NameEnquiryRequest.BankNameEnquiryRequest
	(BankNameEnquiryType)(0),                          // 4: eganow.api.pay_partner.BankNameEnquiryType
}
var file_eganow_api_pay_partner_name_enquiry_payload_proto_depIdxs = []int32{
	3, // 0: eganow.api.pay_partner.NameEnquiryRequest.bank_name_enquiry_request:type_name -> eganow.api.pay_partner.NameEnquiryRequest.BankNameEnquiryRequest
	4, // 1: eganow.api.pay_partner.GhQrNameEnquiryRequest.request_type:type_name -> eganow.api.pay_partner.BankNameEnquiryType
	4, // 2: eganow.api.pay_partner.NameEnquiryRequest.BankNameEnquiryRequest.request_type:type_name -> eganow.api.pay_partner.BankNameEnquiryType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eganow_api_pay_partner_name_enquiry_payload_proto_init() }
func file_eganow_api_pay_partner_name_enquiry_payload_proto_init() {
	if File_eganow_api_pay_partner_name_enquiry_payload_proto != nil {
		return
	}
	file_eganow_api_pay_partner_name_enquiry_entity_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameEnquiryRequest); i {
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
		file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GhQrNameEnquiryRequest); i {
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
		file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GhQrNameEnquiryResponse); i {
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
		file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameEnquiryRequest_BankNameEnquiryRequest); i {
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
	file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*NameEnquiryRequest_BankNameEnquiryRequest_)(nil),
		(*NameEnquiryRequest_PhoneNumber)(nil),
	}
	file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*GhQrNameEnquiryResponse_Name)(nil),
		(*GhQrNameEnquiryResponse_QrCode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eganow_api_pay_partner_name_enquiry_payload_proto_goTypes,
		DependencyIndexes: file_eganow_api_pay_partner_name_enquiry_payload_proto_depIdxs,
		MessageInfos:      file_eganow_api_pay_partner_name_enquiry_payload_proto_msgTypes,
	}.Build()
	File_eganow_api_pay_partner_name_enquiry_payload_proto = out.File
	file_eganow_api_pay_partner_name_enquiry_payload_proto_rawDesc = nil
	file_eganow_api_pay_partner_name_enquiry_payload_proto_goTypes = nil
	file_eganow_api_pay_partner_name_enquiry_payload_proto_depIdxs = nil
}
