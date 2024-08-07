// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/core/CardAuthentication.proto

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

type Card3DsInputDataRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId                 string  `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	TransactionId           string  `protobuf:"bytes,2,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	TranAmount              float64 `protobuf:"fixed64,3,opt,name=tranAmount,proto3" json:"tranAmount,omitempty"`
	CardExpiryMonth         int32   `protobuf:"varint,4,opt,name=cardExpiryMonth,proto3" json:"cardExpiryMonth,omitempty"`
	CardExpiryYear          int32   `protobuf:"varint,5,opt,name=cardExpiryYear,proto3" json:"cardExpiryYear,omitempty"`
	DeviceMobileOrWeb       string  `protobuf:"bytes,6,opt,name=deviceMobileOrWeb,proto3" json:"deviceMobileOrWeb,omitempty"`
	Cvv                     string  `protobuf:"bytes,7,opt,name=cvv,proto3" json:"cvv,omitempty"`
	CardHolderName          string  `protobuf:"bytes,8,opt,name=cardHolderName,proto3" json:"cardHolderName,omitempty"`
	Currency                string  `protobuf:"bytes,9,opt,name=currency,proto3" json:"currency,omitempty"`
	CardNumber              string  `protobuf:"bytes,10,opt,name=cardNumber,proto3" json:"cardNumber,omitempty"`
	TransTypeAuthOrPurchase string  `protobuf:"bytes,11,opt,name=transTypeAuthOrPurchase,proto3" json:"transTypeAuthOrPurchase,omitempty"`
	ServiceNameOrdescriptor string  `protobuf:"bytes,12,opt,name=serviceNameOrdescriptor,proto3" json:"serviceNameOrdescriptor,omitempty"`
	MerchantCardApiId       string  `protobuf:"bytes,13,opt,name=merchantCardApiId,proto3" json:"merchantCardApiId,omitempty"`
	MerchantCardApiPassword string  `protobuf:"bytes,14,opt,name=merchantCardApiPassword,proto3" json:"merchantCardApiPassword,omitempty"`
	EndpointFor3DSCheck     string  `protobuf:"bytes,50,opt,name=endpointFor3DSCheck,proto3" json:"endpointFor3DSCheck,omitempty"`
	ProviderEnumValue       string  `protobuf:"bytes,51,opt,name=providerEnumValue,proto3" json:"providerEnumValue,omitempty"`
}

func (x *Card3DsInputDataRequestMessage) Reset() {
	*x = Card3DsInputDataRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_core_CardAuthentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card3DsInputDataRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card3DsInputDataRequestMessage) ProtoMessage() {}

