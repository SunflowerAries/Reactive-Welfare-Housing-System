// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: Reactive-Welfare-Housing-System/src/messages/verifier.proto

package verifierMessages

import (
	tenantMessages "Reactive-Welfare-Housing-System/src/messages/tenantMessages"
	fmt "fmt"
	_ "github.com/ChaokunChang/protoactor-go/actor"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HouseApplicationRequest struct {
	FamilyID int32 `protobuf:"varint,1,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
	Level    int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	Retry    bool  `protobuf:"varint,3,opt,name=Retry,proto3" json:"Retry,omitempty"`
}

func (m *HouseApplicationRequest) Reset()      { *m = HouseApplicationRequest{} }
func (*HouseApplicationRequest) ProtoMessage() {}
func (*HouseApplicationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d28fb9438c1e6, []int{0}
}
func (m *HouseApplicationRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HouseApplicationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HouseApplicationRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HouseApplicationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseApplicationRequest.Merge(m, src)
}
func (m *HouseApplicationRequest) XXX_Size() int {
	return m.Size()
}
func (m *HouseApplicationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseApplicationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HouseApplicationRequest proto.InternalMessageInfo

func (m *HouseApplicationRequest) GetFamilyID() int32 {
	if m != nil {
		return m.FamilyID
	}
	return 0
}

func (m *HouseApplicationRequest) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *HouseApplicationRequest) GetRetry() bool {
	if m != nil {
		return m.Retry
	}
	return false
}

type HouseApplicationResponse struct {
}

func (m *HouseApplicationResponse) Reset()      { *m = HouseApplicationResponse{} }
func (*HouseApplicationResponse) ProtoMessage() {}
func (*HouseApplicationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d28fb9438c1e6, []int{1}
}
func (m *HouseApplicationResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HouseApplicationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HouseApplicationResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HouseApplicationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseApplicationResponse.Merge(m, src)
}
func (m *HouseApplicationResponse) XXX_Size() int {
	return m.Size()
}
func (m *HouseApplicationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseApplicationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HouseApplicationResponse proto.InternalMessageInfo

type HouseMatchApproveACK struct {
}

func (m *HouseMatchApproveACK) Reset()      { *m = HouseMatchApproveACK{} }
func (*HouseMatchApproveACK) ProtoMessage() {}
func (*HouseMatchApproveACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d28fb9438c1e6, []int{2}
}
func (m *HouseMatchApproveACK) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HouseMatchApproveACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HouseMatchApproveACK.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HouseMatchApproveACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseMatchApproveACK.Merge(m, src)
}
func (m *HouseMatchApproveACK) XXX_Size() int {
	return m.Size()
}
func (m *HouseMatchApproveACK) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseMatchApproveACK.DiscardUnknown(m)
}

var xxx_messageInfo_HouseMatchApproveACK proto.InternalMessageInfo

type HouseMatchRejectACK struct {
}

func (m *HouseMatchRejectACK) Reset()      { *m = HouseMatchRejectACK{} }
func (*HouseMatchRejectACK) ProtoMessage() {}
func (*HouseMatchRejectACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d28fb9438c1e6, []int{3}
}
func (m *HouseMatchRejectACK) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HouseMatchRejectACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HouseMatchRejectACK.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HouseMatchRejectACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseMatchRejectACK.Merge(m, src)
}
func (m *HouseMatchRejectACK) XXX_Size() int {
	return m.Size()
}
func (m *HouseMatchRejectACK) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseMatchRejectACK.DiscardUnknown(m)
}

var xxx_messageInfo_HouseMatchRejectACK proto.InternalMessageInfo

type UnqualifiedResidesACK struct {
}

func (m *UnqualifiedResidesACK) Reset()      { *m = UnqualifiedResidesACK{} }
func (*UnqualifiedResidesACK) ProtoMessage() {}
func (*UnqualifiedResidesACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d28fb9438c1e6, []int{4}
}
func (m *UnqualifiedResidesACK) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnqualifiedResidesACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnqualifiedResidesACK.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnqualifiedResidesACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnqualifiedResidesACK.Merge(m, src)
}
func (m *UnqualifiedResidesACK) XXX_Size() int {
	return m.Size()
}
func (m *UnqualifiedResidesACK) XXX_DiscardUnknown() {
	xxx_messageInfo_UnqualifiedResidesACK.DiscardUnknown(m)
}

var xxx_messageInfo_UnqualifiedResidesACK proto.InternalMessageInfo

type HouseCheckOut struct {
	UserInfo *tenantMessages.HouseApplicationRequest `protobuf:"bytes,1,opt,name=UserInfo,proto3" json:"UserInfo,omitempty"`
	FamilyID int32                                   `protobuf:"varint,2,opt,name=FamilyID,proto3" json:"FamilyID,omitempty"`
	Level    int32                                   `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Retry    bool                                    `protobuf:"varint,4,opt,name=Retry,proto3" json:"Retry,omitempty"`
}

func (m *HouseCheckOut) Reset()      { *m = HouseCheckOut{} }
func (*HouseCheckOut) ProtoMessage() {}
func (*HouseCheckOut) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d28fb9438c1e6, []int{5}
}
func (m *HouseCheckOut) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HouseCheckOut) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HouseCheckOut.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HouseCheckOut) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseCheckOut.Merge(m, src)
}
func (m *HouseCheckOut) XXX_Size() int {
	return m.Size()
}
func (m *HouseCheckOut) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseCheckOut.DiscardUnknown(m)
}

