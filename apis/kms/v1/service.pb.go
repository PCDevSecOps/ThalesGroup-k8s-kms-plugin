// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: apis/kms/v1/service.proto

package kms

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_apis_kms_v1_service_proto protoreflect.FileDescriptor

var file_apis_kms_v1_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x68, 0x61,
	0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xd7, 0x06, 0x0a, 0x14, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43,
	0x41, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x12, 0x2b, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73,
	0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x41, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c,
	0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x41, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12,
	0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x2f, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0xda, 0x41,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x74, 0x0a, 0x07, 0x53, 0x69, 0x67, 0x6e, 0x43, 0x53, 0x52,
	0x12, 0x27, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x43,
	0x53, 0x52, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x68, 0x61, 0x6c,
	0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x43, 0x53, 0x52, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x61, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0xa6, 0x01, 0x0a, 0x0f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x12,
	0x2f, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x22, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x4b, 0x65, 0x79, 0x22, 0x3e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x08, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x3a, 0x0a, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x5f, 0x6b,
	0x65, 0x79, 0xda, 0x41, 0x1f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x2c, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x5f, 0x6b, 0x65, 0x79, 0x12, 0x83, 0x01, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2e, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73,
	0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73,
	0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x4b, 0x65, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a,
	0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x8c, 0x01, 0x0a, 0x07, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x27, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63,
	0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b,
	0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x2c,
	0x70, 0x6c, 0x61, 0x69, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x8d, 0x01, 0x0a, 0x07, 0x44, 0x65,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x12, 0x27, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70,
	0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x3a, 0x01, 0x2a, 0xda, 0x41, 0x0f, 0x6e, 0x61, 0x6d, 0x65, 0x2c, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70,
	0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x6b, 0x38, 0x73, 0x2d, 0x6b, 0x6d, 0x73, 0x2d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6b, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x6b, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apis_kms_v1_service_proto_goTypes = []interface{}{
	(*GetCABundleRequest)(nil),     // 0: thalescpl.io.k8s.kms.v1.GetCABundleRequest
	(*SignCSRRequest)(nil),         // 1: thalescpl.io.k8s.kms.v1.SignCSRRequest
	(*CreateCryptoKeyRequest)(nil), // 2: thalescpl.io.k8s.kms.v1.CreateCryptoKeyRequest
	(*ListCryptoKeysRequest)(nil),  // 3: thalescpl.io.k8s.kms.v1.ListCryptoKeysRequest
	(*EncryptRequest)(nil),         // 4: thalescpl.io.k8s.kms.v1.EncryptRequest
	(*DecryptRequest)(nil),         // 5: thalescpl.io.k8s.kms.v1.DecryptRequest
	(*CABundle)(nil),               // 6: thalescpl.io.k8s.kms.v1.CABundle
	(*SignCSRResponse)(nil),        // 7: thalescpl.io.k8s.kms.v1.SignCSRResponse
	(*CryptoKey)(nil),              // 8: thalescpl.io.k8s.kms.v1.CryptoKey
	(*ListCryptoKeysResponse)(nil), // 9: thalescpl.io.k8s.kms.v1.ListCryptoKeysResponse
	(*EncryptResponse)(nil),        // 10: thalescpl.io.k8s.kms.v1.EncryptResponse
	(*DecryptResponse)(nil),        // 11: thalescpl.io.k8s.kms.v1.DecryptResponse
}
var file_apis_kms_v1_service_proto_depIdxs = []int32{
	0,  // 0: thalescpl.io.k8s.kms.v1.KeyManagementService.GetCABundle:input_type -> thalescpl.io.k8s.kms.v1.GetCABundleRequest
	1,  // 1: thalescpl.io.k8s.kms.v1.KeyManagementService.SignCSR:input_type -> thalescpl.io.k8s.kms.v1.SignCSRRequest
	2,  // 2: thalescpl.io.k8s.kms.v1.KeyManagementService.CreateCryptoKey:input_type -> thalescpl.io.k8s.kms.v1.CreateCryptoKeyRequest
	3,  // 3: thalescpl.io.k8s.kms.v1.KeyManagementService.ListCryptoKeys:input_type -> thalescpl.io.k8s.kms.v1.ListCryptoKeysRequest
	4,  // 4: thalescpl.io.k8s.kms.v1.KeyManagementService.Encrypt:input_type -> thalescpl.io.k8s.kms.v1.EncryptRequest
	5,  // 5: thalescpl.io.k8s.kms.v1.KeyManagementService.Decrypt:input_type -> thalescpl.io.k8s.kms.v1.DecryptRequest
	6,  // 6: thalescpl.io.k8s.kms.v1.KeyManagementService.GetCABundle:output_type -> thalescpl.io.k8s.kms.v1.CABundle
	7,  // 7: thalescpl.io.k8s.kms.v1.KeyManagementService.SignCSR:output_type -> thalescpl.io.k8s.kms.v1.SignCSRResponse
	8,  // 8: thalescpl.io.k8s.kms.v1.KeyManagementService.CreateCryptoKey:output_type -> thalescpl.io.k8s.kms.v1.CryptoKey
	9,  // 9: thalescpl.io.k8s.kms.v1.KeyManagementService.ListCryptoKeys:output_type -> thalescpl.io.k8s.kms.v1.ListCryptoKeysResponse
	10, // 10: thalescpl.io.k8s.kms.v1.KeyManagementService.Encrypt:output_type -> thalescpl.io.k8s.kms.v1.EncryptResponse
	11, // 11: thalescpl.io.k8s.kms.v1.KeyManagementService.Decrypt:output_type -> thalescpl.io.k8s.kms.v1.DecryptResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_apis_kms_v1_service_proto_init() }
func file_apis_kms_v1_service_proto_init() {
	if File_apis_kms_v1_service_proto != nil {
		return
	}
	file_apis_kms_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_kms_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_kms_v1_service_proto_goTypes,
		DependencyIndexes: file_apis_kms_v1_service_proto_depIdxs,
	}.Build()
	File_apis_kms_v1_service_proto = out.File
	file_apis_kms_v1_service_proto_rawDesc = nil
	file_apis_kms_v1_service_proto_goTypes = nil
	file_apis_kms_v1_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeyManagementServiceClient is the client API for KeyManagementService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeyManagementServiceClient interface {
	// Returns the CABundle
	GetCABundle(ctx context.Context, in *GetCABundleRequest, opts ...grpc.CallOption) (*CABundle, error)
	// Create a new [CryptoKey][google.cloud.kms.v1.CryptoKey] within a [KeyRing][google.cloud.kms.v1.KeyRing].
	//
	// [CryptoKey.purpose][google.cloud.kms.v1.CryptoKey.purpose] and
	// [CryptoKey.version_template.algorithm][google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm]
	// are required.
	SignCSR(ctx context.Context, in *SignCSRRequest, opts ...grpc.CallOption) (*SignCSRResponse, error)
	// Create a new [CryptoKey][google.cloud.kms.v1.CryptoKey] within a [KeyRing][google.cloud.kms.v1.KeyRing].
	//
	// [CryptoKey.purpose][google.cloud.kms.v1.CryptoKey.purpose] and
	// [CryptoKey.version_template.algorithm][google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm]
	// are required.
	CreateCryptoKey(ctx context.Context, in *CreateCryptoKeyRequest, opts ...grpc.CallOption) (*CryptoKey, error)
	// Create a new [CryptoKey][google.cloud.kms.v1.CryptoKey] within a [KeyRing][google.cloud.kms.v1.KeyRing].
	//
	// [CryptoKey.purpose][google.cloud.kms.v1.CryptoKey.purpose] and
	// [CryptoKey.version_template.algorithm][google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm]
	// are required.
	ListCryptoKeys(ctx context.Context, in *ListCryptoKeysRequest, opts ...grpc.CallOption) (*ListCryptoKeysResponse, error)
	// Encrypts data, so that it can only be recovered by a call to [Decrypt][thalescpl.io.istio.v1.KeyManagementService.Decrypt].
	// The [CryptoKey.purpose][thalescpl.io.istio.v1.CryptoKey.purpose] must be
	// [ENCRYPT_DECRYPT][thalescpl.io.istio.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT].
	Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error)
	// Decrypts data that was protected by [Encrypt][thalescpl.io.istio.v1.KeyManagementService.Encrypt]. The [CryptoKey.purpose][thalescpl.io.istio.v1.CryptoKey.purpose]
	// must be [ENCRYPT_DECRYPT][thalescpl.io.istio.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT].
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type keyManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyManagementServiceClient(cc grpc.ClientConnInterface) KeyManagementServiceClient {
	return &keyManagementServiceClient{cc}
}

func (c *keyManagementServiceClient) GetCABundle(ctx context.Context, in *GetCABundleRequest, opts ...grpc.CallOption) (*CABundle, error) {
	out := new(CABundle)
	err := c.cc.Invoke(ctx, "/thalescpl.io.k8s.kms.v1.KeyManagementService/GetCABundle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) SignCSR(ctx context.Context, in *SignCSRRequest, opts ...grpc.CallOption) (*SignCSRResponse, error) {
	out := new(SignCSRResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.k8s.kms.v1.KeyManagementService/SignCSR", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) CreateCryptoKey(ctx context.Context, in *CreateCryptoKeyRequest, opts ...grpc.CallOption) (*CryptoKey, error) {
	out := new(CryptoKey)
	err := c.cc.Invoke(ctx, "/thalescpl.io.k8s.kms.v1.KeyManagementService/CreateCryptoKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) ListCryptoKeys(ctx context.Context, in *ListCryptoKeysRequest, opts ...grpc.CallOption) (*ListCryptoKeysResponse, error) {
	out := new(ListCryptoKeysResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.k8s.kms.v1.KeyManagementService/ListCryptoKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) Encrypt(ctx context.Context, in *EncryptRequest, opts ...grpc.CallOption) (*EncryptResponse, error) {
	out := new(EncryptResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.k8s.kms.v1.KeyManagementService/Encrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.k8s.kms.v1.KeyManagementService/Decrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyManagementServiceServer is the server API for KeyManagementService service.
type KeyManagementServiceServer interface {
	// Returns the CABundle
	GetCABundle(context.Context, *GetCABundleRequest) (*CABundle, error)
	// Create a new [CryptoKey][google.cloud.kms.v1.CryptoKey] within a [KeyRing][google.cloud.kms.v1.KeyRing].
	//
	// [CryptoKey.purpose][google.cloud.kms.v1.CryptoKey.purpose] and
	// [CryptoKey.version_template.algorithm][google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm]
	// are required.
	SignCSR(context.Context, *SignCSRRequest) (*SignCSRResponse, error)
	// Create a new [CryptoKey][google.cloud.kms.v1.CryptoKey] within a [KeyRing][google.cloud.kms.v1.KeyRing].
	//
	// [CryptoKey.purpose][google.cloud.kms.v1.CryptoKey.purpose] and
	// [CryptoKey.version_template.algorithm][google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm]
	// are required.
	CreateCryptoKey(context.Context, *CreateCryptoKeyRequest) (*CryptoKey, error)
	// Create a new [CryptoKey][google.cloud.kms.v1.CryptoKey] within a [KeyRing][google.cloud.kms.v1.KeyRing].
	//
	// [CryptoKey.purpose][google.cloud.kms.v1.CryptoKey.purpose] and
	// [CryptoKey.version_template.algorithm][google.cloud.kms.v1.CryptoKeyVersionTemplate.algorithm]
	// are required.
	ListCryptoKeys(context.Context, *ListCryptoKeysRequest) (*ListCryptoKeysResponse, error)
	// Encrypts data, so that it can only be recovered by a call to [Decrypt][thalescpl.io.istio.v1.KeyManagementService.Decrypt].
	// The [CryptoKey.purpose][thalescpl.io.istio.v1.CryptoKey.purpose] must be
	// [ENCRYPT_DECRYPT][thalescpl.io.istio.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT].
	Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error)
	// Decrypts data that was protected by [Encrypt][thalescpl.io.istio.v1.KeyManagementService.Encrypt]. The [CryptoKey.purpose][thalescpl.io.istio.v1.CryptoKey.purpose]
	// must be [ENCRYPT_DECRYPT][thalescpl.io.istio.v1.CryptoKey.CryptoKeyPurpose.ENCRYPT_DECRYPT].
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
}

// UnimplementedKeyManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeyManagementServiceServer struct {
}

func (*UnimplementedKeyManagementServiceServer) GetCABundle(context.Context, *GetCABundleRequest) (*CABundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCABundle not implemented")
}
func (*UnimplementedKeyManagementServiceServer) SignCSR(context.Context, *SignCSRRequest) (*SignCSRResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCSR not implemented")
}
func (*UnimplementedKeyManagementServiceServer) CreateCryptoKey(context.Context, *CreateCryptoKeyRequest) (*CryptoKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCryptoKey not implemented")
}
func (*UnimplementedKeyManagementServiceServer) ListCryptoKeys(context.Context, *ListCryptoKeysRequest) (*ListCryptoKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCryptoKeys not implemented")
}
func (*UnimplementedKeyManagementServiceServer) Encrypt(context.Context, *EncryptRequest) (*EncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (*UnimplementedKeyManagementServiceServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}

func RegisterKeyManagementServiceServer(s *grpc.Server, srv KeyManagementServiceServer) {
	s.RegisterService(&_KeyManagementService_serviceDesc, srv)
}

func _KeyManagementService_GetCABundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCABundleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).GetCABundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.k8s.kms.v1.KeyManagementService/GetCABundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).GetCABundle(ctx, req.(*GetCABundleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_SignCSR_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCSRRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).SignCSR(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.k8s.kms.v1.KeyManagementService/SignCSR",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).SignCSR(ctx, req.(*SignCSRRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_CreateCryptoKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCryptoKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).CreateCryptoKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.k8s.kms.v1.KeyManagementService/CreateCryptoKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).CreateCryptoKey(ctx, req.(*CreateCryptoKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_ListCryptoKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCryptoKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).ListCryptoKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.k8s.kms.v1.KeyManagementService/ListCryptoKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).ListCryptoKeys(ctx, req.(*ListCryptoKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.k8s.kms.v1.KeyManagementService/Encrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Encrypt(ctx, req.(*EncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.k8s.kms.v1.KeyManagementService/Decrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thalescpl.io.k8s.kms.v1.KeyManagementService",
	HandlerType: (*KeyManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCABundle",
			Handler:    _KeyManagementService_GetCABundle_Handler,
		},
		{
			MethodName: "SignCSR",
			Handler:    _KeyManagementService_SignCSR_Handler,
		},
		{
			MethodName: "CreateCryptoKey",
			Handler:    _KeyManagementService_CreateCryptoKey_Handler,
		},
		{
			MethodName: "ListCryptoKeys",
			Handler:    _KeyManagementService_ListCryptoKeys_Handler,
		},
		{
			MethodName: "Encrypt",
			Handler:    _KeyManagementService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _KeyManagementService_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/kms/v1/service.proto",
}
