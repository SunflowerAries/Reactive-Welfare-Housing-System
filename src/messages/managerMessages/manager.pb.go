// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.0
// source: Reactive-Welfare-Housing-System/src/messages/manager.proto

package managerMessages

import (
	distributorMessages "Reactive-Welfare-Housing-System/src/messages/distributorMessages"
	_ "Reactive-Welfare-Housing-System/src/messages/propertyMessages"
	sharedMessages "Reactive-Welfare-Housing-System/src/messages/sharedMessages"
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

type HouseMatchApprove struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match *distributorMessages.HouseMatch `protobuf:"bytes,1,opt,name=Match,proto3" json:"Match,omitempty"`
}

func (x *HouseMatchApprove) Reset() {
	*x = HouseMatchApprove{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseMatchApprove) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseMatchApprove) ProtoMessage() {}

func (x *HouseMatchApprove) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseMatchApprove.ProtoReflect.Descriptor instead.
func (*HouseMatchApprove) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{0}
}

func (x *HouseMatchApprove) GetMatch() *distributorMessages.HouseMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

type HouseCheckOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckOut *distributorMessages.HouseCheckOut `protobuf:"bytes,1,opt,name=CheckOut,proto3" json:"CheckOut,omitempty"`
}

func (x *HouseCheckOut) Reset() {
	*x = HouseCheckOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseCheckOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseCheckOut) ProtoMessage() {}

func (x *HouseCheckOut) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[1]
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
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{1}
}

func (x *HouseCheckOut) GetCheckOut() *distributorMessages.HouseCheckOut {
	if x != nil {
		return x.CheckOut
	}
	return nil
}

type HouseMatchReject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Match  *distributorMessages.HouseMatch `protobuf:"bytes,1,opt,name=Match,proto3" json:"Match,omitempty"`
	Reason int32                           `protobuf:"varint,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
}

func (x *HouseMatchReject) Reset() {
	*x = HouseMatchReject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseMatchReject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseMatchReject) ProtoMessage() {}

func (x *HouseMatchReject) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseMatchReject.ProtoReflect.Descriptor instead.
func (*HouseMatchReject) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{2}
}

func (x *HouseMatchReject) GetMatch() *distributorMessages.HouseMatch {
	if x != nil {
		return x.Match
	}
	return nil
}

func (x *HouseMatchReject) GetReason() int32 {
	if x != nil {
		return x.Reason
	}
	return 0
}

type NewHouses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Houses      *sharedMessages.NewHouses `protobuf:"bytes,1,opt,name=Houses,proto3" json:"Houses,omitempty"`
	CommitIndex int32                     `protobuf:"varint,2,opt,name=CommitIndex,proto3" json:"CommitIndex,omitempty"`
}

func (x *NewHouses) Reset() {
	*x = NewHouses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewHouses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewHouses) ProtoMessage() {}

func (x *NewHouses) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewHouses.ProtoReflect.Descriptor instead.
func (*NewHouses) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{3}
}

func (x *NewHouses) GetHouses() *sharedMessages.NewHouses {
	if x != nil {
		return x.Houses
	}
	return nil
}

func (x *NewHouses) GetCommitIndex() int32 {
	if x != nil {
		return x.CommitIndex
	}
	return 0
}

type HouseMatchACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HouseMatchACK) Reset() {
	*x = HouseMatchACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseMatchACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseMatchACK) ProtoMessage() {}

func (x *HouseMatchACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseMatchACK.ProtoReflect.Descriptor instead.
func (*HouseMatchACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{4}
}

type HouseCheckOutACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseID int32 `protobuf:"varint,1,opt,name=HouseID,proto3" json:"HouseID,omitempty"`
}

func (x *HouseCheckOutACK) Reset() {
	*x = HouseCheckOutACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseCheckOutACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseCheckOutACK) ProtoMessage() {}

func (x *HouseCheckOutACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[5]
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
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{5}
}

func (x *HouseCheckOutACK) GetHouseID() int32 {
	if x != nil {
		return x.HouseID
	}
	return 0
}

type ExaminationReject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseID  int32 `protobuf:"varint,1,opt,name=HouseID,proto3" json:"HouseID,omitempty"`
	Age      int32 `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Area     int32 `protobuf:"varint,3,opt,name=Area,proto3" json:"Area,omitempty"`
	Level    int32 `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
	FamilyID int32 `protobuf:"varint,5,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
}

