// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.10.0
// source: apis/istio/v1/service.proto

package istio

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

var File_apis_istio_v1_service_proto protoreflect.FileDescriptor

var file_apis_istio_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74,
	0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73,
	0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe8, 0x05, 0x0a, 0x14, 0x4b, 0x65, 0x79,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x45, 0x4b,
	0x12, 0x2e, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b,
	0x12, 0x2e, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2f, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x45, 0x4b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x71, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x4b, 0x65,
	0x79, 0x12, 0x2f, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f,
	0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69,
	0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x08, 0x4c, 0x6f, 0x61, 0x64, 0x53, 0x4b, 0x65, 0x79,
	0x12, 0x2b, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x61, 0x64, 0x53, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d,
	0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x61, 0x64, 0x53,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x14,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x12, 0x37, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c,
	0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e,
	0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d,
	0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x12, 0x37, 0x2e, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e,
	0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x74, 0x68, 0x61, 0x6c,
	0x65, 0x73, 0x63, 0x70, 0x6c, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6b, 0x6d, 0x73, 0x2e, 0x69, 0x73,
	0x74, 0x69, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x64, 0x44, 0x65, 0x63, 0x72, 0x79, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x68, 0x61, 0x6c, 0x65, 0x73, 0x63, 0x70, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x6b,
	0x38, 0x73, 0x2d, 0x6b, 0x6d, 0x73, 0x2d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x73, 0x74, 0x69,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apis_istio_v1_service_proto_goTypes = []interface{}{
	(*GenerateKEKRequest)(nil),           // 0: thalescpl.io.ekms.istio.v1.GenerateKEKRequest
	(*GenerateDEKRequest)(nil),           // 1: thalescpl.io.ekms.istio.v1.GenerateDEKRequest
	(*GenerateSKeyRequest)(nil),          // 2: thalescpl.io.ekms.istio.v1.GenerateSKeyRequest
	(*LoadSKeyRequest)(nil),              // 3: thalescpl.io.ekms.istio.v1.LoadSKeyRequest
	(*AuthenticatedEncryptRequest)(nil),  // 4: thalescpl.io.ekms.istio.v1.AuthenticatedEncryptRequest
	(*AuthenticatedDecryptRequest)(nil),  // 5: thalescpl.io.ekms.istio.v1.AuthenticatedDecryptRequest
	(*GenerateKEKResponse)(nil),          // 6: thalescpl.io.ekms.istio.v1.GenerateKEKResponse
	(*GenerateDEKResponse)(nil),          // 7: thalescpl.io.ekms.istio.v1.GenerateDEKResponse
	(*GenerateSKeyResponse)(nil),         // 8: thalescpl.io.ekms.istio.v1.GenerateSKeyResponse
	(*LoadSKeyResponse)(nil),             // 9: thalescpl.io.ekms.istio.v1.LoadSKeyResponse
	(*AuthenticatedEncryptResponse)(nil), // 10: thalescpl.io.ekms.istio.v1.AuthenticatedEncryptResponse
	(*AuthenticatedDecryptResponse)(nil), // 11: thalescpl.io.ekms.istio.v1.AuthenticatedDecryptResponse
}
var file_apis_istio_v1_service_proto_depIdxs = []int32{
	0,  // 0: thalescpl.io.ekms.istio.v1.KeyManagementService.GenerateKEK:input_type -> thalescpl.io.ekms.istio.v1.GenerateKEKRequest
	1,  // 1: thalescpl.io.ekms.istio.v1.KeyManagementService.GenerateDEK:input_type -> thalescpl.io.ekms.istio.v1.GenerateDEKRequest
	2,  // 2: thalescpl.io.ekms.istio.v1.KeyManagementService.GenerateSKey:input_type -> thalescpl.io.ekms.istio.v1.GenerateSKeyRequest
	3,  // 3: thalescpl.io.ekms.istio.v1.KeyManagementService.LoadSKey:input_type -> thalescpl.io.ekms.istio.v1.LoadSKeyRequest
	4,  // 4: thalescpl.io.ekms.istio.v1.KeyManagementService.AuthenticatedEncrypt:input_type -> thalescpl.io.ekms.istio.v1.AuthenticatedEncryptRequest
	5,  // 5: thalescpl.io.ekms.istio.v1.KeyManagementService.AuthenticatedDecrypt:input_type -> thalescpl.io.ekms.istio.v1.AuthenticatedDecryptRequest
	6,  // 6: thalescpl.io.ekms.istio.v1.KeyManagementService.GenerateKEK:output_type -> thalescpl.io.ekms.istio.v1.GenerateKEKResponse
	7,  // 7: thalescpl.io.ekms.istio.v1.KeyManagementService.GenerateDEK:output_type -> thalescpl.io.ekms.istio.v1.GenerateDEKResponse
	8,  // 8: thalescpl.io.ekms.istio.v1.KeyManagementService.GenerateSKey:output_type -> thalescpl.io.ekms.istio.v1.GenerateSKeyResponse
	9,  // 9: thalescpl.io.ekms.istio.v1.KeyManagementService.LoadSKey:output_type -> thalescpl.io.ekms.istio.v1.LoadSKeyResponse
	10, // 10: thalescpl.io.ekms.istio.v1.KeyManagementService.AuthenticatedEncrypt:output_type -> thalescpl.io.ekms.istio.v1.AuthenticatedEncryptResponse
	11, // 11: thalescpl.io.ekms.istio.v1.KeyManagementService.AuthenticatedDecrypt:output_type -> thalescpl.io.ekms.istio.v1.AuthenticatedDecryptResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_apis_istio_v1_service_proto_init() }
func file_apis_istio_v1_service_proto_init() {
	if File_apis_istio_v1_service_proto != nil {
		return
	}
	file_apis_istio_v1_messages_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apis_istio_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apis_istio_v1_service_proto_goTypes,
		DependencyIndexes: file_apis_istio_v1_service_proto_depIdxs,
	}.Build()
	File_apis_istio_v1_service_proto = out.File
	file_apis_istio_v1_service_proto_rawDesc = nil
	file_apis_istio_v1_service_proto_goTypes = nil
	file_apis_istio_v1_service_proto_depIdxs = nil
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
	// GenerateKEK returns the KID of the GeneratedKEK if allowed/successful
	GenerateKEK(ctx context.Context, in *GenerateKEKRequest, opts ...grpc.CallOption) (*GenerateKEKResponse, error)
	// GenerateDEK returns a wrapped (by HSM handled KEK)
	GenerateDEK(ctx context.Context, in *GenerateDEKRequest, opts ...grpc.CallOption) (*GenerateDEKResponse, error)
	// GenerateSKey returns a wrapped (by provided encrypted DEK ), for later use during loading and signing key generation
	GenerateSKey(ctx context.Context, in *GenerateSKeyRequest, opts ...grpc.CallOption) (*GenerateSKeyResponse, error)
	// LoadSKey returns the SKey unwrapped for the controller to use for CA work...
	LoadSKey(ctx context.Context, in *LoadSKeyRequest, opts ...grpc.CallOption) (*LoadSKeyResponse, error)
	// AuthenticatedEncrypt returns a payload protect by AEAD (Authenticated Encryption with Associated Data) AES-GCM with 256-bit
	AuthenticatedEncrypt(ctx context.Context, in *AuthenticatedEncryptRequest, opts ...grpc.CallOption) (*AuthenticatedEncryptResponse, error)
	// AuthenticatedDecrypt returns a payload decrypted from the previously invoke AE
	AuthenticatedDecrypt(ctx context.Context, in *AuthenticatedDecryptRequest, opts ...grpc.CallOption) (*AuthenticatedDecryptResponse, error)
}

type keyManagementServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyManagementServiceClient(cc grpc.ClientConnInterface) KeyManagementServiceClient {
	return &keyManagementServiceClient{cc}
}

func (c *keyManagementServiceClient) GenerateKEK(ctx context.Context, in *GenerateKEKRequest, opts ...grpc.CallOption) (*GenerateKEKResponse, error) {
	out := new(GenerateKEKResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.istio.v1.KeyManagementService/GenerateKEK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) GenerateDEK(ctx context.Context, in *GenerateDEKRequest, opts ...grpc.CallOption) (*GenerateDEKResponse, error) {
	out := new(GenerateDEKResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.istio.v1.KeyManagementService/GenerateDEK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) GenerateSKey(ctx context.Context, in *GenerateSKeyRequest, opts ...grpc.CallOption) (*GenerateSKeyResponse, error) {
	out := new(GenerateSKeyResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.istio.v1.KeyManagementService/GenerateSKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) LoadSKey(ctx context.Context, in *LoadSKeyRequest, opts ...grpc.CallOption) (*LoadSKeyResponse, error) {
	out := new(LoadSKeyResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.istio.v1.KeyManagementService/LoadSKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) AuthenticatedEncrypt(ctx context.Context, in *AuthenticatedEncryptRequest, opts ...grpc.CallOption) (*AuthenticatedEncryptResponse, error) {
	out := new(AuthenticatedEncryptResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.istio.v1.KeyManagementService/AuthenticatedEncrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyManagementServiceClient) AuthenticatedDecrypt(ctx context.Context, in *AuthenticatedDecryptRequest, opts ...grpc.CallOption) (*AuthenticatedDecryptResponse, error) {
	out := new(AuthenticatedDecryptResponse)
	err := c.cc.Invoke(ctx, "/thalescpl.io.ekms.istio.v1.KeyManagementService/AuthenticatedDecrypt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyManagementServiceServer is the server API for KeyManagementService service.
type KeyManagementServiceServer interface {
	// GenerateKEK returns the KID of the GeneratedKEK if allowed/successful
	GenerateKEK(context.Context, *GenerateKEKRequest) (*GenerateKEKResponse, error)
	// GenerateDEK returns a wrapped (by HSM handled KEK)
	GenerateDEK(context.Context, *GenerateDEKRequest) (*GenerateDEKResponse, error)
	// GenerateSKey returns a wrapped (by provided encrypted DEK ), for later use during loading and signing key generation
	GenerateSKey(context.Context, *GenerateSKeyRequest) (*GenerateSKeyResponse, error)
	// LoadSKey returns the SKey unwrapped for the controller to use for CA work...
	LoadSKey(context.Context, *LoadSKeyRequest) (*LoadSKeyResponse, error)
	// AuthenticatedEncrypt returns a payload protect by AEAD (Authenticated Encryption with Associated Data) AES-GCM with 256-bit
	AuthenticatedEncrypt(context.Context, *AuthenticatedEncryptRequest) (*AuthenticatedEncryptResponse, error)
	// AuthenticatedDecrypt returns a payload decrypted from the previously invoke AE
	AuthenticatedDecrypt(context.Context, *AuthenticatedDecryptRequest) (*AuthenticatedDecryptResponse, error)
}

// UnimplementedKeyManagementServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeyManagementServiceServer struct {
}

func (*UnimplementedKeyManagementServiceServer) GenerateKEK(context.Context, *GenerateKEKRequest) (*GenerateKEKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKEK not implemented")
}
func (*UnimplementedKeyManagementServiceServer) GenerateDEK(context.Context, *GenerateDEKRequest) (*GenerateDEKResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDEK not implemented")
}
func (*UnimplementedKeyManagementServiceServer) GenerateSKey(context.Context, *GenerateSKeyRequest) (*GenerateSKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateSKey not implemented")
}
func (*UnimplementedKeyManagementServiceServer) LoadSKey(context.Context, *LoadSKeyRequest) (*LoadSKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadSKey not implemented")
}
func (*UnimplementedKeyManagementServiceServer) AuthenticatedEncrypt(context.Context, *AuthenticatedEncryptRequest) (*AuthenticatedEncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticatedEncrypt not implemented")
}
func (*UnimplementedKeyManagementServiceServer) AuthenticatedDecrypt(context.Context, *AuthenticatedDecryptRequest) (*AuthenticatedDecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthenticatedDecrypt not implemented")
}

func RegisterKeyManagementServiceServer(s *grpc.Server, srv KeyManagementServiceServer) {
	s.RegisterService(&_KeyManagementService_serviceDesc, srv)
}

func _KeyManagementService_GenerateKEK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKEKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).GenerateKEK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.istio.v1.KeyManagementService/GenerateKEK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).GenerateKEK(ctx, req.(*GenerateKEKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_GenerateDEK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDEKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).GenerateDEK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.istio.v1.KeyManagementService/GenerateDEK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).GenerateDEK(ctx, req.(*GenerateDEKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_GenerateSKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateSKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).GenerateSKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.istio.v1.KeyManagementService/GenerateSKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).GenerateSKey(ctx, req.(*GenerateSKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_LoadSKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadSKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).LoadSKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.istio.v1.KeyManagementService/LoadSKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).LoadSKey(ctx, req.(*LoadSKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_AuthenticatedEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticatedEncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).AuthenticatedEncrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.istio.v1.KeyManagementService/AuthenticatedEncrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).AuthenticatedEncrypt(ctx, req.(*AuthenticatedEncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyManagementService_AuthenticatedDecrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticatedDecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyManagementServiceServer).AuthenticatedDecrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thalescpl.io.ekms.istio.v1.KeyManagementService/AuthenticatedDecrypt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyManagementServiceServer).AuthenticatedDecrypt(ctx, req.(*AuthenticatedDecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyManagementService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thalescpl.io.ekms.istio.v1.KeyManagementService",
	HandlerType: (*KeyManagementServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKEK",
			Handler:    _KeyManagementService_GenerateKEK_Handler,
		},
		{
			MethodName: "GenerateDEK",
			Handler:    _KeyManagementService_GenerateDEK_Handler,
		},
		{
			MethodName: "GenerateSKey",
			Handler:    _KeyManagementService_GenerateSKey_Handler,
		},
		{
			MethodName: "LoadSKey",
			Handler:    _KeyManagementService_LoadSKey_Handler,
		},
		{
			MethodName: "AuthenticatedEncrypt",
			Handler:    _KeyManagementService_AuthenticatedEncrypt_Handler,
		},
		{
			MethodName: "AuthenticatedDecrypt",
			Handler:    _KeyManagementService_AuthenticatedDecrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apis/istio/v1/service.proto",
}
