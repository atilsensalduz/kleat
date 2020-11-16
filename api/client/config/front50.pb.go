// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: config/front50.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	storage "github.com/spinnaker/kleat/api/client/storage"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configuration for the front50 microservice.
type Front50 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spinnaker *Front50_Spinnaker `protobuf:"bytes,1,opt,name=spinnaker,proto3" json:"spinnaker,omitempty"`
}

func (x *Front50) Reset() {
	*x = Front50{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_front50_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Front50) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Front50) ProtoMessage() {}

func (x *Front50) ProtoReflect() protoreflect.Message {
	mi := &file_config_front50_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Front50.ProtoReflect.Descriptor instead.
func (*Front50) Descriptor() ([]byte, []int) {
	return file_config_front50_proto_rawDescGZIP(), []int{0}
}

func (x *Front50) GetSpinnaker() *Front50_Spinnaker {
	if x != nil {
		return x.Spinnaker
	}
	return nil
}

type Front50_Spinnaker struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gcs    *storage.Gcs    `protobuf:"bytes,1,opt,name=gcs,proto3" json:"gcs,omitempty"`
	Azs    *storage.Azs    `protobuf:"bytes,2,opt,name=azs,proto3" json:"azs,omitempty"`
	Oracle *storage.Oracle `protobuf:"bytes,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
	S3     *storage.S3     `protobuf:"bytes,4,opt,name=s3,proto3" json:"s3,omitempty"`
}

func (x *Front50_Spinnaker) Reset() {
	*x = Front50_Spinnaker{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_front50_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Front50_Spinnaker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Front50_Spinnaker) ProtoMessage() {}

func (x *Front50_Spinnaker) ProtoReflect() protoreflect.Message {
	mi := &file_config_front50_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Front50_Spinnaker.ProtoReflect.Descriptor instead.
func (*Front50_Spinnaker) Descriptor() ([]byte, []int) {
	return file_config_front50_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Front50_Spinnaker) GetGcs() *storage.Gcs {
	if x != nil {
		return x.Gcs
	}
	return nil
}

func (x *Front50_Spinnaker) GetAzs() *storage.Azs {
	if x != nil {
		return x.Azs
	}
	return nil
}

func (x *Front50_Spinnaker) GetOracle() *storage.Oracle {
	if x != nil {
		return x.Oracle
	}
	return nil
}

func (x *Front50_Spinnaker) GetS3() *storage.S3 {
	if x != nil {
		return x.S3
	}
	return nil
}

var File_config_front50_proto protoreflect.FileDescriptor

var file_config_front50_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x35, 0x30,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x7a,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x67, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x10, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf4, 0x01, 0x0a, 0x07, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x35, 0x30, 0x12, 0x3d,
	0x0a, 0x09, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x35, 0x30, 0x2e, 0x53, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b,
	0x65, 0x72, 0x52, 0x09, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x1a, 0xa9, 0x01,
	0x0a, 0x09, 0x53, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x03, 0x67,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x63, 0x73, 0x52, 0x03, 0x67, 0x63,
	0x73, 0x12, 0x24, 0x0a, 0x03, 0x61, 0x7a, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x41,
	0x7a, 0x73, 0x52, 0x03, 0x61, 0x7a, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x06,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x33, 0x52, 0x02, 0x73, 0x33, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65,
	0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_front50_proto_rawDescOnce sync.Once
	file_config_front50_proto_rawDescData = file_config_front50_proto_rawDesc
)

func file_config_front50_proto_rawDescGZIP() []byte {
	file_config_front50_proto_rawDescOnce.Do(func() {
		file_config_front50_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_front50_proto_rawDescData)
	})
	return file_config_front50_proto_rawDescData
}

var file_config_front50_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_front50_proto_goTypes = []interface{}{
	(*Front50)(nil),           // 0: proto.config.Front50
	(*Front50_Spinnaker)(nil), // 1: proto.config.Front50.Spinnaker
	(*storage.Gcs)(nil),       // 2: proto.storage.Gcs
	(*storage.Azs)(nil),       // 3: proto.storage.Azs
	(*storage.Oracle)(nil),    // 4: proto.storage.Oracle
	(*storage.S3)(nil),        // 5: proto.storage.S3
}
var file_config_front50_proto_depIdxs = []int32{
	1, // 0: proto.config.Front50.spinnaker:type_name -> proto.config.Front50.Spinnaker
	2, // 1: proto.config.Front50.Spinnaker.gcs:type_name -> proto.storage.Gcs
	3, // 2: proto.config.Front50.Spinnaker.azs:type_name -> proto.storage.Azs
	4, // 3: proto.config.Front50.Spinnaker.oracle:type_name -> proto.storage.Oracle
	5, // 4: proto.config.Front50.Spinnaker.s3:type_name -> proto.storage.S3
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_front50_proto_init() }
func file_config_front50_proto_init() {
	if File_config_front50_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_front50_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Front50); i {
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
		file_config_front50_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Front50_Spinnaker); i {
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
			RawDescriptor: file_config_front50_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_front50_proto_goTypes,
		DependencyIndexes: file_config_front50_proto_depIdxs,
		MessageInfos:      file_config_front50_proto_msgTypes,
	}.Build()
	File_config_front50_proto = out.File
	file_config_front50_proto_rawDesc = nil
	file_config_front50_proto_goTypes = nil
	file_config_front50_proto_depIdxs = nil
}
