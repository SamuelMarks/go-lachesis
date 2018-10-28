// Code generated by protoc-gen-go. DO NOT EDIT.
// source: frame.proto

package poset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Frame struct {
	Round                int64    `protobuf:"varint,1,opt,name=Round,proto3" json:"Round,omitempty"`
	Roots                []*Root  `protobuf:"bytes,2,rep,name=Roots,proto3" json:"Roots,omitempty"`
	Events               []*Event `protobuf:"bytes,3,rep,name=Events,proto3" json:"Events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Frame) Reset()         { *m = Frame{} }
func (m *Frame) String() string { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()    {}
func (*Frame) Descriptor() ([]byte, []int) {
	return fileDescriptor_5379e2b825e15002, []int{0}
}

func (m *Frame) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Frame.Unmarshal(m, b)
}
func (m *Frame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Frame.Marshal(b, m, deterministic)
}
func (m *Frame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Frame.Merge(m, src)
}
func (m *Frame) XXX_Size() int {
	return xxx_messageInfo_Frame.Size(m)
}
func (m *Frame) XXX_DiscardUnknown() {
	xxx_messageInfo_Frame.DiscardUnknown(m)
}

var xxx_messageInfo_Frame proto.InternalMessageInfo

func (m *Frame) GetRound() int64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *Frame) GetRoots() []*Root {
	if m != nil {
		return m.Roots
	}
	return nil
}

func (m *Frame) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func init() {
	proto.RegisterType((*Frame)(nil), "poset.Frame")
}

func init() { proto.RegisterFile("frame.proto", fileDescriptor_5379e2b825e15002) }

var fileDescriptor_5379e2b825e15002 = []byte{
	// 133 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x2b, 0x4a, 0xcc,
	0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xc8, 0x2f, 0x4e, 0x2d, 0x91, 0xe2,
	0x2a, 0xca, 0xcf, 0x2f, 0x81, 0x08, 0x49, 0x71, 0xa7, 0x96, 0xa5, 0xe6, 0x41, 0x39, 0x4a, 0x69,
	0x5c, 0xac, 0x6e, 0x20, 0xe5, 0x42, 0x22, 0x5c, 0xac, 0x41, 0xf9, 0xa5, 0x79, 0x29, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x90, 0x22, 0x48, 0x34, 0xbf, 0xa4, 0x58, 0x82, 0x49,
	0x81, 0x59, 0x83, 0xdb, 0x88, 0x5b, 0x0f, 0x6c, 0x9c, 0x1e, 0x48, 0x2c, 0x08, 0x22, 0x23, 0xa4,
	0xc2, 0xc5, 0xe6, 0x0a, 0x32, 0xb0, 0x58, 0x82, 0x19, 0xac, 0x86, 0x07, 0xaa, 0x06, 0x2c, 0x18,
	0x04, 0x95, 0x4b, 0x62, 0x03, 0x5b, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x0b, 0xcc,
	0xdb, 0x9d, 0x00, 0x00, 0x00,
}
