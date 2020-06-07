// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: Reactive-Welfare-Housing-System/src/messages/distributor.proto

package distributorMessages

import (
	verifierMessages "Reactive-Welfare-Housing-System/src/messages/verifierMessages"
	actor "github.com/AsynkronIT/protoactor-go/actor"
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

type HouseApplicationReject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *verifierMessages.HouseApplicationRequest `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
	Reason  string                                    `protobuf:"bytes,2,opt,name=Reason,proto3" json:"Reason,omitempty"`
}

func (x *HouseApplicationReject) Reset() {
	*x = HouseApplicationReject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseApplicationReject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseApplicationReject) ProtoMessage() {}

func (x *HouseApplicationReject) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseApplicationReject.ProtoReflect.Descriptor instead.
func (*HouseApplicationReject) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{0}
}

func (x *HouseApplicationReject) GetRequest() *verifierMessages.HouseApplicationRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *HouseApplicationReject) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

type NewHousesACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewHousesACK) Reset() {
	*x = NewHousesACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewHousesACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewHousesACK) ProtoMessage() {}

func (x *NewHousesACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewHousesACK.ProtoReflect.Descriptor instead.
func (*NewHousesACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{1}
}

type HouseApplicationACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HouseApplicationACK) Reset() {
	*x = HouseApplicationACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseApplicationACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseApplicationACK) ProtoMessage() {}

func (x *HouseApplicationACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseApplicationACK.ProtoReflect.Descriptor instead.
func (*HouseApplicationACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{2}
}

type MatchEmptyHouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *verifierMessages.HouseApplicationRequest `protobuf:"bytes,1,opt,name=Request,proto3" json:"Request,omitempty"`
	Sender  *actor.PID                                `protobuf:"bytes,2,opt,name=Sender,proto3" json:"Sender,omitempty"`
}

func (x *MatchEmptyHouse) Reset() {
	*x = MatchEmptyHouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchEmptyHouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchEmptyHouse) ProtoMessage() {}

func (x *MatchEmptyHouse) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchEmptyHouse.ProtoReflect.Descriptor instead.
func (*MatchEmptyHouse) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{3}
}

func (x *MatchEmptyHouse) GetRequest() *verifierMessages.HouseApplicationRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *MatchEmptyHouse) GetSender() *actor.PID {
	if x != nil {
		return x.Sender
	}
	return nil
}

type HouseCheckOutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer string `protobuf:"bytes,1,opt,name=Answer,proto3" json:"Answer,omitempty"`
}

func (x *HouseCheckOutResponse) Reset() {
	*x = HouseCheckOutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseCheckOutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseCheckOutResponse) ProtoMessage() {}

func (x *HouseCheckOutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseCheckOutResponse.ProtoReflect.Descriptor instead.
func (*HouseCheckOutResponse) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{4}
}

func (x *HouseCheckOutResponse) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type HouseCheckOut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyID int32 `protobuf:"varint,1,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
	HouseID  int32 `protobuf:"varint,2,opt,name=HouseID,proto3" json:"HouseID,omitempty"`
	Retry    bool  `protobuf:"varint,3,opt,name=Retry,proto3" json:"Retry,omitempty"`
}

func (x *HouseCheckOut) Reset() {
	*x = HouseCheckOut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseCheckOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseCheckOut) ProtoMessage() {}

func (x *HouseCheckOut) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[5]
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
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{5}
}

func (x *HouseCheckOut) GetFamilyID() int32 {
	if x != nil {
		return x.FamilyID
	}
	return 0
}

func (x *HouseCheckOut) GetHouseID() int32 {
	if x != nil {
		return x.HouseID
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
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseCheckOutACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseCheckOutACK) ProtoMessage() {}

func (x *HouseCheckOutACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[6]
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
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{6}
}

type UnqualifiedHousesACK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnqualifiedHousesACK) Reset() {
	*x = UnqualifiedHousesACK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnqualifiedHousesACK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnqualifiedHousesACK) ProtoMessage() {}

func (x *UnqualifiedHousesACK) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnqualifiedHousesACK.ProtoReflect.Descriptor instead.
func (*UnqualifiedHousesACK) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{7}
}

type HouseMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyID int32 `protobuf:"varint,1,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
	HouseID  int32 `protobuf:"varint,2,opt,name=HouseID,proto3" json:"HouseID,omitempty"`
	Level    int32 `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *HouseMatch) Reset() {
	*x = HouseMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HouseMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HouseMatch) ProtoMessage() {}

