// Code generated by protoc-gen-go. DO NOT EDIT.
// source: object.proto

/*
Package storage is a generated protocol buffer package.

It is generated from these files:
	object.proto

It has these top-level messages:
	Object
*/
package storage

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

// Object is the storage object for a draft applications build history.
type Object struct {
	BuildID     string `protobuf:"bytes,1,opt,name=buildID" json:"buildID,omitempty"`
	Release     string `protobuf:"bytes,2,opt,name=release" json:"release,omitempty"`
	ContextID   []byte `protobuf:"bytes,3,opt,name=contextID,proto3" json:"contextID,omitempty"`
	LogsFileRef string `protobuf:"bytes,4,opt,name=logs_file_ref,json=logsFileRef" json:"logs_file_ref,omitempty"`
}

func (m *Object) Reset()                    { *m = Object{} }
func (m *Object) String() string            { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()               {}
func (*Object) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Object) GetBuildID() string {
	if m != nil {
		return m.BuildID
	}
	return ""
}

func (m *Object) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Object) GetContextID() []byte {
	if m != nil {
		return m.ContextID
	}
	return nil
}

func (m *Object) GetLogsFileRef() string {
	if m != nil {
		return m.LogsFileRef
	}
	return ""
}

func init() {
	proto.RegisterType((*Object)(nil), "storage.Object")
}

func init() { proto.RegisterFile("object.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4f, 0xca, 0x4a,
	0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c,
	0x4f, 0x55, 0xaa, 0xe3, 0x62, 0xf3, 0x07, 0x4b, 0x08, 0x49, 0x70, 0xb1, 0x27, 0x95, 0x66, 0xe6,
	0xa4, 0x78, 0xba, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x20, 0x99, 0xa2, 0xd4,
	0x9c, 0xd4, 0xc4, 0xe2, 0x54, 0x09, 0x26, 0x88, 0x0c, 0x94, 0x2b, 0x24, 0xc3, 0xc5, 0x99, 0x9c,
	0x9f, 0x57, 0x92, 0x5a, 0x51, 0xe2, 0xe9, 0x22, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x13, 0x84, 0x10,
	0x10, 0x52, 0xe2, 0xe2, 0xcd, 0xc9, 0x4f, 0x2f, 0x8e, 0x4f, 0xcb, 0xcc, 0x49, 0x8d, 0x2f, 0x4a,
	0x4d, 0x93, 0x60, 0x01, 0xeb, 0xe6, 0x06, 0x09, 0xba, 0x65, 0xe6, 0xa4, 0x06, 0xa5, 0xa6, 0x25,
	0xb1, 0x81, 0xdd, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x73, 0xd8, 0xd6, 0x04, 0x9f, 0x00,
	0x00, 0x00,
}