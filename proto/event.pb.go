// Code generated by protoc-gen-go. DO NOT EDIT.
// source: event.proto

/*
Package event is a generated protocol buffer package.

It is generated from these files:
	event.proto

It has these top-level messages:
	Transport
	DummyMessage
*/
package event

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

type Dummy int32

const (
	Dummy_Subject1     Dummy = 0
	Dummy_QueueName1   Dummy = 1
	Dummy_DurableName1 Dummy = 2
	Dummy_Subject2     Dummy = 3
	Dummy_QueueName2   Dummy = 4
	Dummy_DurableName2 Dummy = 5
)

var Dummy_name = map[int32]string{
	0: "Subject1",
	1: "QueueName1",
	2: "DurableName1",
	3: "Subject2",
	4: "QueueName2",
	5: "DurableName2",
}
var Dummy_value = map[string]int32{
	"Subject1":     0,
	"QueueName1":   1,
	"DurableName1": 2,
	"Subject2":     3,
	"QueueName2":   4,
	"DurableName2": 5,
}

func (x Dummy) String() string {
	return proto.EnumName(Dummy_name, int32(x))
}
func (Dummy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Transport struct {
	CoorlationId string `protobuf:"bytes,1,opt,name=coorlationId" json:"coorlationId,omitempty"`
	Signature    string `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Payload      []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Transport) Reset()                    { *m = Transport{} }
func (m *Transport) String() string            { return proto.CompactTextString(m) }
func (*Transport) ProtoMessage()               {}
func (*Transport) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Transport) GetCoorlationId() string {
	if m != nil {
		return m.CoorlationId
	}
	return ""
}

func (m *Transport) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Transport) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type DummyMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *DummyMessage) Reset()                    { *m = DummyMessage{} }
func (m *DummyMessage) String() string            { return proto.CompactTextString(m) }
func (*DummyMessage) ProtoMessage()               {}
func (*DummyMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DummyMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Transport)(nil), "event.Transport")
	proto.RegisterType((*DummyMessage)(nil), "event.DummyMessage")
	proto.RegisterEnum("event.Dummy", Dummy_name, Dummy_value)
}

func init() { proto.RegisterFile("event.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x4d, 0x6b, 0xd4, 0x1d, 0x83, 0x84, 0xc1, 0x43, 0x0e, 0x1e, 0xca, 0xe2, 0xa1, 0x78,
	0x10, 0xba, 0xbe, 0x82, 0x17, 0x0f, 0x0a, 0x56, 0x5f, 0x60, 0xb6, 0x1d, 0x42, 0x25, 0x9b, 0x2c,
	0xf9, 0x53, 0xe8, 0xdb, 0x8b, 0xb1, 0xa2, 0xeb, 0xf1, 0xfb, 0xf8, 0x98, 0x81, 0x1f, 0x5c, 0xf2,
	0x9e, 0x7d, 0xbe, 0x1f, 0x63, 0xc8, 0x01, 0x65, 0x85, 0xd6, 0x42, 0xf3, 0x1e, 0xc9, 0xa7, 0x31,
	0xc4, 0x8c, 0x2d, 0xa8, 0x4d, 0x08, 0xd1, 0x51, 0xde, 0x05, 0xff, 0xb4, 0x35, 0x62, 0x21, 0x96,
	0xcd, 0x7a, 0xe2, 0xf0, 0x06, 0x9a, 0xb4, 0xb3, 0x9e, 0x72, 0x89, 0x6c, 0x66, 0x35, 0xf8, 0x15,
	0x68, 0xe0, 0x7c, 0xa4, 0x83, 0x0b, 0xb4, 0x35, 0xf3, 0x85, 0x58, 0xaa, 0xf5, 0x0f, 0xb6, 0xb7,
	0xa0, 0x1e, 0xcb, 0x30, 0x1c, 0x9e, 0x39, 0x25, 0xb2, 0x8c, 0xd7, 0x20, 0xf7, 0xe4, 0x0a, 0x1f,
	0x9f, 0x7c, 0xc3, 0x9d, 0x05, 0x59, 0x2b, 0x54, 0x70, 0xf1, 0x56, 0xfa, 0x0f, 0xde, 0xe4, 0x95,
	0x3e, 0xc1, 0x2b, 0x80, 0xd7, 0xc2, 0x85, 0x5f, 0x68, 0xe0, 0x95, 0x16, 0xa8, 0xbf, 0x8e, 0x45,
	0xea, 0xdd, 0xd1, 0xcc, 0xfe, 0xf4, 0x9d, 0x9e, 0x4f, 0xfa, 0x4e, 0x9f, 0xfe, 0xeb, 0x3b, 0x2d,
	0xfb, 0xb3, 0xba, 0xc2, 0xc3, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xb3, 0xda, 0xe8, 0x14,
	0x01, 0x00, 0x00,
}
