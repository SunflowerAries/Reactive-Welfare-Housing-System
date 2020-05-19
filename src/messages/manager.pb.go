// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.0
// source: manager.proto

package messages

import (
	_ "github.com/AsynkronIT/protoactor-go/actor"
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

type Houses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Houses []*House `protobuf:"bytes,1,rep,name=houses,proto3" json:"houses,omitempty"`
}

func (x *Houses) Reset() {
	*x = Houses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Houses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Houses) ProtoMessage() {}

func (x *Houses) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Houses.ProtoReflect.Descriptor instead.
func (*Houses) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{0}
}

func (x *Houses) GetHouses() []*House {
	if x != nil {
		return x.Houses
	}
	return nil
}

type House struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level int32 `protobuf:"varint,1,opt,name=Level,proto3" json:"Level,omitempty"`
	Age   int32 `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Area  int32 `protobuf:"varint,3,opt,name=Area,proto3" json:"Area,omitempty"`
}

func (x *House) Reset() {
	*x = House{}
	if protoimpl.UnsafeEnabled {
		mi := &file_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *House) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*House) ProtoMessage() {}

func (x *House) ProtoReflect() protoreflect.Message {
	mi := &file_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use House.ProtoReflect.Descriptor instead.
func (*House) Descriptor() ([]byte, []int) {
	return file_manager_proto_rawDescGZIP(), []int{1}
}

func (x *House) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *House) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *House) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

var File_manager_proto protoreflect.FileDescriptor

var file_manager_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x79, 0x6e, 0x6b, 0x72, 0x6f, 0x6e, 0x49, 0x54,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x31, 0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x06, 0x68,
	0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x06, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_manager_proto_rawDescOnce sync.Once
	file_manager_proto_rawDescData = file_manager_proto_rawDesc
)

func file_manager_proto_rawDescGZIP() []byte {
	file_manager_proto_rawDescOnce.Do(func() {
		file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_manager_proto_rawDescData)
	})
	return file_manager_proto_rawDescData
}

var file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_manager_proto_goTypes = []interface{}{
	(*Houses)(nil), // 0: messages.Houses
	(*House)(nil),  // 1: messages.House
}
var file_manager_proto_depIdxs = []int32{
	1, // 0: messages.Houses.houses:type_name -> messages.House
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_manager_proto_init() }
func file_manager_proto_init() {
	if File_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Houses); i {
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
		file_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*House); i {
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
			RawDescriptor: file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_manager_proto_goTypes,
		DependencyIndexes: file_manager_proto_depIdxs,
		MessageInfos:      file_manager_proto_msgTypes,
	}.Build()
	File_manager_proto = out.File
	file_manager_proto_rawDesc = nil
	file_manager_proto_goTypes = nil
	file_manager_proto_depIdxs = nil
}
