// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: canary/signalfx.proto

package canary

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

// Configuration for the SignalFx canary integration.
type SignalFx struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether SignalFx is enabled as a metric store provider.
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*SignalFxAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *SignalFx) Reset() {
	*x = SignalFx{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_signalfx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalFx) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalFx) ProtoMessage() {}

func (x *SignalFx) ProtoReflect() protoreflect.Message {
	mi := &file_canary_signalfx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalFx.ProtoReflect.Descriptor instead.
func (*SignalFx) Descriptor() ([]byte, []int) {
	return file_canary_signalfx_proto_rawDescGZIP(), []int{0}
}

func (x *SignalFx) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *SignalFx) GetAccounts() []*SignalFxAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type SignalFxAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The SignalFx access token.
	AccessToken string `protobuf:"bytes,2,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	// The SignalFx server endpoint.
	Endpoint *SignalFxAccount_Endpoint `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// The scope key, which is used to distinguish between base and canary
	// deployments. If omitted, each request must supply the `_scope_key` param
	// in extended scope params.
	DefaultScopeKey string `protobuf:"bytes,4,opt,name=defaultScopeKey,proto3" json:"defaultScopeKey,omitempty"`
	// The location key, which is used to filter by deployment region. If
	// omitted, each request must supply the `_location_key` if it is needed.
	DefaultLocationKey string `protobuf:"bytes,5,opt,name=defaultLocationKey,proto3" json:"defaultLocationKey,omitempty"`
}

func (x *SignalFxAccount) Reset() {
	*x = SignalFxAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_signalfx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalFxAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalFxAccount) ProtoMessage() {}

func (x *SignalFxAccount) ProtoReflect() protoreflect.Message {
	mi := &file_canary_signalfx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalFxAccount.ProtoReflect.Descriptor instead.
func (*SignalFxAccount) Descriptor() ([]byte, []int) {
	return file_canary_signalfx_proto_rawDescGZIP(), []int{1}
}

func (x *SignalFxAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SignalFxAccount) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignalFxAccount) GetEndpoint() *SignalFxAccount_Endpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *SignalFxAccount) GetDefaultScopeKey() string {
	if x != nil {
		return x.DefaultScopeKey
	}
	return ""
}

func (x *SignalFxAccount) GetDefaultLocationKey() string {
	if x != nil {
		return x.DefaultLocationKey
	}
	return ""
}

// The SignalFx server endpoint.
type SignalFxAccount_Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The base URL to the SignalFx server. Defaults to
	// https://stream.signalfx.com
	BaseUrl string `protobuf:"bytes,1,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
}

func (x *SignalFxAccount_Endpoint) Reset() {
	*x = SignalFxAccount_Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_canary_signalfx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignalFxAccount_Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignalFxAccount_Endpoint) ProtoMessage() {}

func (x *SignalFxAccount_Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_canary_signalfx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignalFxAccount_Endpoint.ProtoReflect.Descriptor instead.
func (*SignalFxAccount_Endpoint) Descriptor() ([]byte, []int) {
	return file_canary_signalfx_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SignalFxAccount_Endpoint) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

var File_canary_signalfx_proto protoreflect.FileDescriptor

var file_canary_signalfx_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x66,
	0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x61, 0x6e, 0x61, 0x72, 0x79, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x46,
	0x78, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x46,
	0x78, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x8b, 0x02, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x46, 0x78, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x42, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x46, 0x78, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x4b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2e, 0x0a, 0x12, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x1a, 0x24, 0x0a, 0x08, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x61, 0x6e, 0x61, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_canary_signalfx_proto_rawDescOnce sync.Once
	file_canary_signalfx_proto_rawDescData = file_canary_signalfx_proto_rawDesc
)

func file_canary_signalfx_proto_rawDescGZIP() []byte {
	file_canary_signalfx_proto_rawDescOnce.Do(func() {
		file_canary_signalfx_proto_rawDescData = protoimpl.X.CompressGZIP(file_canary_signalfx_proto_rawDescData)
	})
	return file_canary_signalfx_proto_rawDescData
}

var file_canary_signalfx_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_canary_signalfx_proto_goTypes = []interface{}{
	(*SignalFx)(nil),                 // 0: proto.canary.SignalFx
	(*SignalFxAccount)(nil),          // 1: proto.canary.SignalFxAccount
	(*SignalFxAccount_Endpoint)(nil), // 2: proto.canary.SignalFxAccount.Endpoint
	(*wrapperspb.BoolValue)(nil),     // 3: google.protobuf.BoolValue
}
var file_canary_signalfx_proto_depIdxs = []int32{
	3, // 0: proto.canary.SignalFx.enabled:type_name -> google.protobuf.BoolValue
	1, // 1: proto.canary.SignalFx.accounts:type_name -> proto.canary.SignalFxAccount
	2, // 2: proto.canary.SignalFxAccount.endpoint:type_name -> proto.canary.SignalFxAccount.Endpoint
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_canary_signalfx_proto_init() }
func file_canary_signalfx_proto_init() {
	if File_canary_signalfx_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_canary_signalfx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalFx); i {
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
		file_canary_signalfx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalFxAccount); i {
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
		file_canary_signalfx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignalFxAccount_Endpoint); i {
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
			RawDescriptor: file_canary_signalfx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_canary_signalfx_proto_goTypes,
		DependencyIndexes: file_canary_signalfx_proto_depIdxs,
		MessageInfos:      file_canary_signalfx_proto_msgTypes,
	}.Build()
	File_canary_signalfx_proto = out.File
	file_canary_signalfx_proto_rawDesc = nil
	file_canary_signalfx_proto_goTypes = nil
	file_canary_signalfx_proto_depIdxs = nil
}