func (x *Card3DsInputDataRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_core_CardAuthentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card3DsInputDataRequestMessage.ProtoReflect.Descriptor instead.
func (*Card3DsInputDataRequestMessage) Descriptor() ([]byte, []int) {
	return file_eganow_api_core_CardAuthentication_proto_rawDescGZIP(), []int{0}
}

func (x *Card3DsInputDataRequestMessage) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetTranAmount() float64 {
	if x != nil {
		return x.TranAmount
	}
	return 0
}

func (x *Card3DsInputDataRequestMessage) GetCardExpiryMonth() int32 {
	if x != nil {
		return x.CardExpiryMonth
	}
	return 0
}

func (x *Card3DsInputDataRequestMessage) GetCardExpiryYear() int32 {
	if x != nil {
		return x.CardExpiryYear
	}
	return 0
}

func (x *Card3DsInputDataRequestMessage) GetDeviceMobileOrWeb() string {
	if x != nil {
		return x.DeviceMobileOrWeb
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetCardHolderName() string {
	if x != nil {
		return x.CardHolderName
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetTransTypeAuthOrPurchase() string {
	if x != nil {
		return x.TransTypeAuthOrPurchase
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetServiceNameOrdescriptor() string {
	if x != nil {
		return x.ServiceNameOrdescriptor
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetMerchantCardApiId() string {
	if x != nil {
		return x.MerchantCardApiId
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetMerchantCardApiPassword() string {
	if x != nil {
		return x.MerchantCardApiPassword
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetEndpointFor3DSCheck() string {
	if x != nil {
		return x.EndpointFor3DSCheck
	}
	return ""
}

func (x *Card3DsInputDataRequestMessage) GetProviderEnumValue() string {
	if x != nil {
		return x.ProviderEnumValue
	}
	return ""
}

type Card3DsDataResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess                           bool   `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	Message                             string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Brand                               string `protobuf:"bytes,3,opt,name=brand,proto3" json:"brand,omitempty"`
	SuccessOrFailOrPendingTransStatus   string `protobuf:"bytes,4,opt,name=successOrFailOrPendingTransStatus,proto3" json:"successOrFailOrPendingTransStatus,omitempty"`
	Scheme                              string `protobuf:"bytes,5,opt,name=scheme,proto3" json:"scheme,omitempty"`
	CardCurrencyIso                     string `protobuf:"bytes,6,opt,name=cardCurrencyIso,proto3" json:"cardCurrencyIso,omitempty"`
	RedirectUrl                         string `protobuf:"bytes,7,opt,name=redirectUrl,proto3" json:"redirectUrl,omitempty"`
	CardIs3DsYesOrNo                    string `protobuf:"bytes,8,opt,name=cardIs3DsYesOrNo,proto3" json:"cardIs3DsYesOrNo,omitempty"`
	GatewayCode                         string `protobuf:"bytes,9,opt,name=gatewayCode,proto3" json:"gatewayCode,omitempty"`
	GatewayRecommendation               string `protobuf:"bytes,10,opt,name=gatewayRecommendation,proto3" json:"gatewayRecommendation,omitempty"`
	OrderIdFromCardGateway              string `protobuf:"bytes,11,opt,name=orderIdFromCardGateway,proto3" json:"orderIdFromCardGateway,omitempty"`
	CardDsTransactionIdFromCardGateway  string `protobuf:"bytes,12,opt,name=cardDsTransactionIdFromCardGateway,proto3" json:"cardDsTransactionIdFromCardGateway,omitempty"`
	CardAcsTransactionIdFromCardGateway string `protobuf:"bytes,13,opt,name=cardAcsTransactionIdFromCardGateway,proto3" json:"cardAcsTransactionIdFromCardGateway,omitempty"`
	Card3DsAuthResponseUrl              string `protobuf:"bytes,14,opt,name=card3DsAuthResponseUrl,proto3" json:"card3DsAuthResponseUrl,omitempty"`
	CardTypeDebitCreditPrepaidVirtual   string `protobuf:"bytes,20,opt,name=cardTypeDebitCreditPrepaidVirtual,proto3" json:"cardTypeDebitCreditPrepaidVirtual,omitempty"`
	CardCountryOfIssue                  string `protobuf:"bytes,21,opt,name=cardCountryOfIssue,proto3" json:"cardCountryOfIssue,omitempty"`
	CardIssuingBankOrgName              string `protobuf:"bytes,22,opt,name=cardIssuingBankOrgName,proto3" json:"cardIssuingBankOrgName,omitempty"`
}

func (x *Card3DsDataResponseMessage) Reset() {
	*x = Card3DsDataResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eganow_api_core_CardAuthentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card3DsDataResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card3DsDataResponseMessage) ProtoMessage() {}

func (x *Card3DsDataResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_eganow_api_core_CardAuthentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card3DsDataResponseMessage.ProtoReflect.Descriptor instead.
func (*Card3DsDataResponseMessage) Descriptor() ([]byte, []int) {
	return file_eganow_api_core_CardAuthentication_proto_rawDescGZIP(), []int{1}
}

func (x *Card3DsDataResponseMessage) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *Card3DsDataResponseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetSuccessOrFailOrPendingTransStatus() string {
	if x != nil {
		return x.SuccessOrFailOrPendingTransStatus
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCardCurrencyIso() string {
	if x != nil {
		return x.CardCurrencyIso
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCardIs3DsYesOrNo() string {
	if x != nil {
		return x.CardIs3DsYesOrNo
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetGatewayCode() string {
	if x != nil {
		return x.GatewayCode
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetGatewayRecommendation() string {
	if x != nil {
		return x.GatewayRecommendation
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetOrderIdFromCardGateway() string {
	if x != nil {
		return x.OrderIdFromCardGateway
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCardDsTransactionIdFromCardGateway() string {
	if x != nil {
		return x.CardDsTransactionIdFromCardGateway
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCardAcsTransactionIdFromCardGateway() string {
	if x != nil {
		return x.CardAcsTransactionIdFromCardGateway
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCard3DsAuthResponseUrl() string {
	if x != nil {
		return x.Card3DsAuthResponseUrl
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCardTypeDebitCreditPrepaidVirtual() string {
	if x != nil {
		return x.CardTypeDebitCreditPrepaidVirtual
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCardCountryOfIssue() string {
	if x != nil {
		return x.CardCountryOfIssue
	}
	return ""
}

func (x *Card3DsDataResponseMessage) GetCardIssuingBankOrgName() string {
	if x != nil {
		return x.CardIssuingBankOrgName
	}
	return ""
}

var File_eganow_api_core_CardAuthentication_proto protoreflect.FileDescriptor

var file_eganow_api_core_CardAuthentication_proto_rawDesc = []byte{
	0x0a, 0x28, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x43, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x43, 0x61, 0x72, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb2,
	0x05, 0x0a, 0x1e, 0x43, 0x61, 0x72, 0x64, 0x33, 0x44, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61, 0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x64,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x0e, 0x63,
	0x61, 0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x59, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x63, 0x61, 0x72, 0x64, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x4f, 0x72, 0x57, 0x65, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x4f, 0x72, 0x57, 0x65,
	0x62, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x76, 0x76, 0x12, 0x26, 0x0a, 0x0e, 0x63, 0x61, 0x72, 0x64, 0x48, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x61, 0x72,
	0x64, 0x48, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x12, 0x2c, 0x0a, 0x11, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x41, 0x70, 0x69, 0x49, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x41, 0x70, 0x69, 0x49, 0x64, 0x12, 0x38, 0x0a, 0x17, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x41, 0x70, 0x69, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x41, 0x70, 0x69, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46,
	0x6f, 0x72, 0x33, 0x44, 0x53, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x33, 0x44, 0x53,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0xe8, 0x06, 0x0a, 0x1a, 0x43, 0x61, 0x72, 0x64, 0x33, 0x44, 0x73, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x4c, 0x0a, 0x21, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x46, 0x61, 0x69,
	0x6c, 0x4f, 0x72, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x21, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4f, 0x72, 0x46, 0x61, 0x69, 0x6c, 0x4f, 0x72, 0x50, 0x65, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x61, 0x72, 0x64, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x73, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x61, 0x72, 0x64, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x73, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55,
	0x72, 0x6c, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x61, 0x72, 0x64, 0x49, 0x73, 0x33, 0x44, 0x73, 0x59,
	0x65, 0x73, 0x4f, 0x72, 0x4e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x61,
	0x72, 0x64, 0x49, 0x73, 0x33, 0x44, 0x73, 0x59, 0x65, 0x73, 0x4f, 0x72, 0x4e, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x34, 0x0a, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x4e,
	0x0a, 0x22, 0x63, 0x61, 0x72, 0x64, 0x44, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x47, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x22, 0x63, 0x61, 0x72, 0x64,
	0x44, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x50,
	0x0a, 0x23, 0x63, 0x61, 0x72, 0x64, 0x41, 0x63, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x47, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x23, 0x63, 0x61, 0x72,
	0x64, 0x41, 0x63, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x61, 0x72, 0x64, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x12, 0x36, 0x0a, 0x16, 0x63, 0x61, 0x72, 0x64, 0x33, 0x44, 0x73, 0x41, 0x75, 0x74, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x63, 0x61, 0x72, 0x64, 0x33, 0x44, 0x73, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x4c, 0x0a, 0x21, 0x63, 0x61, 0x72, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x62, 0x69, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x69, 0x64, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x21, 0x63, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x62,
	0x69, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x50, 0x72, 0x65, 0x70, 0x61, 0x69, 0x64, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x61, 0x72, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x66, 0x49, 0x73, 0x73, 0x75, 0x65, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x12, 0x63, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4f,
	0x66, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x63, 0x61, 0x72, 0x64, 0x49, 0x73,
	0x73, 0x75, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6e, 0x6b, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x63, 0x61, 0x72, 0x64, 0x49, 0x73, 0x73, 0x75,
	0x69, 0x6e, 0x67, 0x42, 0x61, 0x6e, 0x6b, 0x4f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0xb8,
	0x02, 0x0a, 0x15, 0x43, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x76, 0x63, 0x12, 0x88, 0x01, 0x0a, 0x22, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x33, 0x64, 0x73, 0x12,
	0x32, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x33, 0x44, 0x73, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x2e, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x33, 0x44, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x93, 0x01, 0x0a, 0x2d, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x61, 0x72,
	0x64, 0x46, 0x6f, 0x72, 0x54, 0x68, 0x72, 0x65, 0x65, 0x44, 0x53, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x41, 0x73, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x32, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x33,
	0x44, 0x73, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x2e, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43,
	0x61, 0x72, 0x64, 0x33, 0x44, 0x73, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x41, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x47, 0x72, 0x70, 0x63, 0x45, 0x4d, 0x56, 0x43, 0x61,
	0x72, 0x64, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eganow_api_core_CardAuthentication_proto_rawDescOnce sync.Once
	file_eganow_api_core_CardAuthentication_proto_rawDescData = file_eganow_api_core_CardAuthentication_proto_rawDesc
)

func file_eganow_api_core_CardAuthentication_proto_rawDescGZIP() []byte {
	file_eganow_api_core_CardAuthentication_proto_rawDescOnce.Do(func() {
		file_eganow_api_core_CardAuthentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_eganow_api_core_CardAuthentication_proto_rawDescData)
	})
	return file_eganow_api_core_CardAuthentication_proto_rawDescData
}

var file_eganow_api_core_CardAuthentication_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_eganow_api_core_CardAuthentication_proto_goTypes = []interface{}{
	(*Card3DsInputDataRequestMessage)(nil), // 0: CardAuthentication.Card3DsInputDataRequestMessage
	(*Card3DsDataResponseMessage)(nil),     // 1: CardAuthentication.Card3DsDataResponseMessage
}
var file_eganow_api_core_CardAuthentication_proto_depIdxs = []int32{
	0, // 0: CardAuthentication.CardAuthenticationSvc.AuthenticateCardTransactionWith3ds:input_type -> CardAuthentication.Card3DsInputDataRequestMessage
	0, // 1: CardAuthentication.CardAuthenticationSvc.CheckCardForThreeDSToAddAsEganowPaymentMethod:input_type -> CardAuthentication.Card3DsInputDataRequestMessage
	1, // 2: CardAuthentication.CardAuthenticationSvc.AuthenticateCardTransactionWith3ds:output_type -> CardAuthentication.Card3DsDataResponseMessage
	1, // 3: CardAuthentication.CardAuthenticationSvc.CheckCardForThreeDSToAddAsEganowPaymentMethod:output_type -> CardAuthentication.Card3DsDataResponseMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eganow_api_core_CardAuthentication_proto_init() }
func file_eganow_api_core_CardAuthentication_proto_init() {
	if File_eganow_api_core_CardAuthentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eganow_api_core_CardAuthentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card3DsInputDataRequestMessage); i {
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
		file_eganow_api_core_CardAuthentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card3DsDataResponseMessage); i {
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
			RawDescriptor: file_eganow_api_core_CardAuthentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eganow_api_core_CardAuthentication_proto_goTypes,
		DependencyIndexes: file_eganow_api_core_CardAuthentication_proto_depIdxs,
		MessageInfos:      file_eganow_api_core_CardAuthentication_proto_msgTypes,
	}.Build()
	File_eganow_api_core_CardAuthentication_proto = out.File
	file_eganow_api_core_CardAuthentication_proto_rawDesc = nil
	file_eganow_api_core_CardAuthentication_proto_goTypes = nil
	file_eganow_api_core_CardAuthentication_proto_depIdxs = nil
}