func (x *ExaminationReject) Reset() {
	*x = ExaminationReject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExaminationReject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExaminationReject) ProtoMessage() {}

func (x *ExaminationReject) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[6]
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
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{6}
}

func (x *ExaminationReject) GetHouseID() int32 {
	if x != nil {
		return x.HouseID
	}
	return 0
}

func (x *ExaminationReject) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *ExaminationReject) GetArea() int32 {
	if x != nil {
		return x.Area
	}
	return 0
}

func (x *ExaminationReject) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ExaminationReject) GetFamilyID() int32 {
	if x != nil {
		return x.FamilyID
	}
	return 0
}

type UnqualifiedHouses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Houses []*ExaminationReject `protobuf:"bytes,1,rep,name=Houses,proto3" json:"Houses,omitempty"`
}

func (x *UnqualifiedHouses) Reset() {
	*x = UnqualifiedHouses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnqualifiedHouses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnqualifiedHouses) ProtoMessage() {}

func (x *UnqualifiedHouses) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnqualifiedHouses.ProtoReflect.Descriptor instead.
func (*UnqualifiedHouses) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{7}
}

func (x *UnqualifiedHouses) GetHouses() []*ExaminationReject {
	if x != nil {
		return x.Houses
	}
	return nil
}

type UnqualifiedResides struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resides []*ExaminationReject `protobuf:"bytes,1,rep,name=Resides,proto3" json:"Resides,omitempty"`
}

func (x *UnqualifiedResides) Reset() {
	*x = UnqualifiedResides{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnqualifiedResides) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnqualifiedResides) ProtoMessage() {}

func (x *UnqualifiedResides) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnqualifiedResides.ProtoReflect.Descriptor instead.
func (*UnqualifiedResides) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{8}
}

func (x *UnqualifiedResides) GetResides() []*ExaminationReject {
	if x != nil {
		return x.Resides
	}
	return nil
}

type ExaminationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HouseID []int32 `protobuf:"varint,1,rep,packed,name=HouseID,proto3" json:"HouseID,omitempty"`
}

func (x *ExaminationList) Reset() {
	*x = ExaminationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExaminationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExaminationList) ProtoMessage() {}

func (x *ExaminationList) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExaminationList.ProtoReflect.Descriptor instead.
func (*ExaminationList) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP(), []int{9}
}

func (x *ExaminationList) GetHouseID() []int32 {
	if x != nil {
		return x.HouseID
	}
	return nil
}

var File_Reactive_Welfare_Housing_System_src_messages_manager_proto protoreflect.FileDescriptor

var file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61,
	0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x79, 0x6e, 0x6b, 0x72, 0x6f, 0x6e, 0x49, 0x54, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x52,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61, 0x72, 0x65, 0x2d,
	0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x52,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61, 0x72, 0x65, 0x2d,
	0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73,
	0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61, 0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69,
	0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x11, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x52, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x47, 0x0a, 0x0d, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x12, 0x36, 0x0a, 0x08, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x22, 0x59, 0x0a, 0x10, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x05,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x58, 0x0a,
	0x09, 0x4e, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x4e, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x52, 0x06, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x0f, 0x0a, 0x0d, 0x48, 0x6f, 0x75, 0x73, 0x65,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x41, 0x43, 0x4b, 0x22, 0x2c, 0x0a, 0x10, 0x48, 0x6f, 0x75, 0x73,
	0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x41, 0x43, 0x4b, 0x12, 0x18, 0x0a, 0x07,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x45, 0x78, 0x61, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x48,
	0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x65, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x41, 0x72, 0x65, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x22, 0x47,
	0x0a, 0x11, 0x55, 0x6e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x78,
	0x61, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x12, 0x55, 0x6e, 0x71, 0x75, 0x61,
	0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x73, 0x12, 0x34, 0x0a,
	0x07, 0x52, 0x65, 0x73, 0x69, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x07, 0x52, 0x65, 0x73, 0x69,
	0x64, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x45, 0x78, 0x61, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44,
	0x42, 0x3e, 0x5a, 0x3c, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c,
	0x66, 0x61, 0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescOnce sync.Once
	file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescData = file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDesc
)

