// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/health_check/v2/health_check.proto

package envoy_config_filter_http_health_check_v2

import (
	fmt "fmt"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	_type "github.com/envoyproxy/go-control-plane/envoy/type"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type HealthCheck struct {
	PassThroughMode              *wrappers.BoolValue       `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	CacheTime                    *duration.Duration        `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	ClusterMinHealthyPercentages map[string]*_type.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Headers                      []*route.HeaderMatcher    `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                  `json:"-"`
	XXX_unrecognized             []byte                    `json:"-"`
	XXX_sizecache                int32                     `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_75439d7b4d98e201, []int{0}
}

func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPassThroughMode() *wrappers.BoolValue {
	if m != nil {
		return m.PassThroughMode
	}
	return nil
}

func (m *HealthCheck) GetCacheTime() *duration.Duration {
	if m != nil {
		return m.CacheTime
	}
	return nil
}

func (m *HealthCheck) GetClusterMinHealthyPercentages() map[string]*_type.Percent {
	if m != nil {
		return m.ClusterMinHealthyPercentages
	}
	return nil
}

func (m *HealthCheck) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck")
	proto.RegisterMapType((map[string]*_type.Percent)(nil), "envoy.config.filter.http.health_check.v2.HealthCheck.ClusterMinHealthyPercentagesEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/health_check/v2/health_check.proto", fileDescriptor_75439d7b4d98e201)
}

var fileDescriptor_75439d7b4d98e201 = []byte{
	// 446 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x55, 0x9a, 0x15, 0x3a, 0xf7, 0x40, 0x09, 0x07, 0x42, 0x85, 0x60, 0xe3, 0x54, 0x2e, 0xb6,
	0x54, 0x24, 0x34, 0xb1, 0x5b, 0x06, 0xd2, 0x84, 0xa8, 0x54, 0x45, 0x13, 0x1c, 0x23, 0xcf, 0xf9,
	0x1a, 0x5b, 0x4b, 0x6d, 0xcb, 0x71, 0x02, 0xf9, 0x0b, 0xfc, 0x08, 0xfe, 0x21, 0x7f, 0x80, 0x13,
	0xb2, 0x9d, 0x88, 0x4d, 0x93, 0x58, 0x2f, 0x91, 0xe3, 0xf7, 0xbd, 0xf7, 0x3d, 0xbf, 0x87, 0xce,
	0x41, 0x76, 0xaa, 0x27, 0x4c, 0xc9, 0x9d, 0xa8, 0xc8, 0x4e, 0xd4, 0x16, 0x0c, 0xe1, 0xd6, 0x6a,
	0xc2, 0x81, 0xd6, 0x96, 0x17, 0x8c, 0x03, 0xbb, 0x21, 0xdd, 0xfa, 0xce, 0x3f, 0xd6, 0x46, 0x59,
	0x95, 0xac, 0x3c, 0x19, 0x07, 0x32, 0x0e, 0x64, 0xec, 0xc8, 0xf8, 0xce, 0x70, 0xb7, 0x5e, 0xbe,
	0xaa, 0x94, 0xaa, 0x6a, 0x20, 0x9e, 0x77, 0xdd, 0xee, 0x48, 0xd9, 0x1a, 0x6a, 0x85, 0x92, 0x41,
	0xe9, 0x3e, 0xfe, 0xdd, 0x50, 0xad, 0xc1, 0x34, 0x23, 0x1e, 0x6c, 0x52, 0x2d, 0x9c, 0x15, 0xa3,
	0x5a, 0x0b, 0xe1, 0x3b, 0xe0, 0x69, 0xc0, 0x6d, 0xaf, 0x81, 0x68, 0x30, 0x0c, 0xa4, 0x1d, 0x90,
	0xe7, 0x1d, 0xad, 0x45, 0x49, 0x2d, 0x90, 0xf1, 0x10, 0x80, 0x37, 0xbf, 0x63, 0x34, 0xbf, 0xf4,
	0x36, 0x2f, 0x9c, 0xcb, 0x64, 0x8b, 0x9e, 0x6a, 0xda, 0x34, 0x85, 0xe5, 0x46, 0xb5, 0x15, 0x2f,
	0xf6, 0xaa, 0x84, 0x34, 0x3a, 0x89, 0x56, 0xf3, 0xf5, 0x12, 0x07, 0x7b, 0x78, 0xb4, 0x87, 0x33,
	0xa5, 0xea, 0xaf, 0xb4, 0x6e, 0x21, 0x9b, 0xfd, 0xc9, 0xa6, 0x3f, 0xa3, 0xc9, 0x22, 0xca, 0x9f,
	0x38, 0xfa, 0x55, 0x60, 0x6f, 0x54, 0x09, 0xc9, 0x19, 0x42, 0x8c, 0x32, 0x0e, 0x85, 0x15, 0x7b,
	0x48, 0x63, 0x2f, 0xf5, 0xe2, 0x9e, 0xd4, 0xc7, 0x21, 0x89, 0xfc, 0xd8, 0x0f, 0x5f, 0x89, 0x3d,
	0x24, 0xbf, 0x22, 0xf4, 0x9a, 0xd5, 0x6d, 0x63, 0xc1, 0x14, 0x7b, 0x21, 0x8b, 0x10, 0x67, 0x5f,
	0x0c, 0x4f, 0xa3, 0x15, 0x34, 0xe9, 0xd1, 0x49, 0xbc, 0x9a, 0xaf, 0xbf, 0xe1, 0x43, 0x3b, 0xc0,
	0xb7, 0x1e, 0x8b, 0x2f, 0x82, 0xf8, 0x46, 0xc8, 0x70, 0xdb, 0x6f, 0xff, 0x29, 0x7f, 0x92, 0xd6,
	0xf4, 0xf9, 0x4b, 0xf6, 0x9f, 0x91, 0xe4, 0x1c, 0x3d, 0xe6, 0x40, 0x4b, 0x30, 0x4d, 0x3a, 0xf5,
	0x3e, 0x4e, 0x07, 0x1f, 0x54, 0x0b, 0xb7, 0x2b, 0x74, 0x73, 0xe9, 0x47, 0x36, 0xd4, 0x32, 0x0e,
	0x26, 0x1f, 0x19, 0xcb, 0x12, 0x9d, 0x3e, 0xb8, 0x3f, 0x59, 0xa0, 0xf8, 0x06, 0x7a, 0x5f, 0xc0,
	0x71, 0xee, 0x8e, 0xc9, 0x5b, 0x34, 0xed, 0x5c, 0xe4, 0xe9, 0xc4, 0x27, 0xf9, 0x6c, 0xd8, 0xe8,
	0x3a, 0xc7, 0x03, 0x3d, 0x0f, 0x13, 0x1f, 0x26, 0x67, 0xd1, 0xe7, 0xa3, 0xd9, 0x64, 0x11, 0xe7,
	0x33, 0x90, 0xa5, 0x56, 0x42, 0xda, 0xec, 0x0b, 0x7a, 0x2f, 0x54, 0xe0, 0x68, 0xa3, 0x7e, 0xf4,
	0x07, 0x07, 0x97, 0x2d, 0x6e, 0x25, 0xb7, 0x75, 0xb5, 0x6d, 0xa3, 0xeb, 0x47, 0xbe, 0xbf, 0x77,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xad, 0xeb, 0x3b, 0x40, 0x03, 0x00, 0x00,
}
