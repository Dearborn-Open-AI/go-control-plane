// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/thrift_proxy/v3alpha/route.proto

package envoy_config_filter_network_thrift_proxy_v3alpha

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/route"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type RouteConfiguration struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Routes               []*Route `protobuf:"bytes,2,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RouteConfiguration) Reset()         { *m = RouteConfiguration{} }
func (m *RouteConfiguration) String() string { return proto.CompactTextString(m) }
func (*RouteConfiguration) ProtoMessage()    {}
func (*RouteConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{0}
}

func (m *RouteConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfiguration.Unmarshal(m, b)
}
func (m *RouteConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfiguration.Marshal(b, m, deterministic)
}
func (m *RouteConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfiguration.Merge(m, src)
}
func (m *RouteConfiguration) XXX_Size() int {
	return xxx_messageInfo_RouteConfiguration.Size(m)
}
func (m *RouteConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfiguration proto.InternalMessageInfo

func (m *RouteConfiguration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfiguration) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type Route struct {
	Match                *RouteMatch  `protobuf:"bytes,1,opt,name=match,proto3" json:"match,omitempty"`
	Route                *RouteAction `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *Route) GetRoute() *RouteAction {
	if m != nil {
		return m.Route
	}
	return nil
}

type RouteMatch struct {
	// Types that are valid to be assigned to MatchSpecifier:
	//	*RouteMatch_MethodName
	//	*RouteMatch_ServiceName
	MatchSpecifier       isRouteMatch_MatchSpecifier `protobuf_oneof:"match_specifier"`
	Invert               bool                        `protobuf:"varint,3,opt,name=invert,proto3" json:"invert,omitempty"`
	Headers              []*route.HeaderMatcher      `protobuf:"bytes,4,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{2}
}

func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (m *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(m, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_MatchSpecifier interface {
	isRouteMatch_MatchSpecifier()
}

type RouteMatch_MethodName struct {
	MethodName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3,oneof"`
}

type RouteMatch_ServiceName struct {
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3,oneof"`
}

func (*RouteMatch_MethodName) isRouteMatch_MatchSpecifier() {}

func (*RouteMatch_ServiceName) isRouteMatch_MatchSpecifier() {}

func (m *RouteMatch) GetMatchSpecifier() isRouteMatch_MatchSpecifier {
	if m != nil {
		return m.MatchSpecifier
	}
	return nil
}

func (m *RouteMatch) GetMethodName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_MethodName); ok {
		return x.MethodName
	}
	return ""
}

func (m *RouteMatch) GetServiceName() string {
	if x, ok := m.GetMatchSpecifier().(*RouteMatch_ServiceName); ok {
		return x.ServiceName
	}
	return ""
}

func (m *RouteMatch) GetInvert() bool {
	if m != nil {
		return m.Invert
	}
	return false
}

func (m *RouteMatch) GetHeaders() []*route.HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteMatch_MethodName)(nil),
		(*RouteMatch_ServiceName)(nil),
	}
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterSpecifier:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedClusters
	ClusterSpecifier     isRouteAction_ClusterSpecifier `protobuf_oneof:"cluster_specifier"`
	MetadataMatch        *core.Metadata                 `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	RateLimits           []*route.RateLimit             `protobuf:"bytes,4,rep,name=rate_limits,json=rateLimits,proto3" json:"rate_limits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{3}
}