var xxx_messageInfo_HouseCheckOut proto.InternalMessageInfo

func (m *HouseCheckOut) GetUserInfo() *tenantMessages.HouseApplicationRequest {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *HouseCheckOut) GetFamilyID() int32 {
	if m != nil {
		return m.FamilyID
	}
	return 0
}

func (m *HouseCheckOut) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *HouseCheckOut) GetRetry() bool {
	if m != nil {
		return m.Retry
	}
	return false
}

type HouseCheckOutACK struct {
}

func (m *HouseCheckOutACK) Reset()      { *m = HouseCheckOutACK{} }
func (*HouseCheckOutACK) ProtoMessage() {}
func (*HouseCheckOutACK) Descriptor() ([]byte, []int) {
	return fileDescriptor_198d28fb9438c1e6, []int{6}
}
func (m *HouseCheckOutACK) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HouseCheckOutACK) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HouseCheckOutACK.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HouseCheckOutACK) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HouseCheckOutACK.Merge(m, src)
}
func (m *HouseCheckOutACK) XXX_Size() int {
	return m.Size()
}
func (m *HouseCheckOutACK) XXX_DiscardUnknown() {
	xxx_messageInfo_HouseCheckOutACK.DiscardUnknown(m)
}

var xxx_messageInfo_HouseCheckOutACK proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HouseApplicationRequest)(nil), "verifier.HouseApplicationRequest")
	proto.RegisterType((*HouseApplicationResponse)(nil), "verifier.HouseApplicationResponse")
	proto.RegisterType((*HouseMatchApproveACK)(nil), "verifier.HouseMatchApproveACK")
	proto.RegisterType((*HouseMatchRejectACK)(nil), "verifier.HouseMatchRejectACK")
	proto.RegisterType((*UnqualifiedResidesACK)(nil), "verifier.UnqualifiedResidesACK")
	proto.RegisterType((*HouseCheckOut)(nil), "verifier.HouseCheckOut")
	proto.RegisterType((*HouseCheckOutACK)(nil), "verifier.HouseCheckOutACK")
}

func init() {
	proto.RegisterFile("Reactive-Welfare-Housing-System/src/messages/verifier.proto", fileDescriptor_198d28fb9438c1e6)
}

