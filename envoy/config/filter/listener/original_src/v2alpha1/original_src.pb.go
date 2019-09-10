// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/listener/original_src/v2alpha1/original_src.proto

package envoy_config_filter_listener_original_src_v2alpha1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OriginalSrc struct {
	BindPort             bool     `protobuf:"varint,1,opt,name=bind_port,json=bindPort,proto3" json:"bind_port,omitempty"`
	Mark                 uint32   `protobuf:"varint,2,opt,name=mark,proto3" json:"mark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OriginalSrc) Reset()         { *m = OriginalSrc{} }
func (m *OriginalSrc) String() string { return proto.CompactTextString(m) }
func (*OriginalSrc) ProtoMessage()    {}
func (*OriginalSrc) Descriptor() ([]byte, []int) {
	return fileDescriptor_19e4bd40556a6972, []int{0}
}

func (m *OriginalSrc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OriginalSrc.Unmarshal(m, b)
}
func (m *OriginalSrc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OriginalSrc.Marshal(b, m, deterministic)
}
func (m *OriginalSrc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginalSrc.Merge(m, src)
}
func (m *OriginalSrc) XXX_Size() int {
	return xxx_messageInfo_OriginalSrc.Size(m)
}
func (m *OriginalSrc) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginalSrc.DiscardUnknown(m)
}

var xxx_messageInfo_OriginalSrc proto.InternalMessageInfo

func (m *OriginalSrc) GetBindPort() bool {
	if m != nil {
		return m.BindPort
	}
	return false
}

func (m *OriginalSrc) GetMark() uint32 {
	if m != nil {
		return m.Mark
	}
	return 0
}

func init() {
	proto.RegisterType((*OriginalSrc)(nil), "envoy.config.filter.listener.original_src.v2alpha1.OriginalSrc")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/original_src/v2alpha1/original_src.proto", fileDescriptor_19e4bd40556a6972)
}

var fileDescriptor_19e4bd40556a6972 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0xce, 0x41, 0x4b, 0x87, 0x30,
	0x18, 0xc7, 0x71, 0x16, 0x11, 0xb6, 0x08, 0x62, 0x97, 0xa4, 0x2e, 0xd2, 0xc9, 0xd3, 0x46, 0x76,
	0x8f, 0x10, 0x3a, 0x27, 0x06, 0x5d, 0x65, 0xea, 0xb4, 0x87, 0xd6, 0x9e, 0xf1, 0x38, 0x24, 0xdf,
	0x7d, 0x38, 0x13, 0xf2, 0xf8, 0xbf, 0x3d, 0x3c, 0xf0, 0xfb, 0xf0, 0xe5, 0xaf, 0xc6, 0xcd, 0xb8,
	0xa8, 0x0e, 0xdd, 0x00, 0xa3, 0x1a, 0xc0, 0x06, 0x43, 0xca, 0xc2, 0x14, 0x8c, 0x33, 0xa4, 0x90,
	0x60, 0x04, 0xa7, 0x6d, 0x33, 0x51, 0xa7, 0xe6, 0x42, 0x5b, 0xff, 0xa9, 0x1f, 0x0f, 0x5f, 0xe9,
	0x09, 0x03, 0x8a, 0x22, 0x32, 0x72, 0x63, 0xe4, 0xc6, 0xc8, 0x9d, 0x91, 0x87, 0xc1, 0xce, 0xdc,
	0xdd, 0xce, 0xda, 0x42, 0xaf, 0x83, 0x51, 0xfb, 0xb1, 0x61, 0x0f, 0xcf, 0xfc, 0xea, 0xed, 0x6f,
	0xf1, 0x4e, 0x9d, 0xb8, 0xe7, 0x97, 0x2d, 0xb8, 0xbe, 0xf1, 0x48, 0x21, 0x65, 0x19, 0xcb, 0x93,
	0x3a, 0x59, 0x1f, 0x15, 0x52, 0x10, 0x82, 0x9f, 0x7f, 0x6b, 0xfa, 0x4a, 0xcf, 0x32, 0x96, 0x5f,
	0xd7, 0xf1, 0x2e, 0x3f, 0xf8, 0x0b, 0xa0, 0x8c, 0x45, 0x9e, 0xf0, 0x67, 0x91, 0xa7, 0xc7, 0x95,
	0x37, 0xff, 0x0a, 0xaa, 0xb5, 0xaa, 0x62, 0xed, 0x45, 0xcc, 0x7b, 0xfa, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0xd4, 0x1c, 0x5e, 0x8b, 0x34, 0x01, 0x00, 0x00,
}
