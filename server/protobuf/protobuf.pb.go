// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	protobuf.proto

It has these top-level messages:
	TheMsg
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TheMsg struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Num  int32  `protobuf:"varint,2,opt,name=num" json:"num,omitempty"`
}

func (m *TheMsg) Reset()                    { *m = TheMsg{} }
func (m *TheMsg) String() string            { return proto.CompactTextString(m) }
func (*TheMsg) ProtoMessage()               {}
func (*TheMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TheMsg) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TheMsg) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func init() {
	proto.RegisterType((*TheMsg)(nil), "TheMsg")
}

func init() { proto.RegisterFile("protobuf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 84 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x03, 0x33, 0x94, 0xf4, 0xb8, 0xd8, 0x42, 0x32, 0x52, 0x7d, 0x8b,
	0xd3, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0xbc, 0xd2, 0x5c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xd6, 0x20,
	0x10, 0x33, 0x89, 0x0d, 0xac, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x92, 0x69, 0xce, 0x54,
	0x48, 0x00, 0x00, 0x00,
}