// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: genesis.v1.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

type GenesisCurrent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha256      []byte `protobuf:"bytes,1,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Commitmsg   string `protobuf:"bytes,2,opt,name=commitmsg,proto3" json:"commitmsg,omitempty"`
	Stable      bool   `protobuf:"varint,3,opt,name=stable,proto3" json:"stable,omitempty"`
	Contentsize uint64 `protobuf:"varint,4,opt,name=contentsize,proto3" json:"contentsize,omitempty"`
}

func (x *GenesisCurrent) Reset() {
	*x = GenesisCurrent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genesis_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisCurrent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisCurrent) ProtoMessage() {}

func (x *GenesisCurrent) ProtoReflect() protoreflect.Message {
	mi := &file_genesis_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisCurrent.ProtoReflect.Descriptor instead.
func (*GenesisCurrent) Descriptor() ([]byte, []int) {
	return file_genesis_v1_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisCurrent) GetSha256() []byte {
	if x != nil {
		return x.Sha256
	}
	return nil
}

func (x *GenesisCurrent) GetCommitmsg() string {
	if x != nil {
		return x.Commitmsg
	}
	return ""
}

func (x *GenesisCurrent) GetStable() bool {
	if x != nil {
		return x.Stable
	}
	return false
}

func (x *GenesisCurrent) GetContentsize() uint64 {
	if x != nil {
		return x.Contentsize
	}
	return 0
}

type GenesisUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha256      []byte `protobuf:"bytes,1,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Commitmsg   string `protobuf:"bytes,2,opt,name=commitmsg,proto3" json:"commitmsg,omitempty"`
	Parent      []byte `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Contentsize uint64 `protobuf:"varint,4,opt,name=contentsize,proto3" json:"contentsize,omitempty"`
}

func (x *GenesisUpdate) Reset() {
	*x = GenesisUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_genesis_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisUpdate) ProtoMessage() {}

func (x *GenesisUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_genesis_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisUpdate.ProtoReflect.Descriptor instead.
func (*GenesisUpdate) Descriptor() ([]byte, []int) {
	return file_genesis_v1_proto_rawDescGZIP(), []int{1}
}

func (x *GenesisUpdate) GetSha256() []byte {
	if x != nil {
		return x.Sha256
	}
	return nil
}

func (x *GenesisUpdate) GetCommitmsg() string {
	if x != nil {
		return x.Commitmsg
	}
	return ""
}

func (x *GenesisUpdate) GetParent() []byte {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *GenesisUpdate) GetContentsize() uint64 {
	if x != nil {
		return x.Contentsize
	}
	return 0
}

var File_genesis_v1_proto protoreflect.FileDescriptor

var file_genesis_v1_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x76, 0x32, 0x22, 0x80,
	0x01, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x7a,
	0x65, 0x22, 0x7f, 0x0a, 0x0d, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f,
	0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x69,
	0x7a, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_genesis_v1_proto_rawDescOnce sync.Once
	file_genesis_v1_proto_rawDescData = file_genesis_v1_proto_rawDesc
)

func file_genesis_v1_proto_rawDescGZIP() []byte {
	file_genesis_v1_proto_rawDescOnce.Do(func() {
		file_genesis_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_genesis_v1_proto_rawDescData)
	})
	return file_genesis_v1_proto_rawDescData
}

var file_genesis_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_genesis_v1_proto_goTypes = []interface{}{
	(*GenesisCurrent)(nil), // 0: genesis.v2.GenesisCurrent
	(*GenesisUpdate)(nil),  // 1: genesis.v2.GenesisUpdate
}
var file_genesis_v1_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_genesis_v1_proto_init() }
func file_genesis_v1_proto_init() {
	if File_genesis_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_genesis_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisCurrent); i {
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
		file_genesis_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisUpdate); i {
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
			RawDescriptor: file_genesis_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_genesis_v1_proto_goTypes,
		DependencyIndexes: file_genesis_v1_proto_depIdxs,
		MessageInfos:      file_genesis_v1_proto_msgTypes,
	}.Build()
	File_genesis_v1_proto = out.File
	file_genesis_v1_proto_rawDesc = nil
	file_genesis_v1_proto_goTypes = nil
	file_genesis_v1_proto_depIdxs = nil
}