var fileDescriptor_198d28fb9438c1e6 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xb1, 0x6e, 0xd4, 0x30,
	0x18, 0xc7, 0xe3, 0x96, 0xa2, 0xc8, 0x08, 0x09, 0x85, 0x96, 0x46, 0x19, 0x4c, 0x95, 0xa9, 0x4b,
	0x2e, 0x12, 0x2c, 0xa0, 0x8a, 0x21, 0x04, 0x21, 0xaa, 0x52, 0x21, 0x19, 0x55, 0x48, 0x6c, 0xae,
	0xfb, 0x5d, 0x62, 0x9a, 0xd8, 0xa9, 0xed, 0x44, 0xba, 0x8d, 0x47, 0x60, 0xe0, 0x21, 0x78, 0x14,
	0xc6, 0x1b, 0x6f, 0xe4, 0x72, 0x0b, 0xe3, 0x3d, 0x02, 0x4a, 0x72, 0x1c, 0x9c, 0xe0, 0x06, 0xd8,
	0xfc, 0xff, 0x7f, 0xdf, 0x67, 0xff, 0x7f, 0xb6, 0xf1, 0x09, 0x05, 0xc6, 0xad, 0x68, 0x20, 0x7a,
	0x07, 0xc5, 0x98, 0x69, 0x88, 0x5e, 0xa9, 0xda, 0x08, 0x99, 0x45, 0x6f, 0x27, 0xc6, 0x42, 0x19,
	0x1b, 0xcd, 0xe3, 0x12, 0x8c, 0x61, 0x19, 0x98, 0xb8, 0x01, 0x2d, 0xc6, 0x02, 0xf4, 0xa8, 0xd2,
	0xca, 0x2a, 0xcf, 0xfd, 0xa9, 0x83, 0x27, 0x99, 0xb0, 0x79, 0x7d, 0x39, 0xe2, 0xaa, 0x8c, 0xd3,
	0x9c, 0xa9, 0xeb, 0x5a, 0xa6, 0x39, 0x93, 0x59, 0xdc, 0x37, 0x32, 0x6e, 0x95, 0x8e, 0x32, 0x15,
	0xf7, 0x8b, 0xc1, 0x33, 0xc3, 0x1e, 0xc1, 0xd3, 0x7f, 0x0a, 0x60, 0x41, 0x32, 0x69, 0x87, 0xd1,
	0x90, 0xe1, 0xc3, 0xae, 0x17, 0x92, 0xaa, 0x2a, 0x04, 0x67, 0x56, 0x28, 0x49, 0xe1, 0xa6, 0x06,
	0x63, 0xbd, 0x00, 0xbb, 0x2f, 0x59, 0x29, 0x8a, 0xc9, 0xe9, 0x0b, 0x1f, 0x1d, 0xa1, 0xe3, 0x3d,
	0xba, 0xd6, 0xde, 0x3e, 0xde, 0x7b, 0x0d, 0x0d, 0x14, 0xfe, 0x4e, 0x5f, 0x18, 0x44, 0xe7, 0x52,
	0xb0, 0x7a, 0xe2, 0xef, 0x1e, 0xa1, 0x63, 0x97, 0x0e, 0x22, 0x0c, 0xb0, 0xff, 0xe7, 0x11, 0xa6,
	0x52, 0xd2, 0x40, 0xf8, 0x00, 0xef, 0xf7, 0xb5, 0x73, 0x66, 0x79, 0x9e, 0x54, 0x95, 0x56, 0x0d,
	0x24, 0xe9, 0x59, 0x78, 0x80, 0xef, 0xff, 0xf2, 0x29, 0x7c, 0x00, 0x6e, 0x3b, 0xfb, 0x10, 0x1f,
	0x5c, 0xc8, 0x9b, 0x9a, 0x15, 0xdd, 0x8d, 0x5d, 0x51, 0x30, 0xe2, 0x0a, 0x4c, 0x57, 0xf8, 0x8c,
	0xf0, 0xdd, 0x7e, 0x20, 0xcd, 0x81, 0x5f, 0xbf, 0xa9, 0xad, 0x77, 0x82, 0xdd, 0x0b, 0x03, 0xfa,
	0x54, 0x8e, 0x55, 0x9f, 0xfe, 0xce, 0xa3, 0x87, 0xa3, 0x15, 0xf9, 0x16, 0x60, 0xba, 0x1e, 0xd8,
	0x40, 0xdf, 0xd9, 0x86, 0xbe, 0xfb, 0x57, 0xf4, 0x5b, 0xbf, 0xa3, 0x7b, 0xf8, 0xde, 0x46, 0xaa,
	0x24, 0x3d, 0x7b, 0x6e, 0xa6, 0x73, 0xe2, 0xcc, 0xe6, 0xc4, 0x59, 0xce, 0x09, 0xfa, 0xd8, 0x12,
	0xf4, 0xa5, 0x25, 0xe8, 0x6b, 0x4b, 0xd0, 0xb4, 0x25, 0xe8, 0x5b, 0x4b, 0xd0, 0xf7, 0x96, 0x38,
	0xcb, 0x96, 0xa0, 0x4f, 0x0b, 0xe2, 0x4c, 0x17, 0xc4, 0x99, 0x2d, 0x88, 0xf3, 0xfe, 0xd9, 0x7f,
	0xfd, 0xb1, 0xf3, 0x95, 0x71, 0x79, 0xbb, 0x7f, 0xed, 0xc7, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x78, 0x37, 0xb7, 0xe7, 0xab, 0x02, 0x00, 0x00,
}

