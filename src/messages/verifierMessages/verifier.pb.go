// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: Reactive-Welfare-Housing-System/src/messages/verifier.proto

package verifierMessages

import (
	tenantMessages "Reactive-Welfare-Housing-System/src/messages/tenantMessages"
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

type HouseApplicationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyID int32 `protobuf:"varint,1,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
	Level    int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	Retry    bool  `protobuf:"varint,3,opt,name=Retry,proto3" json:"Retry,omitempty"`
}

func (x *HouseApplicationRequest) Reset() {
	*x = HouseApplicationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseApplicationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseApplicationRequest) ProtoMessage() {}

func (x *HouseApplicationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseApplicationRequest.ProtoReflect.Descriptor instead.
func (*HouseApplicationRequest) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP(), []int{0}
}

func (x *HouseApplicationRequest) GetFamilyID() int32 {
	if x != nil {
		return x.FamilyID
	}
	return 0
}

func (x *HouseApplicationRequest) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *HouseApplicationRequest) GetRetry() bool {
	if x != nil {
		return x.Retry
	}
	return false
}

type HouseApplicationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HouseApplicationResponse) Reset() {
	*x = HouseApplicationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseApplicationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseApplicationResponse) ProtoMessage() {}

func (x *HouseApplicationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseApplicationResponse.ProtoReflect.Descriptor instead.
func (*HouseApplicationResponse) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP(), []int{1}
}

type HouseMatchApproveACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HouseMatchApproveACK) Reset() {
	*x = HouseMatchApproveACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseMatchApproveACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseMatchApproveACK) ProtoMessage() {}

func (x *HouseMatchApproveACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseMatchApproveACK.ProtoReflect.Descriptor instead.
func (*HouseMatchApproveACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP(), []int{2}
}

type HouseMatchRejectACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HouseMatchRejectACK) Reset() {
	*x = HouseMatchRejectACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseMatchRejectACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseMatchRejectACK) ProtoMessage() {}

func (x *HouseMatchRejectACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseMatchRejectACK.ProtoReflect.Descriptor instead.
func (*HouseMatchRejectACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP(), []int{3}
}

type UnqualifiedResidesACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnqualifiedResidesACK) Reset() {
	*x = UnqualifiedResidesACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnqualifiedResidesACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnqualifiedResidesACK) ProtoMessage() {}

func (x *UnqualifiedResidesACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnqualifiedResidesACK.ProtoReflect.Descriptor instead.
func (*UnqualifiedResidesACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP(), []int{4}
}

type HouseCheckOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *tenantMessages.HouseApplicationRequest `protobuf:"bytes,1,opt,name=UserInfo,proto3" json:"UserInfo,omitempty"`
	FamilyID int32                                   `protobuf:"varint,2,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
	Level    int32                                   `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Retry    bool                                    `protobuf:"varint,4,opt,name=Retry,proto3" json:"Retry,omitempty"`
}

func (x *HouseCheckOut) Reset() {
	*x = HouseCheckOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseCheckOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseCheckOut) ProtoMessage() {}

func (x *HouseCheckOut) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseCheckOut.ProtoReflect.Descriptor instead.
func (*HouseCheckOut) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP(), []int{5}
}

func (x *HouseCheckOut) GetUserInfo() *tenantMessages.HouseApplicationRequest {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *HouseCheckOut) GetFamilyID() int32 {
	if x != nil {
		return x.FamilyID
	}
	return 0
}

func (x *HouseCheckOut) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *HouseCheckOut) GetRetry() bool {
	if x != nil {
		return x.Retry
	}
	return false
}

type HouseCheckOutACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HouseCheckOutACK) Reset() {
	*x = HouseCheckOutACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseCheckOutACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseCheckOutACK) ProtoMessage() {}

func (x *HouseCheckOutACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseCheckOutACK.ProtoReflect.Descriptor instead.
func (*HouseCheckOutACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP(), []int{6}
}

var File_Reactive_Welfare_Housing_System_src_messages_verifier_proto protoreflect.FileDescriptor

var file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61,
	0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x79, 0x6e, 0x6b, 0x72, 0x6f, 0x6e, 0x49, 0x54, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x39, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61, 0x72,
	0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x17, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x22, 0x1a, 0x0a,
	0x18, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x41, 0x43,
	0x4b, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x41, 0x43, 0x4b, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x6e, 0x71, 0x75,
	0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x73, 0x41, 0x43,
	0x4b, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x2e, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x41, 0x43, 0x4b, 0x42, 0x3f, 0x5a, 0x3d,
	0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61, 0x72, 0x65,
	0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescOnce sync.Once
	file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescData = file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDesc
)

func file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescGZIP() []byte {
	file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescOnce.Do(func() {
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescData)
	})
	return file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDescData
}

var file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_goTypes = []interface{}{
	(*HouseApplicationRequest)(nil),                // 0: verifier.HouseApplicationRequest
	(*HouseApplicationResponse)(nil),               // 1: verifier.HouseApplicationResponse
	(*HouseMatchApproveACK)(nil),                   // 2: verifier.HouseMatchApproveACK
	(*HouseMatchRejectACK)(nil),                    // 3: verifier.HouseMatchRejectACK
	(*UnqualifiedResidesACK)(nil),                  // 4: verifier.UnqualifiedResidesACK
	(*HouseCheckOut)(nil),                          // 5: verifier.HouseCheckOut
	(*HouseCheckOutACK)(nil),                       // 6: verifier.HouseCheckOutACK
	(*tenantMessages.HouseApplicationRequest)(nil), // 7: tenant.HouseApplicationRequest
}
var file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_depIdxs = []int32{
	7, // 0: verifier.HouseCheckOut.UserInfo:type_name -> tenant.HouseApplicationRequest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_init() }
func file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_init() {
	if File_Reactive_Welfare_Housing_System_src_messages_verifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseApplicationRequest); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseApplicationResponse); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseMatchApproveACK); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseMatchRejectACK); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnqualifiedResidesACK); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseCheckOut); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseCheckOutACK); i {
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
			RawDescriptor: file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_goTypes,
		DependencyIndexes: file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_depIdxs,
		MessageInfos:      file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_msgTypes,
	}.Build()
	File_Reactive_Welfare_Housing_System_src_messages_verifier_proto = out.File
	file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_rawDesc = nil
	file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_goTypes = nil
	file_Reactive_Welfare_Housing_System_src_messages_verifier_proto_depIdxs = nil
}
