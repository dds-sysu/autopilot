// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: test_api.proto

package v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
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

type PaintColor struct {
	Color                string   `protobuf:"bytes,1,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PaintColor) Reset()         { *m = PaintColor{} }
func (m *PaintColor) String() string { return proto.CompactTextString(m) }
func (*PaintColor) ProtoMessage()    {}
func (*PaintColor) Descriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{0}
}
func (m *PaintColor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaintColor.Unmarshal(m, b)
}
func (m *PaintColor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaintColor.Marshal(b, m, deterministic)
}
func (m *PaintColor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaintColor.Merge(m, src)
}
func (m *PaintColor) XXX_Size() int {
	return xxx_messageInfo_PaintColor.Size(m)
}
func (m *PaintColor) XXX_DiscardUnknown() {
	xxx_messageInfo_PaintColor.DiscardUnknown(m)
}

var xxx_messageInfo_PaintColor proto.InternalMessageInfo

func (m *PaintColor) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

type TubeStatus struct {
	PercentRemaining     int64    `protobuf:"varint,1,opt,name=percent_remaining,json=percentRemaining,proto3" json:"percent_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TubeStatus) Reset()         { *m = TubeStatus{} }
func (m *TubeStatus) String() string { return proto.CompactTextString(m) }
func (*TubeStatus) ProtoMessage()    {}
func (*TubeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_77683351be7bc655, []int{1}
}
func (m *TubeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TubeStatus.Unmarshal(m, b)
}
func (m *TubeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TubeStatus.Marshal(b, m, deterministic)
}
func (m *TubeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TubeStatus.Merge(m, src)
}
func (m *TubeStatus) XXX_Size() int {
	return xxx_messageInfo_TubeStatus.Size(m)
}
func (m *TubeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_TubeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_TubeStatus proto.InternalMessageInfo

func (m *TubeStatus) GetPercentRemaining() int64 {
	if m != nil {
		return m.PercentRemaining
	}
	return 0
}

func init() {
	proto.RegisterType((*PaintColor)(nil), "things.test.io.PaintColor")
	proto.RegisterType((*TubeStatus)(nil), "things.test.io.TubeStatus")
}

func init() { proto.RegisterFile("test_api.proto", fileDescriptor_77683351be7bc655) }

var fileDescriptor_77683351be7bc655 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8e, 0x31, 0xab, 0xc2, 0x30,
	0x14, 0x85, 0x29, 0x8f, 0xf7, 0xe0, 0x65, 0x28, 0x5a, 0x1c, 0x1c, 0xa5, 0x93, 0x20, 0x26, 0x88,
	0x93, 0x93, 0xa8, 0x7f, 0x40, 0xaa, 0x93, 0x4b, 0x49, 0xe3, 0x25, 0xbd, 0xd0, 0xe6, 0x86, 0xe4,
	0xc6, 0xdf, 0x2f, 0xad, 0x3a, 0xb8, 0x9d, 0xf3, 0x9d, 0x33, 0x7c, 0x22, 0x67, 0x88, 0x5c, 0x6b,
	0x8f, 0xd2, 0x07, 0x62, 0x2a, 0x72, 0x6e, 0xd1, 0xd9, 0x28, 0x07, 0x2c, 0x91, 0xca, 0x52, 0x88,
	0xb3, 0x46, 0xc7, 0x27, 0xea, 0x28, 0x14, 0x33, 0xf1, 0x6b, 0x86, 0x30, 0xcf, 0x16, 0xd9, 0xf2,
	0xbf, 0x7a, 0x95, 0x72, 0x27, 0xc4, 0x35, 0x35, 0x70, 0x61, 0xcd, 0x29, 0x16, 0x2b, 0x31, 0xf5,
	0x10, 0x0c, 0x38, 0xae, 0x03, 0xf4, 0x1a, 0x1d, 0x3a, 0x3b, 0xfe, 0x7f, 0xaa, 0xc9, 0x7b, 0xa8,
	0x3e, 0xfc, 0x78, 0xb8, 0xed, 0x2d, 0x72, 0x9b, 0x1a, 0x69, 0xa8, 0x57, 0x91, 0x3a, 0x5a, 0x23,
	0x29, 0x9d, 0x98, 0x3c, 0x76, 0xc4, 0xca, 0xd0, 0x1d, 0x2c, 0x38, 0x35, 0xe8, 0x28, 0xed, 0x51,
	0x7d, 0xeb, 0xa9, 0xc7, 0xa6, 0xf9, 0x1b, 0xc5, 0xb7, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21,
	0x4a, 0x78, 0x38, 0xca, 0x00, 0x00, 0x00,
}