// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: artifact/oracle.proto

package artifact

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
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

// Configuration for the Oracle artifact provider.
type Oracle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the Oracle artifact provider is enabled.
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Oracle artifact accounts.
	Accounts []*OracleAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *Oracle) Reset() {
	*x = Oracle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_oracle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oracle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oracle) ProtoMessage() {}

func (x *Oracle) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_oracle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oracle.ProtoReflect.Descriptor instead.
func (*Oracle) Descriptor() ([]byte, []int) {
	return file_artifact_oracle_proto_rawDescGZIP(), []int{0}
}

func (x *Oracle) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Oracle) GetAccounts() []*OracleAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for an Oracle artifact account.
type OracleAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The fingerprint of the public key.
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// The namespace in which the bucket and objects will be created.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// The passphrase used for the private key, if it is encrypted.
	PrivateKeyPassphrase string `protobuf:"bytes,4,opt,name=privateKeyPassphrase,proto3" json:"privateKeyPassphrase,omitempty"`
	// An Oracle region (for example, `us-phoenix-1`).
	Region string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
	// Path to the private key in PEM format.
	SshPrivateKeyFilePath string `protobuf:"bytes,6,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// The OCID of the Oracle Tenancy to use.
	TenancyId string `protobuf:"bytes,7,opt,name=tenancyId,proto3" json:"tenancyId,omitempty"`
	// The OCID of the Oracle User with which to authenticate.
	UserId string `protobuf:"bytes,8,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *OracleAccount) Reset() {
	*x = OracleAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_oracle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleAccount) ProtoMessage() {}

func (x *OracleAccount) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_oracle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleAccount.ProtoReflect.Descriptor instead.
func (*OracleAccount) Descriptor() ([]byte, []int) {
	return file_artifact_oracle_proto_rawDescGZIP(), []int{1}
}

func (x *OracleAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OracleAccount) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *OracleAccount) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *OracleAccount) GetPrivateKeyPassphrase() string {
	if x != nil {
		return x.PrivateKeyPassphrase
	}
	return ""
}

func (x *OracleAccount) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *OracleAccount) GetSshPrivateKeyFilePath() string {
	if x != nil {
		return x.SshPrivateKeyFilePath
	}
	return ""
}

func (x *OracleAccount) GetTenancyId() string {
	if x != nil {
		return x.TenancyId
	}
	return ""
}

func (x *OracleAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_artifact_oracle_proto protoreflect.FileDescriptor

var file_artifact_oracle_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x06, 0x4f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x9b, 0x02, 0x0a, 0x0d, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x65, 0x6e, 0x61, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_oracle_proto_rawDescOnce sync.Once
	file_artifact_oracle_proto_rawDescData = file_artifact_oracle_proto_rawDesc
)

func file_artifact_oracle_proto_rawDescGZIP() []byte {
	file_artifact_oracle_proto_rawDescOnce.Do(func() {
		file_artifact_oracle_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_oracle_proto_rawDescData)
	})
	return file_artifact_oracle_proto_rawDescData
}

var file_artifact_oracle_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_artifact_oracle_proto_goTypes = []interface{}{
	(*Oracle)(nil),               // 0: proto.artifact.Oracle
	(*OracleAccount)(nil),        // 1: proto.artifact.OracleAccount
	(*wrapperspb.BoolValue)(nil), // 2: google.protobuf.BoolValue
}
var file_artifact_oracle_proto_depIdxs = []int32{
	2, // 0: proto.artifact.Oracle.enabled:type_name -> google.protobuf.BoolValue
	1, // 1: proto.artifact.Oracle.accounts:type_name -> proto.artifact.OracleAccount
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_artifact_oracle_proto_init() }
func file_artifact_oracle_proto_init() {
	if File_artifact_oracle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_oracle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oracle); i {
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
		file_artifact_oracle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleAccount); i {
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
			RawDescriptor: file_artifact_oracle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_oracle_proto_goTypes,
		DependencyIndexes: file_artifact_oracle_proto_depIdxs,
		MessageInfos:      file_artifact_oracle_proto_msgTypes,
	}.Build()
	File_artifact_oracle_proto = out.File
	file_artifact_oracle_proto_rawDesc = nil
	file_artifact_oracle_proto_goTypes = nil
	file_artifact_oracle_proto_depIdxs = nil
}