func file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescGZIP() []byte {
	file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescOnce.Do(func() {
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescData)
	})
	return file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDescData
}

var file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_Reactive_Welfare_Housing_System_src_messages_manager_proto_goTypes = []interface{}{
	(*HouseMatchApprove)(nil),                 // 0: manager.HouseMatchApprove
	(*HouseCheckOut)(nil),                     // 1: manager.HouseCheckOut
	(*HouseMatchReject)(nil),                  // 2: manager.HouseMatchReject
	(*NewHouses)(nil),                         // 3: manager.NewHouses
	(*HouseMatchACK)(nil),                     // 4: manager.HouseMatchACK
	(*HouseCheckOutACK)(nil),                  // 5: manager.HouseCheckOutACK
	(*ExaminationReject)(nil),                 // 6: manager.ExaminationReject
	(*UnqualifiedHouses)(nil),                 // 7: manager.UnqualifiedHouses
	(*UnqualifiedResides)(nil),                // 8: manager.UnqualifiedResides
	(*ExaminationList)(nil),                   // 9: manager.ExaminationList
	(*distributorMessages.HouseMatch)(nil),    // 10: distributor.HouseMatch
	(*distributorMessages.HouseCheckOut)(nil), // 11: distributor.HouseCheckOut
	(*sharedMessages.NewHouses)(nil),          // 12: shared.NewHouses
}
var file_Reactive_Welfare_Housing_System_src_messages_manager_proto_depIdxs = []int32{
	10, // 0: manager.HouseMatchApprove.Match:type_name -> distributor.HouseMatch
	11, // 1: manager.HouseCheckOut.CheckOut:type_name -> distributor.HouseCheckOut
	10, // 2: manager.HouseMatchReject.Match:type_name -> distributor.HouseMatch
	12, // 3: manager.NewHouses.Houses:type_name -> shared.NewHouses
	6,  // 4: manager.UnqualifiedHouses.Houses:type_name -> manager.ExaminationReject
	6,  // 5: manager.UnqualifiedResides.Resides:type_name -> manager.ExaminationReject
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_Reactive_Welfare_Housing_System_src_messages_manager_proto_init() }
func file_Reactive_Welfare_Housing_System_src_messages_manager_proto_init() {
	if File_Reactive_Welfare_Housing_System_src_messages_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseMatchApprove); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseMatchReject); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewHouses); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseMatchACK); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnqualifiedHouses); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnqualifiedResides); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExaminationList); i {
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
			RawDescriptor: file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Reactive_Welfare_Housing_System_src_messages_manager_proto_goTypes,
		DependencyIndexes: file_Reactive_Welfare_Housing_System_src_messages_manager_proto_depIdxs,
		MessageInfos:      file_Reactive_Welfare_Housing_System_src_messages_manager_proto_msgTypes,
	}.Build()
	File_Reactive_Welfare_Housing_System_src_messages_manager_proto = out.File
	file_Reactive_Welfare_Housing_System_src_messages_manager_proto_rawDesc = nil
	file_Reactive_Welfare_Housing_System_src_messages_manager_proto_goTypes = nil
	file_Reactive_Welfare_Housing_System_src_messages_manager_proto_depIdxs = nil
}