func (this *HouseApplicationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HouseApplicationRequest)
	if !ok {
		that2, ok := that.(HouseApplicationRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.FamilyID != that1.FamilyID {
		return false
	}
	if this.Level != that1.Level {
		return false
	}
	if this.Retry != that1.Retry {
		return false
	}
	return true
}
func (this *HouseApplicationResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HouseApplicationResponse)
	if !ok {
		that2, ok := that.(HouseApplicationResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *HouseMatchApproveACK) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HouseMatchApproveACK)
	if !ok {
		that2, ok := that.(HouseMatchApproveACK)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *HouseMatchRejectACK) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HouseMatchRejectACK)
	if !ok {
		that2, ok := that.(HouseMatchRejectACK)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *UnqualifiedResidesACK) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UnqualifiedResidesACK)
	if !ok {
		that2, ok := that.(UnqualifiedResidesACK)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *HouseCheckOut) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HouseCheckOut)
	if !ok {
		that2, ok := that.(HouseCheckOut)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UserInfo.Equal(that1.UserInfo) {
		return false
	}
	if this.FamilyID != that1.FamilyID {
		return false
	}
	if this.Level != that1.Level {
		return false
	}
	if this.Retry != that1.Retry {
		return false
	}
	return true
}
func (this *HouseCheckOutACK) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HouseCheckOutACK)
	if !ok {
		that2, ok := that.(HouseCheckOutACK)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *HouseApplicationRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&verifierMessages.HouseApplicationRequest{")
	s = append(s, "FamilyID: "+fmt.Sprintf("%#v", this.FamilyID)+",\n")
	s = append(s, "Level: "+fmt.Sprintf("%#v", this.Level)+",\n")
	s = append(s, "Retry: "+fmt.Sprintf("%#v", this.Retry)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HouseApplicationResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&verifierMessages.HouseApplicationResponse{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HouseMatchApproveACK) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&verifierMessages.HouseMatchApproveACK{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HouseMatchRejectACK) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&verifierMessages.HouseMatchRejectACK{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *UnqualifiedResidesACK) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&verifierMessages.UnqualifiedResidesACK{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HouseCheckOut) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&verifierMessages.HouseCheckOut{")
	if this.UserInfo != nil {
		s = append(s, "UserInfo: "+fmt.Sprintf("%#v", this.UserInfo)+",\n")
	}
	s = append(s, "FamilyID: "+fmt.Sprintf("%#v", this.FamilyID)+",\n")
	s = append(s, "Level: "+fmt.Sprintf("%#v", this.Level)+",\n")
	s = append(s, "Retry: "+fmt.Sprintf("%#v", this.Retry)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HouseCheckOutACK) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&verifierMessages.HouseCheckOutACK{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringVerifier(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *HouseApplicationRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HouseApplicationRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HouseApplicationRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Retry {
		i--
		if m.Retry {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Level != 0 {
		i = encodeVarintVerifier(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x10
	}
	if m.FamilyID != 0 {
		i = encodeVarintVerifier(dAtA, i, uint64(m.FamilyID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *HouseApplicationResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HouseApplicationResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HouseApplicationResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *HouseMatchApproveACK) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HouseMatchApproveACK) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HouseMatchApproveACK) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *HouseMatchRejectACK) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HouseMatchRejectACK) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HouseMatchRejectACK) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *UnqualifiedResidesACK) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnqualifiedResidesACK) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnqualifiedResidesACK) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *HouseCheckOut) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HouseCheckOut) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HouseCheckOut) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Retry {
		i--
		if m.Retry {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.Level != 0 {
		i = encodeVarintVerifier(dAtA, i, uint64(m.Level))
		i--
		dAtA[i] = 0x18
	}
	if m.FamilyID != 0 {
		i = encodeVarintVerifier(dAtA, i, uint64(m.FamilyID))
		i--
		dAtA[i] = 0x10
	}
	if m.UserInfo != nil {
		{
			size, err := m.UserInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVerifier(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HouseCheckOutACK) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HouseCheckOutACK) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HouseCheckOutACK) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintVerifier(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerifier(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HouseApplicationRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FamilyID != 0 {
		n += 1 + sovVerifier(uint64(m.FamilyID))
	}
	if m.Level != 0 {
		n += 1 + sovVerifier(uint64(m.Level))
	}
	if m.Retry {
		n += 2
	}
	return n
}

func (m *HouseApplicationResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *HouseMatchApproveACK) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *HouseMatchRejectACK) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *UnqualifiedResidesACK) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *HouseCheckOut) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UserInfo != nil {
		l = m.UserInfo.Size()
		n += 1 + l + sovVerifier(uint64(l))
	}
	if m.FamilyID != 0 {
		n += 1 + sovVerifier(uint64(m.FamilyID))
	}
	if m.Level != 0 {
		n += 1 + sovVerifier(uint64(m.Level))
	}
	if m.Retry {
		n += 2
	}
	return n
}