func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (m *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(m, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterSpecifier interface {
	isRouteAction_ClusterSpecifier()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3,oneof"`
}

type RouteAction_WeightedClusters struct {
	WeightedClusters *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_clusters,json=weightedClusters,proto3,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterSpecifier() {}

func (*RouteAction_WeightedClusters) isRouteAction_ClusterSpecifier() {}

func (m *RouteAction) GetClusterSpecifier() isRouteAction_ClusterSpecifier {
	if m != nil {
		return m.ClusterSpecifier
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedClusters() *WeightedCluster {
	if x, ok := m.GetClusterSpecifier().(*RouteAction_WeightedClusters); ok {
		return x.WeightedClusters
	}
	return nil
}

func (m *RouteAction) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func (m *RouteAction) GetRateLimits() []*route.RateLimit {
	if m != nil {
		return m.RateLimits
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*RouteAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedClusters)(nil),
	}
}

type WeightedCluster struct {
	Clusters             []*WeightedCluster_ClusterWeight `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{4}
}

func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (m *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(m, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeight {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeight struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=weight,proto3" json:"weight,omitempty"`
	MetadataMatch        *core.Metadata        `protobuf:"bytes,3,opt,name=metadata_match,json=metadataMatch,proto3" json:"metadata_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WeightedCluster_ClusterWeight) Reset()         { *m = WeightedCluster_ClusterWeight{} }
func (m *WeightedCluster_ClusterWeight) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeight) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b3b4ebd82e248c, []int{4, 0}
}

func (m *WeightedCluster_ClusterWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Marshal(b, m, deterministic)
}
func (m *WeightedCluster_ClusterWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeight.Merge(m, src)
}
func (m *WeightedCluster_ClusterWeight) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeight.Size(m)
}
func (m *WeightedCluster_ClusterWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeight.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeight proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeight) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeight) GetWeight() *wrappers.UInt32Value {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *WeightedCluster_ClusterWeight) GetMetadataMatch() *core.Metadata {
	if m != nil {
		return m.MetadataMatch
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfiguration)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteConfiguration")
	proto.RegisterType((*Route)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.Route")
	proto.RegisterType((*RouteMatch)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteMatch")
	proto.RegisterType((*RouteAction)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.RouteAction")
	proto.RegisterType((*WeightedCluster)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeight)(nil), "envoy.config.filter.network.thrift_proxy.v3alpha.WeightedCluster.ClusterWeight")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/thrift_proxy/v3alpha/route.proto", fileDescriptor_82b3b4ebd82e248c)
}

var fileDescriptor_82b3b4ebd82e248c = []byte{
	// 629 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6f, 0x13, 0x3f,
	0x10, 0xad, 0x37, 0x4d, 0x9a, 0xdf, 0xe4, 0x57, 0xda, 0xfa, 0x50, 0xa2, 0x80, 0x50, 0x9a, 0x4a,
	0x28, 0xe2, 0xe0, 0x45, 0xed, 0x81, 0x4b, 0x8b, 0xe8, 0xf6, 0x40, 0x90, 0xe8, 0x1f, 0x59, 0x02,
	0x2e, 0xa0, 0xc8, 0x4d, 0x26, 0x89, 0xc5, 0x66, 0xbd, 0xf2, 0x3a, 0x09, 0xfd, 0x0a, 0x1c, 0xf9,
	0x2c, 0xdc, 0xb9, 0x20, 0x4e, 0x7c, 0x20, 0xd4, 0x13, 0x5a, 0xdb, 0x9b, 0xa6, 0x05, 0x0e, 0x2d,
	0x9c, 0xd6, 0x3b, 0x9e, 0xf7, 0xde, 0xcc, 0x1b, 0xdb, 0xb0, 0x87, 0xc9, 0x54, 0x9d, 0x87, 0x3d,
	0x95, 0x0c, 0xe4, 0x30, 0x1c, 0xc8, 0xd8, 0xa0, 0x0e, 0x13, 0x34, 0x33, 0xa5, 0xdf, 0x87, 0x66,
	0xa4, 0xe5, 0xc0, 0x74, 0x53, 0xad, 0x3e, 0x9c, 0x87, 0xd3, 0x5d, 0x11, 0xa7, 0x23, 0x11, 0x6a,
	0x35, 0x31, 0xc8, 0x52, 0xad, 0x8c, 0xa2, 0x8f, 0x2d, 0x9a, 0x39, 0x34, 0x73, 0x68, 0xe6, 0xd1,
	0x6c, 0x11, 0xcd, 0x3c, 0xba, 0xb1, 0xe5, 0xf4, 0x44, 0x2a, 0xe7, 0x84, 0x3d, 0xa5, 0x31, 0x3c,
	0x13, 0x99, 0x27, 0x6d, 0x6c, 0xff, 0x9a, 0x62, 0x35, 0x17, 0x95, 0x1b, 0x0f, 0x86, 0x4a, 0x0d,
	0x63, 0x0c, 0xed, 0xdf, 0xd9, 0x64, 0x10, 0xce, 0xb4, 0x48, 0x53, 0xd4, 0x99, 0xdf, 0xbf, 0x3b,
	0x15, 0xb1, 0xec, 0x0b, 0x83, 0x61, 0xb1, 0x70, 0x1b, 0xad, 0x73, 0xa0, 0x3c, 0xe7, 0x39, 0xb4,
	0x35, 0x4f, 0xb4, 0x30, 0x52, 0x25, 0x94, 0xc2, 0x72, 0x22, 0xc6, 0x58, 0x27, 0x4d, 0xd2, 0xfe,
	0x8f, 0xdb, 0x35, 0x3d, 0x81, 0x8a, 0x55, 0xcc, 0xea, 0x41, 0xb3, 0xd4, 0xae, 0xed, 0x3c, 0x61,
	0x37, 0xed, 0x96, 0x59, 0x25, 0xee, 0x69, 0x5a, 0x5f, 0x09, 0x94, 0x6d, 0x84, 0xbe, 0x85, 0xf2,
	0x58, 0x98, 0xde, 0xc8, 0xea, 0xd5, 0x76, 0xf6, 0x6e, 0xc9, 0x7c, 0x94, 0x73, 0x44, 0xd5, 0x8b,
	0xa8, 0xfc, 0x91, 0x04, 0xeb, 0x84, 0x3b, 0x52, 0xfa, 0x0e, 0xca, 0x56, 0xb1, 0x1e, 0x58, 0xf6,
	0xfd, 0x5b, 0xb2, 0x1f, 0xf4, 0x72, 0x6b, 0x16, 0xe9, 0x2d, 0x6b, 0xeb, 0x1b, 0x01, 0xb8, 0x94,
	0xa7, 0x5b, 0x50, 0x1b, 0xa3, 0x19, 0xa9, 0x7e, 0xf7, 0xd2, 0xc1, 0xce, 0x12, 0x07, 0x17, 0x3c,
	0xce, 0x9d, 0xdc, 0x86, 0xff, 0x33, 0xd4, 0x53, 0xd9, 0x43, 0x97, 0x13, 0xf8, 0x9c, 0x9a, 0x8f,
	0xda, 0xa4, 0x4d, 0xa8, 0xc8, 0x64, 0x8a, 0xda, 0xd4, 0x4b, 0x4d, 0xd2, 0xae, 0x72, 0xff, 0x47,
	0x9f, 0xc1, 0xca, 0x08, 0x45, 0x1f, 0x75, 0x56, 0x5f, 0xb6, 0x73, 0x78, 0xe8, 0xfb, 0x11, 0xa9,
	0x9c, 0x17, 0xec, 0x8e, 0x46, 0xc7, 0xe6, 0xd9, 0xb2, 0x50, 0xf3, 0x02, 0x16, 0x6d, 0xc2, 0x9a,
	0x35, 0xa6, 0x9b, 0xa5, 0xd8, 0x93, 0x03, 0x89, 0x9a, 0x96, 0x7e, 0x44, 0xa4, 0xf5, 0x3d, 0x80,
	0xda, 0x42, 0xa7, 0x74, 0x1b, 0x56, 0x7a, 0xf1, 0x24, 0x33, 0xa8, 0x5d, 0x17, 0xd1, 0xca, 0x45,
	0xb4, 0xac, 0x83, 0x26, 0xe9, 0x2c, 0xf1, 0x62, 0x87, 0xa6, 0xb0, 0x31, 0x43, 0x39, 0x1c, 0x19,
	0xec, 0x77, 0x7d, 0x2c, 0xf3, 0x46, 0x1f, 0xdc, 0xdc, 0xe8, 0x37, 0x9e, 0xea, 0xd0, 0x31, 0x75,
	0x96, 0xf8, 0xfa, 0xec, 0x6a, 0x28, 0xa3, 0xcf, 0xe1, 0xce, 0x18, 0x8d, 0xe8, 0x0b, 0x23, 0xba,
	0xee, 0xd4, 0x94, 0xac, 0x5c, 0xf3, 0x37, 0x3e, 0xe4, 0x77, 0x89, 0x1d, 0xf9, 0x6c, 0xbe, 0x5a,
	0xe0, 0xdc, 0xa4, 0x0e, 0xa1, 0xa6, 0x85, 0xc1, 0x6e, 0x2c, 0xc7, 0xd2, 0x14, 0x6e, 0xb6, 0xfe,
	0xe8, 0x26, 0x17, 0x06, 0x5f, 0xe6, 0xa9, 0x1c, 0x74, 0xb1, 0xcc, 0xa2, 0x3a, 0x6c, 0xf8, 0xb6,
	0xaf, 0xdb, 0xf9, 0x25, 0x80, 0xb5, 0x6b, 0xfd, 0xd0, 0x09, 0x54, 0xe7, 0x26, 0x11, 0xab, 0x77,
	0xf2, 0xd7, 0x26, 0x31, 0xff, 0x75, 0x61, 0x7b, 0x3e, 0x3f, 0x91, 0xa0, 0x4a, 0xf8, 0x5c, 0xaa,
	0xf1, 0x99, 0xc0, 0xea, 0x95, 0x2c, 0x7a, 0x6f, 0xf1, 0x82, 0xcf, 0x07, 0xeb, 0x6f, 0xfa, 0x3e,
	0x54, 0x9c, 0xeb, 0x7e, 0x90, 0xf7, 0x99, 0x7b, 0x5d, 0x58, 0xf1, 0xba, 0xb0, 0x57, 0x2f, 0x12,
	0xb3, 0xbb, 0xf3, 0x5a, 0xc4, 0x13, 0xb4, 0xe0, 0x47, 0x41, 0x9b, 0x70, 0x0f, 0xfa, 0x67, 0x03,
	0x8a, 0x8e, 0xe1, 0xa9, 0x54, 0x0e, 0xe4, 0x1c, 0xb8, 0xa9, 0x55, 0x91, 0xbb, 0x98, 0xa7, 0x79,
	0xd9, 0xa7, 0xe4, 0xac, 0x62, 0xeb, 0xdf, 0xfd, 0x19, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xa4, 0x02,
	0x85, 0xe5, 0x05, 0x00, 0x00,
}
