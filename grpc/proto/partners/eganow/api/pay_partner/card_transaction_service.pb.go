// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: eganow/api/pay_partner/card_transaction_service.proto

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

var File_eganow_api_pay_partner_card_transaction_service_proto protoreflect.FileDescriptor

var file_eganow_api_pay_partner_card_transaction_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x35, 0x65,
	0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdd, 0x02, 0x0a, 0x12, 0x43, 0x61,
	0x72, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x76, 0x63,
	0x12, 0x9f, 0x01, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79,
	0x12, 0x2d, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61,
	0x79, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79,
	0x5f, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x32,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a, 0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x74, 0x65, 0x12, 0xa4, 0x01, 0x0a, 0x15, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x4d, 0x70, 0x67,
	0x73, 0x50, 0x61, 0x79, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x2e, 0x2e, 0x65,
	0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x4d, 0x70, 0x67, 0x73, 0x50, 0x61, 0x79, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x65,
	0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x61, 0x79, 0x5f, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x32, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2c, 0x3a, 0x01, 0x2a,
	0x22, 0x27, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x73, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x61, 0x79,
	0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x67, 0x61, 0x6e, 0x6f, 0x77, 0x2f, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_eganow_api_pay_partner_card_transaction_service_proto_goTypes = []interface{}{
	(*PayTransactionRequest)(nil),  // 0: eganow.api.pay_partner.PayTransactionRequest
	(*MpgsPayCallbackRequest)(nil), // 1: eganow.api.pay_partner.MpgsPayCallbackRequest
	(*PayPartnerStringValue)(nil),  // 2: eganow.api.pay_partner.PayPartnerStringValue
	(*PayPartnerEmpty)(nil),        // 3: eganow.api.pay_partner.PayPartnerEmpty
}
var file_eganow_api_pay_partner_card_transaction_service_proto_depIdxs = []int32{
	0, // 0: eganow.api.pay_partner.CardTransactionSvc.InitiatePay:input_type -> eganow.api.pay_partner.PayTransactionRequest
	1, // 1: eganow.api.pay_partner.CardTransactionSvc.HandleMpgsPayCallback:input_type -> eganow.api.pay_partner.MpgsPayCallbackRequest
	2, // 2: eganow.api.pay_partner.CardTransactionSvc.InitiatePay:output_type -> eganow.api.pay_partner.PayPartnerStringValue
	3, // 3: eganow.api.pay_partner.CardTransactionSvc.HandleMpgsPayCallback:output_type -> eganow.api.pay_partner.PayPartnerEmpty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_eganow_api_pay_partner_card_transaction_service_proto_init() }
func file_eganow_api_pay_partner_card_transaction_service_proto_init() {
	if File_eganow_api_pay_partner_card_transaction_service_proto != nil {
		return
	}
	file_eganow_api_pay_partner_card_transaction_payload_proto_init()
	file_eganow_api_pay_partner_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eganow_api_pay_partner_card_transaction_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eganow_api_pay_partner_card_transaction_service_proto_goTypes,
		DependencyIndexes: file_eganow_api_pay_partner_card_transaction_service_proto_depIdxs,
	}.Build()
	File_eganow_api_pay_partner_card_transaction_service_proto = out.File
	file_eganow_api_pay_partner_card_transaction_service_proto_rawDesc = nil
	file_eganow_api_pay_partner_card_transaction_service_proto_goTypes = nil
	file_eganow_api_pay_partner_card_transaction_service_proto_depIdxs = nil
}
