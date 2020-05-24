// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.0
// source: Reactive-Welfare-Housing-System/src/messages/property.proto

package propertyMessages

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

type ExaminedHouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	Age   int32 `protobuf:"varint,3,opt,name=Age,proto3" json:"Age,omitempty"`
	Area  int32 `protobuf:"varint,4,opt,name=Area,proto3" json:"Area,omitempty"`
}

func (x *ExaminedHouse) Reset() {
	*x = ExaminedHouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExaminedHouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExaminedHouse) ProtoMessage() {}

func (x *ExaminedHouse) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExaminedHouse.ProtoReflect.Descriptor instead.
func (*ExaminedHouse) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescGZIP(), []int{0}
}

func (x *ExaminedHouse) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ExaminedHouse) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ExaminedHouse) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *ExaminedHouse) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

type ExaminedHouses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Houses []*ExaminedHouse `protobuf:"bytes,1,rep,name=Houses,proto3" json:"Houses,omitempty"`
}

func (x *ExaminedHouses) Reset() {
	*x = ExaminedHouses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExaminedHouses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExaminedHouses) ProtoMessage() {}

func (x *ExaminedHouses) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExaminedHouses.ProtoReflect.Descriptor instead.
func (*ExaminedHouses) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescGZIP(), []int{1}
}

func (x *ExaminedHouses) GetHouses() []*ExaminedHouse {
	if x != nil {
		return x.Houses
	}
	return nil
}

type ExaminationReject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	House    *ExaminedHouse `protobuf:"bytes,1,opt,name=House,proto3" json:"House,omitempty"`
	FamilyID int32          `protobuf:"varint,2,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
}

func (x *ExaminationReject) Reset() {
	*x = ExaminationReject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExaminationReject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExaminationReject) ProtoMessage() {}

func (x *ExaminationReject) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExaminationReject.ProtoReflect.Descriptor instead.
func (*ExaminationReject) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescGZIP(), []int{2}
}

func (x *ExaminationReject) GetHouse() *ExaminedHouse {
	if x != nil {
		return x.House
	}
	return nil
}

func (x *ExaminationReject) GetFamilyID() int32 {
	if x != nil {
		return x.FamilyID
	}
	return 0
}

type ExaminationRejects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Houses []*ExaminationReject `protobuf:"bytes,1,rep,name=Houses,proto3" json:"Houses,omitempty"`
}

func (x *ExaminationRejects) Reset() {
	*x = ExaminationRejects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExaminationRejects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExaminationRejects) ProtoMessage() {}

func (x *ExaminationRejects) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExaminationRejects.ProtoReflect.Descriptor instead.
func (*ExaminationRejects) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescGZIP(), []int{3}
}

func (x *ExaminationRejects) GetHouses() []*ExaminationReject {
	if x != nil {
		return x.Houses
	}
	return nil
}

var File_Reactive_Welfare_Housing_System_src_messages_property_proto protoreflect.FileDescriptor

var file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61,
	0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x79, 0x6e, 0x6b, 0x72, 0x6f, 0x6e, 0x49, 0x54, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5b, 0x0a, 0x0d, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x22, 0x41, 0x0a, 0x0e,
	0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x2f,
	0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e,
	0x65, 0x64, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x22,
	0x5e, 0x0a, 0x11, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x45,
	0x78, 0x61, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x05, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x22,
	0x49, 0x0a, 0x12, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x42, 0x3f, 0x5a, 0x3d, 0x52, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61, 0x72, 0x65, 0x2d, 0x48,
	0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescOnce sync.Once
	file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescData = file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDesc
)

func file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescGZIP() []byte {
	file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescOnce.Do(func() {
		file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescData = protoimpl.X.CompressGZIP(file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescData)
	})
	return file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDescData
}

var file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Reactive_Welfare_Housing_System_src_messages_property_proto_goTypes = []interface{}{
	(*ExaminedHouse)(nil),      // 0: property.ExaminedHouse
	(*ExaminedHouses)(nil),     // 1: property.ExaminedHouses
	(*ExaminationReject)(nil),  // 2: property.ExaminationReject
	(*ExaminationRejects)(nil), // 3: property.ExaminationRejects
}
var file_Reactive_Welfare_Housing_System_src_messages_property_proto_depIdxs = []int32{
	0, // 0: property.ExaminedHouses.Houses:type_name -> property.ExaminedHouse
	0, // 1: property.ExaminationReject.House:type_name -> property.ExaminedHouse
	2, // 2: property.ExaminationRejects.Houses:type_name -> property.ExaminationReject
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Reactive_Welfare_Housing_System_src_messages_property_proto_init() }
func file_Reactive_Welfare_Housing_System_src_messages_property_proto_init() {
	if File_Reactive_Welfare_Housing_System_src_messages_property_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExaminedHouse); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExaminedHouses); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExaminationReject); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExaminationRejects); i {
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
			RawDescriptor: file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Reactive_Welfare_Housing_System_src_messages_property_proto_goTypes,
		DependencyIndexes: file_Reactive_Welfare_Housing_System_src_messages_property_proto_depIdxs,
		MessageInfos:      file_Reactive_Welfare_Housing_System_src_messages_property_proto_msgTypes,
	}.Build()
	File_Reactive_Welfare_Housing_System_src_messages_property_proto = out.File
	file_Reactive_Welfare_Housing_System_src_messages_property_proto_rawDesc = nil
	file_Reactive_Welfare_Housing_System_src_messages_property_proto_goTypes = nil
	file_Reactive_Welfare_Housing_System_src_messages_property_proto_depIdxs = nil
}