func (m *HouseCheckOutACK) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovVerifier(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerifier(x uint64) (n int) {
	return sovVerifier(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HouseApplicationRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HouseApplicationRequest{`,
		`FamilyID:` + fmt.Sprintf("%v", this.FamilyID) + `,`,
		`Level:` + fmt.Sprintf("%v", this.Level) + `,`,
		`Retry:` + fmt.Sprintf("%v", this.Retry) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HouseApplicationResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HouseApplicationResponse{`,
		`}`,
	}, "")
	return s
}
func (this *HouseMatchApproveACK) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HouseMatchApproveACK{`,
		`}`,
	}, "")
	return s
}
func (this *HouseMatchRejectACK) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HouseMatchRejectACK{`,
		`}`,
	}, "")
	return s
}
func (this *UnqualifiedResidesACK) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&UnqualifiedResidesACK{`,
		`}`,
	}, "")
	return s
}
func (this *HouseCheckOut) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HouseCheckOut{`,
		`UserInfo:` + strings.Replace(fmt.Sprintf("%v", this.UserInfo), "HouseApplicationRequest", "tenantMessages.HouseApplicationRequest", 1) + `,`,
		`FamilyID:` + fmt.Sprintf("%v", this.FamilyID) + `,`,
		`Level:` + fmt.Sprintf("%v", this.Level) + `,`,
		`Retry:` + fmt.Sprintf("%v", this.Retry) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HouseCheckOutACK) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HouseCheckOutACK{`,
		`}`,
	}, "")
	return s
}
func valueToStringVerifier(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HouseApplicationRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HouseApplicationRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HouseApplicationRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FamilyID", wireType)
			}
			m.FamilyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FamilyID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Retry", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Retry = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HouseApplicationResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HouseApplicationResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HouseApplicationResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HouseMatchApproveACK) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HouseMatchApproveACK: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HouseMatchApproveACK: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HouseMatchRejectACK) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HouseMatchRejectACK: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HouseMatchRejectACK: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnqualifiedResidesACK) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnqualifiedResidesACK: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnqualifiedResidesACK: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HouseCheckOut) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HouseCheckOut: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HouseCheckOut: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVerifier
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerifier
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.UserInfo == nil {
				m.UserInfo = &tenantMessages.HouseApplicationRequest{}
			}
			if err := m.UserInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FamilyID", wireType)
			}
			m.FamilyID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FamilyID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Level", wireType)
			}
			m.Level = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Level |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Retry", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Retry = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HouseCheckOutACK) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HouseCheckOutACK: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HouseCheckOutACK: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipVerifier(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVerifier
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipVerifier(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerifier
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowVerifier
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthVerifier
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerifier
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerifier
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerifier        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerifier          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerifier = fmt.Errorf("proto: unexpected end of group")
)