func (x *HouseMatch) ProtoReflect() protoreflect.Message {
	mi := &file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HouseMatch.ProtoReflect.Descriptor instead.
func (*HouseMatch) Descriptor() ([]byte, []int) {
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP(), []int{8}
}

func (x *HouseMatch) GetFamilyID() int32 {
	if x != nil {
		return x.FamilyID
	}
	return 0
}

func (x *HouseMatch) GetHouseID() int32 {
	if x != nil {
		return x.HouseID
	}
	return 0
}

func (x *HouseMatch) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_Reactive_Welfare_Housing_System_src_messages_distributor_proto protoreflect.FileDescriptor

var file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c, 0x66, 0x61,
	0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x73, 0x79, 0x6e, 0x6b, 0x72,
	0x6f, 0x6e, 0x49, 0x54, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d,
	0x67, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d,
	0x57, 0x65, 0x6c, 0x66, 0x61, 0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x16, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x41, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x22, 0x0e, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x41, 0x43,
	0x4b, 0x22, 0x15, 0x0a, 0x13, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x43, 0x4b, 0x22, 0x72, 0x0a, 0x0f, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x41, 0x70, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x50, 0x49, 0x44, 0x52, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x15,
	0x48, 0x6f, 0x75, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x5b, 0x0a,
	0x0d, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x52, 0x65, 0x74, 0x72, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x41, 0x43, 0x4b, 0x22, 0x16,
	0x0a, 0x14, 0x55, 0x6e, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x66, 0x69, 0x65, 0x64, 0x48, 0x6f, 0x75,
	0x73, 0x65, 0x73, 0x41, 0x43, 0x4b, 0x22, 0x58, 0x0a, 0x0a, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x42, 0x42, 0x5a, 0x40, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x57, 0x65, 0x6c,
	0x66, 0x61, 0x72, 0x65, 0x2d, 0x48, 0x6f, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescOnce sync.Once
	file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescData = file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDesc
)

func file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescGZIP() []byte {
	file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescOnce.Do(func() {
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescData = protoimpl.X.CompressGZIP(file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescData)
	})
	return file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDescData
}

var file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_goTypes = []interface{}{
	(*HouseApplicationReject)(nil),                   // 0: distributor.HouseApplicationReject
	(*NewHousesACK)(nil),                             // 1: distributor.NewHousesACK
	(*HouseApplicationACK)(nil),                      // 2: distributor.HouseApplicationACK
	(*MatchEmptyHouse)(nil),                          // 3: distributor.MatchEmptyHouse
	(*HouseCheckOutResponse)(nil),                    // 4: distributor.HouseCheckOutResponse
	(*HouseCheckOut)(nil),                            // 5: distributor.HouseCheckOut
	(*HouseCheckOutACK)(nil),                         // 6: distributor.HouseCheckOutACK
	(*UnqualifiedHousesACK)(nil),                     // 7: distributor.UnqualifiedHousesACK
	(*HouseMatch)(nil),                               // 8: distributor.HouseMatch
	(*verifierMessages.HouseApplicationRequest)(nil), // 9: verifier.HouseApplicationRequest
	(*actor.PID)(nil),                                // 10: actor.PID
}
var file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_depIdxs = []int32{
	9,  // 0: distributor.HouseApplicationReject.Request:type_name -> verifier.HouseApplicationRequest
	9,  // 1: distributor.MatchEmptyHouse.Request:type_name -> verifier.HouseApplicationRequest
	10, // 2: distributor.MatchEmptyHouse.Sender:type_name -> actor.PID
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_init() }
func file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_init() {
	if File_Reactive_Welfare_Housing_System_src_messages_distributor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseApplicationReject); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewHousesACK); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseApplicationACK); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchEmptyHouse); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseCheckOutResponse); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnqualifiedHousesACK); i {
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
		file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HouseMatch); i {
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
			RawDescriptor: file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_goTypes,
		DependencyIndexes: file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_depIdxs,
		MessageInfos:      file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_msgTypes,
	}.Build()
	File_Reactive_Welfare_Housing_System_src_messages_distributor_proto = out.File
	file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_rawDesc = nil
	file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_goTypes = nil
	file_Reactive_Welfare_Housing_System_src_messages_distributor_proto_depIdxs = nil
}
