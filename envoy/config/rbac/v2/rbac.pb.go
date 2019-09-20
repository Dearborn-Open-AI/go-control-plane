// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/rbac/v2/rbac.proto

package envoy_config_rbac_v2

import (
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	route "github.com/envoyproxy/go-control-plane/envoy/api/v2/route"
	matcher "github.com/envoyproxy/go-control-plane/envoy/type/matcher"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1alpha1 "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
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

type RBAC_Action int32

const (
	RBAC_ALLOW RBAC_Action = 0
	RBAC_DENY  RBAC_Action = 1
)

var RBAC_Action_name = map[int32]string{
	0: "ALLOW",
	1: "DENY",
}

var RBAC_Action_value = map[string]int32{
	"ALLOW": 0,
	"DENY":  1,
}

func (x RBAC_Action) String() string {
	return proto.EnumName(RBAC_Action_name, int32(x))
}

func (RBAC_Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{0, 0}
}

type RBAC struct {
	Action               RBAC_Action        `protobuf:"varint,1,opt,name=action,proto3,enum=envoy.config.rbac.v2.RBAC_Action" json:"action,omitempty"`
	Policies             map[string]*Policy `protobuf:"bytes,2,rep,name=policies,proto3" json:"policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RBAC) Reset()         { *m = RBAC{} }
func (m *RBAC) String() string { return proto.CompactTextString(m) }
func (*RBAC) ProtoMessage()    {}
func (*RBAC) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{0}
}

func (m *RBAC) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RBAC.Unmarshal(m, b)
}
func (m *RBAC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RBAC.Marshal(b, m, deterministic)
}
func (m *RBAC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RBAC.Merge(m, src)
}
func (m *RBAC) XXX_Size() int {
	return xxx_messageInfo_RBAC.Size(m)
}
func (m *RBAC) XXX_DiscardUnknown() {
	xxx_messageInfo_RBAC.DiscardUnknown(m)
}

var xxx_messageInfo_RBAC proto.InternalMessageInfo

func (m *RBAC) GetAction() RBAC_Action {
	if m != nil {
		return m.Action
	}
	return RBAC_ALLOW
}

func (m *RBAC) GetPolicies() map[string]*Policy {
	if m != nil {
		return m.Policies
	}
	return nil
}

type Policy struct {
	Permissions          []*Permission  `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	Principals           []*Principal   `protobuf:"bytes,2,rep,name=principals,proto3" json:"principals,omitempty"`
	Condition            *v1alpha1.Expr `protobuf:"bytes,3,opt,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Policy) Reset()         { *m = Policy{} }
func (m *Policy) String() string { return proto.CompactTextString(m) }
func (*Policy) ProtoMessage()    {}
func (*Policy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{1}
}

func (m *Policy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Policy.Unmarshal(m, b)
}
func (m *Policy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Policy.Marshal(b, m, deterministic)
}
func (m *Policy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Policy.Merge(m, src)
}
func (m *Policy) XXX_Size() int {
	return xxx_messageInfo_Policy.Size(m)
}
func (m *Policy) XXX_DiscardUnknown() {
	xxx_messageInfo_Policy.DiscardUnknown(m)
}

var xxx_messageInfo_Policy proto.InternalMessageInfo

func (m *Policy) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *Policy) GetPrincipals() []*Principal {
	if m != nil {
		return m.Principals
	}
	return nil
}

func (m *Policy) GetCondition() *v1alpha1.Expr {
	if m != nil {
		return m.Condition
	}
	return nil
}

type Permission struct {
	// Types that are valid to be assigned to Rule:
	//	*Permission_AndRules
	//	*Permission_OrRules
	//	*Permission_Any
	//	*Permission_Header
	//	*Permission_DestinationIp
	//	*Permission_DestinationPort
	//	*Permission_Metadata
	//	*Permission_NotRule
	//	*Permission_RequestedServerName
	Rule                 isPermission_Rule `protobuf_oneof:"rule"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{2}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

type isPermission_Rule interface {
	isPermission_Rule()
}

type Permission_AndRules struct {
	AndRules *Permission_Set `protobuf:"bytes,1,opt,name=and_rules,json=andRules,proto3,oneof"`
}

type Permission_OrRules struct {
	OrRules *Permission_Set `protobuf:"bytes,2,opt,name=or_rules,json=orRules,proto3,oneof"`
}

type Permission_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Permission_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,4,opt,name=header,proto3,oneof"`
}

type Permission_DestinationIp struct {
	DestinationIp *core.CidrRange `protobuf:"bytes,5,opt,name=destination_ip,json=destinationIp,proto3,oneof"`
}

type Permission_DestinationPort struct {
	DestinationPort uint32 `protobuf:"varint,6,opt,name=destination_port,json=destinationPort,proto3,oneof"`
}

type Permission_Metadata struct {
	Metadata *matcher.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Permission_NotRule struct {
	NotRule *Permission `protobuf:"bytes,8,opt,name=not_rule,json=notRule,proto3,oneof"`
}

type Permission_RequestedServerName struct {
	RequestedServerName *matcher.StringMatcher `protobuf:"bytes,9,opt,name=requested_server_name,json=requestedServerName,proto3,oneof"`
}

func (*Permission_AndRules) isPermission_Rule() {}

func (*Permission_OrRules) isPermission_Rule() {}

func (*Permission_Any) isPermission_Rule() {}

func (*Permission_Header) isPermission_Rule() {}

func (*Permission_DestinationIp) isPermission_Rule() {}

func (*Permission_DestinationPort) isPermission_Rule() {}

func (*Permission_Metadata) isPermission_Rule() {}

func (*Permission_NotRule) isPermission_Rule() {}

func (*Permission_RequestedServerName) isPermission_Rule() {}

func (m *Permission) GetRule() isPermission_Rule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func (m *Permission) GetAndRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_AndRules); ok {
		return x.AndRules
	}
	return nil
}

func (m *Permission) GetOrRules() *Permission_Set {
	if x, ok := m.GetRule().(*Permission_OrRules); ok {
		return x.OrRules
	}
	return nil
}

func (m *Permission) GetAny() bool {
	if x, ok := m.GetRule().(*Permission_Any); ok {
		return x.Any
	}
	return false
}

func (m *Permission) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetRule().(*Permission_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Permission) GetDestinationIp() *core.CidrRange {
	if x, ok := m.GetRule().(*Permission_DestinationIp); ok {
		return x.DestinationIp
	}
	return nil
}

func (m *Permission) GetDestinationPort() uint32 {
	if x, ok := m.GetRule().(*Permission_DestinationPort); ok {
		return x.DestinationPort
	}
	return 0
}

func (m *Permission) GetMetadata() *matcher.MetadataMatcher {
	if x, ok := m.GetRule().(*Permission_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Permission) GetNotRule() *Permission {
	if x, ok := m.GetRule().(*Permission_NotRule); ok {
		return x.NotRule
	}
	return nil
}

func (m *Permission) GetRequestedServerName() *matcher.StringMatcher {
	if x, ok := m.GetRule().(*Permission_RequestedServerName); ok {
		return x.RequestedServerName
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Permission) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Permission_AndRules)(nil),
		(*Permission_OrRules)(nil),
		(*Permission_Any)(nil),
		(*Permission_Header)(nil),
		(*Permission_DestinationIp)(nil),
		(*Permission_DestinationPort)(nil),
		(*Permission_Metadata)(nil),
		(*Permission_NotRule)(nil),
		(*Permission_RequestedServerName)(nil),
	}
}

type Permission_Set struct {
	Rules                []*Permission `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Permission_Set) Reset()         { *m = Permission_Set{} }
func (m *Permission_Set) String() string { return proto.CompactTextString(m) }
func (*Permission_Set) ProtoMessage()    {}
func (*Permission_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{2, 0}
}

func (m *Permission_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission_Set.Unmarshal(m, b)
}
func (m *Permission_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission_Set.Marshal(b, m, deterministic)
}
func (m *Permission_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission_Set.Merge(m, src)
}
func (m *Permission_Set) XXX_Size() int {
	return xxx_messageInfo_Permission_Set.Size(m)
}
func (m *Permission_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Permission_Set proto.InternalMessageInfo

func (m *Permission_Set) GetRules() []*Permission {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Principal struct {
	// Types that are valid to be assigned to Identifier:
	//	*Principal_AndIds
	//	*Principal_OrIds
	//	*Principal_Any
	//	*Principal_Authenticated_
	//	*Principal_SourceIp
	//	*Principal_Header
	//	*Principal_Metadata
	//	*Principal_NotId
	Identifier           isPrincipal_Identifier `protobuf_oneof:"identifier"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal) Reset()         { *m = Principal{} }
func (m *Principal) String() string { return proto.CompactTextString(m) }
func (*Principal) ProtoMessage()    {}
func (*Principal) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{3}
}

func (m *Principal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal.Unmarshal(m, b)
}
func (m *Principal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal.Marshal(b, m, deterministic)
}
func (m *Principal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal.Merge(m, src)
}
func (m *Principal) XXX_Size() int {
	return xxx_messageInfo_Principal.Size(m)
}
func (m *Principal) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal.DiscardUnknown(m)
}

var xxx_messageInfo_Principal proto.InternalMessageInfo

type isPrincipal_Identifier interface {
	isPrincipal_Identifier()
}

type Principal_AndIds struct {
	AndIds *Principal_Set `protobuf:"bytes,1,opt,name=and_ids,json=andIds,proto3,oneof"`
}

type Principal_OrIds struct {
	OrIds *Principal_Set `protobuf:"bytes,2,opt,name=or_ids,json=orIds,proto3,oneof"`
}

type Principal_Any struct {
	Any bool `protobuf:"varint,3,opt,name=any,proto3,oneof"`
}

type Principal_Authenticated_ struct {
	Authenticated *Principal_Authenticated `protobuf:"bytes,4,opt,name=authenticated,proto3,oneof"`
}

type Principal_SourceIp struct {
	SourceIp *core.CidrRange `protobuf:"bytes,5,opt,name=source_ip,json=sourceIp,proto3,oneof"`
}

type Principal_Header struct {
	Header *route.HeaderMatcher `protobuf:"bytes,6,opt,name=header,proto3,oneof"`
}

type Principal_Metadata struct {
	Metadata *matcher.MetadataMatcher `protobuf:"bytes,7,opt,name=metadata,proto3,oneof"`
}

type Principal_NotId struct {
	NotId *Principal `protobuf:"bytes,8,opt,name=not_id,json=notId,proto3,oneof"`
}

func (*Principal_AndIds) isPrincipal_Identifier() {}

func (*Principal_OrIds) isPrincipal_Identifier() {}

func (*Principal_Any) isPrincipal_Identifier() {}

func (*Principal_Authenticated_) isPrincipal_Identifier() {}

func (*Principal_SourceIp) isPrincipal_Identifier() {}

func (*Principal_Header) isPrincipal_Identifier() {}

func (*Principal_Metadata) isPrincipal_Identifier() {}

func (*Principal_NotId) isPrincipal_Identifier() {}

func (m *Principal) GetIdentifier() isPrincipal_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *Principal) GetAndIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_AndIds); ok {
		return x.AndIds
	}
	return nil
}

func (m *Principal) GetOrIds() *Principal_Set {
	if x, ok := m.GetIdentifier().(*Principal_OrIds); ok {
		return x.OrIds
	}
	return nil
}

func (m *Principal) GetAny() bool {
	if x, ok := m.GetIdentifier().(*Principal_Any); ok {
		return x.Any
	}
	return false
}

func (m *Principal) GetAuthenticated() *Principal_Authenticated {
	if x, ok := m.GetIdentifier().(*Principal_Authenticated_); ok {
		return x.Authenticated
	}
	return nil
}

func (m *Principal) GetSourceIp() *core.CidrRange {
	if x, ok := m.GetIdentifier().(*Principal_SourceIp); ok {
		return x.SourceIp
	}
	return nil
}

func (m *Principal) GetHeader() *route.HeaderMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Header); ok {
		return x.Header
	}
	return nil
}

func (m *Principal) GetMetadata() *matcher.MetadataMatcher {
	if x, ok := m.GetIdentifier().(*Principal_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (m *Principal) GetNotId() *Principal {
	if x, ok := m.GetIdentifier().(*Principal_NotId); ok {
		return x.NotId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Principal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Principal_AndIds)(nil),
		(*Principal_OrIds)(nil),
		(*Principal_Any)(nil),
		(*Principal_Authenticated_)(nil),
		(*Principal_SourceIp)(nil),
		(*Principal_Header)(nil),
		(*Principal_Metadata)(nil),
		(*Principal_NotId)(nil),
	}
}

type Principal_Set struct {
	Ids                  []*Principal `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Principal_Set) Reset()         { *m = Principal_Set{} }
func (m *Principal_Set) String() string { return proto.CompactTextString(m) }
func (*Principal_Set) ProtoMessage()    {}
func (*Principal_Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{3, 0}
}

func (m *Principal_Set) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Set.Unmarshal(m, b)
}
func (m *Principal_Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Set.Marshal(b, m, deterministic)
}
func (m *Principal_Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Set.Merge(m, src)
}
func (m *Principal_Set) XXX_Size() int {
	return xxx_messageInfo_Principal_Set.Size(m)
}
func (m *Principal_Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Set proto.InternalMessageInfo

func (m *Principal_Set) GetIds() []*Principal {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Principal_Authenticated struct {
	PrincipalName        *matcher.StringMatcher `protobuf:"bytes,2,opt,name=principal_name,json=principalName,proto3" json:"principal_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Principal_Authenticated) Reset()         { *m = Principal_Authenticated{} }
func (m *Principal_Authenticated) String() string { return proto.CompactTextString(m) }
func (*Principal_Authenticated) ProtoMessage()    {}
func (*Principal_Authenticated) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8a2b527e1e731e1, []int{3, 1}
}

func (m *Principal_Authenticated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Principal_Authenticated.Unmarshal(m, b)
}
func (m *Principal_Authenticated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Principal_Authenticated.Marshal(b, m, deterministic)
}
func (m *Principal_Authenticated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Principal_Authenticated.Merge(m, src)
}
func (m *Principal_Authenticated) XXX_Size() int {
	return xxx_messageInfo_Principal_Authenticated.Size(m)
}
func (m *Principal_Authenticated) XXX_DiscardUnknown() {
	xxx_messageInfo_Principal_Authenticated.DiscardUnknown(m)
}

var xxx_messageInfo_Principal_Authenticated proto.InternalMessageInfo

func (m *Principal_Authenticated) GetPrincipalName() *matcher.StringMatcher {
	if m != nil {
		return m.PrincipalName
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.rbac.v2.RBAC_Action", RBAC_Action_name, RBAC_Action_value)
	proto.RegisterType((*RBAC)(nil), "envoy.config.rbac.v2.RBAC")
	proto.RegisterMapType((map[string]*Policy)(nil), "envoy.config.rbac.v2.RBAC.PoliciesEntry")
	proto.RegisterType((*Policy)(nil), "envoy.config.rbac.v2.Policy")
	proto.RegisterType((*Permission)(nil), "envoy.config.rbac.v2.Permission")
	proto.RegisterType((*Permission_Set)(nil), "envoy.config.rbac.v2.Permission.Set")
	proto.RegisterType((*Principal)(nil), "envoy.config.rbac.v2.Principal")
	proto.RegisterType((*Principal_Set)(nil), "envoy.config.rbac.v2.Principal.Set")
	proto.RegisterType((*Principal_Authenticated)(nil), "envoy.config.rbac.v2.Principal.Authenticated")
}

func init() { proto.RegisterFile("envoy/config/rbac/v2/rbac.proto", fileDescriptor_e8a2b527e1e731e1) }

var fileDescriptor_e8a2b527e1e731e1 = []byte{
	// 887 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x18, 0x8d, 0x9d, 0xc4, 0xb5, 0xbf, 0x28, 0x4b, 0x18, 0x40, 0x58, 0x01, 0x76, 0xb3, 0x59, 0x90,
	0x22, 0x24, 0x6c, 0x35, 0x48, 0x68, 0xa1, 0x80, 0x88, 0xbb, 0x15, 0x09, 0xea, 0x2e, 0x95, 0x2b,
	0xb4, 0xda, 0xab, 0x6a, 0xea, 0x99, 0x4d, 0x07, 0x92, 0x19, 0x33, 0x9e, 0x44, 0xcd, 0x25, 0xaf,
	0xc0, 0x03, 0xf0, 0x10, 0xbc, 0x10, 0xaf, 0x81, 0x7a, 0xb3, 0x68, 0x66, 0x9c, 0x34, 0x96, 0xba,
	0x9b, 0x8d, 0xc4, 0x4d, 0x6b, 0xc5, 0xe7, 0x9c, 0xef, 0x67, 0xce, 0x19, 0xc3, 0x03, 0xca, 0x97,
	0x62, 0x15, 0x67, 0x82, 0xbf, 0x64, 0xd3, 0x58, 0x5e, 0xe2, 0x2c, 0x5e, 0x0e, 0xcd, 0xff, 0x28,
	0x97, 0x42, 0x09, 0xf4, 0xbe, 0x01, 0x44, 0x16, 0x10, 0x99, 0x17, 0xcb, 0x61, 0xf7, 0xc3, 0x25,
	0x9e, 0x31, 0x82, 0x15, 0x8d, 0xd7, 0x0f, 0x16, 0xde, 0x2d, 0xf5, 0x70, 0xce, 0xb4, 0x4e, 0x26,
	0x24, 0x8d, 0x31, 0x21, 0x92, 0x16, 0x45, 0x09, 0xb8, 0x5f, 0x01, 0x48, 0xb1, 0x50, 0xd4, 0xfe,
	0x2d, 0xdf, 0x3f, 0xb4, 0xef, 0xd5, 0x2a, 0xa7, 0xf1, 0x1c, 0xab, 0xec, 0x8a, 0xca, 0x78, 0x4e,
	0x15, 0x26, 0x58, 0xe1, 0x6a, 0x8d, 0x0a, 0xa4, 0x50, 0x92, 0xf1, 0x69, 0x09, 0xf8, 0x6c, 0x2a,
	0xc4, 0x74, 0x46, 0x4d, 0x11, 0x7a, 0x9d, 0xcb, 0x78, 0x79, 0x88, 0x67, 0xf9, 0x15, 0x3e, 0x8c,
	0x8b, 0x15, 0x57, 0xf8, 0xda, 0xc2, 0xfa, 0x7f, 0xb8, 0xd0, 0x48, 0x93, 0xd1, 0x31, 0xfa, 0x1a,
	0x3c, 0x9c, 0x29, 0x26, 0x78, 0xe8, 0xf4, 0x9c, 0xc1, 0xbd, 0xe1, 0xc3, 0xe8, 0xae, 0xa1, 0x23,
	0x8d, 0x8d, 0x46, 0x06, 0x98, 0x96, 0x04, 0xf4, 0x04, 0xfc, 0x5c, 0xcc, 0x58, 0xc6, 0x68, 0x11,
	0xba, 0xbd, 0xfa, 0xa0, 0x35, 0x1c, 0xbc, 0x81, 0x7c, 0x56, 0x42, 0x4f, 0xb8, 0x92, 0xab, 0x74,
	0xc3, 0xec, 0xbe, 0x80, 0x76, 0xe5, 0x15, 0xea, 0x40, 0xfd, 0x37, 0xba, 0x32, 0xed, 0x04, 0xa9,
	0x7e, 0x44, 0x43, 0x68, 0x2e, 0xf1, 0x6c, 0x41, 0x43, 0xb7, 0xe7, 0x0c, 0x5a, 0xc3, 0x8f, 0xef,
	0xae, 0x62, 0x54, 0x56, 0xa9, 0x85, 0x7e, 0xe3, 0x3e, 0x76, 0xfa, 0x9f, 0x80, 0x67, 0x5b, 0x46,
	0x01, 0x34, 0x47, 0xa7, 0xa7, 0x3f, 0x3f, 0xef, 0xd4, 0x90, 0x0f, 0x8d, 0x27, 0x27, 0xcf, 0x5e,
	0x74, 0x9c, 0xfe, 0x3f, 0x0e, 0x78, 0x96, 0x84, 0x4e, 0xa1, 0x95, 0x53, 0x39, 0x67, 0x45, 0xc1,
	0x04, 0x2f, 0x42, 0xc7, 0x4c, 0xd3, 0x7b, 0x4d, 0x9d, 0x0d, 0x30, 0xf1, 0x6f, 0x92, 0xe6, 0x9f,
	0x8e, 0xeb, 0x3b, 0xe9, 0x36, 0x1d, 0x4d, 0x00, 0x72, 0xc9, 0x78, 0xc6, 0x72, 0x3c, 0x5b, 0xaf,
	0xe6, 0xc1, 0x6b, 0xc4, 0xd6, 0xb8, 0x2d, 0xad, 0x2d, 0x32, 0xfa, 0x16, 0x82, 0x4c, 0x70, 0xc2,
	0xcc, 0x09, 0xd5, 0xcd, 0xf8, 0xf7, 0x23, 0x7b, 0xc4, 0x11, 0xce, 0x59, 0xa4, 0x8f, 0x38, 0x5a,
	0x1f, 0x71, 0x74, 0x72, 0x9d, 0xcb, 0xf4, 0x96, 0xd0, 0xff, 0xab, 0x09, 0x70, 0xdb, 0x2e, 0x3a,
	0x86, 0x00, 0x73, 0x72, 0x21, 0x17, 0x33, 0x5a, 0x98, 0xfd, 0xb6, 0x86, 0x9f, 0xee, 0x9a, 0x31,
	0x3a, 0xa7, 0x6a, 0x5c, 0x4b, 0x7d, 0xcc, 0x49, 0xaa, 0x79, 0x68, 0x04, 0xbe, 0x90, 0xa5, 0x86,
	0xbb, 0x97, 0xc6, 0x81, 0x90, 0x56, 0xe2, 0x23, 0xa8, 0x63, 0xbe, 0x32, 0xe3, 0xf8, 0xc9, 0xc1,
	0x4d, 0xd2, 0xf8, 0xd5, 0xf5, 0x9d, 0x71, 0x2d, 0xd5, 0xbf, 0xa2, 0x23, 0xf0, 0xae, 0x28, 0x26,
	0x54, 0x86, 0x0d, 0xa3, 0xbe, 0x36, 0xa4, 0x9e, 0x76, 0x39, 0x8c, 0x6c, 0x5e, 0xc6, 0x06, 0xf1,
	0xd4, 0x06, 0x60, 0x5c, 0x4b, 0x4b, 0x0a, 0x3a, 0x81, 0x7b, 0x84, 0x16, 0x8a, 0x71, 0xac, 0xe7,
	0xbf, 0x60, 0x79, 0xd8, 0xac, 0x58, 0xa6, 0x14, 0xd1, 0xd9, 0x8c, 0x8e, 0x19, 0x91, 0x29, 0xe6,
	0x53, 0x3a, 0xae, 0xa5, 0xed, 0x2d, 0xd6, 0x24, 0x47, 0x5f, 0x41, 0x67, 0x5b, 0x26, 0x17, 0x52,
	0x85, 0x5e, 0xcf, 0x19, 0xb4, 0x93, 0xe0, 0x26, 0xf1, 0x3e, 0x6f, 0x84, 0xaf, 0x5e, 0xd5, 0xc7,
	0xb5, 0xf4, 0x9d, 0x2d, 0xd0, 0x99, 0x90, 0x4a, 0xef, 0x66, 0x9d, 0xd7, 0xf0, 0xc0, 0x14, 0x7e,
	0x54, 0x16, 0xd6, 0x81, 0x8d, 0xca, 0xc0, 0x46, 0x4f, 0x4b, 0xcc, 0x6d, 0xff, 0x1b, 0x1a, 0xfa,
	0x0e, 0x7c, 0x2e, 0x94, 0xd9, 0x6f, 0xe8, 0x1b, 0x89, 0x9d, 0x36, 0xd4, 0xab, 0xe5, 0x42, 0xe9,
	0xdd, 0xa2, 0xe7, 0xf0, 0x81, 0xa4, 0xbf, 0x2f, 0x68, 0xa1, 0x28, 0xb9, 0x28, 0xa8, 0x5c, 0x52,
	0x79, 0xc1, 0xf1, 0x9c, 0x86, 0x41, 0x65, 0x99, 0x95, 0x76, 0xce, 0xcd, 0xfd, 0x71, 0xdb, 0xcc,
	0x7b, 0x1b, 0x85, 0x73, 0x23, 0xf0, 0x0c, 0xcf, 0x69, 0xf7, 0x47, 0xa8, 0x9f, 0x53, 0x85, 0x7e,
	0x80, 0xe6, 0xda, 0x3e, 0xfb, 0x46, 0xc4, 0x12, 0x93, 0x16, 0x34, 0xf4, 0x03, 0xaa, 0xff, 0x9b,
	0x38, 0xfd, 0xbf, 0x9b, 0x10, 0x6c, 0x22, 0x80, 0xbe, 0x87, 0x03, 0xed, 0x4f, 0x46, 0xd6, 0xee,
	0x7c, 0xb4, 0x23, 0x34, 0xa5, 0xb1, 0x3c, 0xcc, 0xc9, 0x84, 0xe8, 0xb0, 0x78, 0x42, 0x1a, 0xba,
	0xbb, 0x0f, 0xbd, 0x29, 0xa4, 0x66, 0xbf, 0xd1, 0x95, 0xbf, 0x40, 0x1b, 0x2f, 0xd4, 0x15, 0xe5,
	0x8a, 0x65, 0x58, 0x51, 0x52, 0x9a, 0xf3, 0x8b, 0x5d, 0x15, 0x46, 0xdb, 0x24, 0x6d, 0xb4, 0x8a,
	0x0a, 0x3a, 0x82, 0xa0, 0x10, 0x0b, 0x99, 0xd1, 0xb7, 0xb7, 0xaa, 0x6f, 0x09, 0x93, 0x7c, 0x2b,
	0x29, 0xde, 0xfe, 0x49, 0xf9, 0x1f, 0xac, 0xfa, 0x18, 0x3c, 0x6d, 0x55, 0x46, 0x4a, 0xa3, 0xee,
	0xba, 0xe2, 0xf4, 0xaa, 0xb9, 0x50, 0x13, 0xd2, 0x4d, 0xac, 0x99, 0x8e, 0xa0, 0x6e, 0xcf, 0x7a,
	0xcf, 0x0b, 0x52, 0xb3, 0xba, 0x14, 0xda, 0x95, 0xe5, 0xa2, 0x31, 0xdc, 0xdb, 0x5c, 0x9c, 0xd6,
	0xf3, 0xee, 0x5b, 0x7a, 0x3e, 0x6d, 0x6f, 0x88, 0xda, 0xeb, 0x3f, 0x35, 0x7c, 0xa7, 0xe3, 0xa6,
	0x0d, 0xad, 0x91, 0xbc, 0x0b, 0xc0, 0x88, 0x2e, 0xf2, 0x92, 0x51, 0x69, 0x4c, 0x9b, 0x1c, 0x42,
	0x9f, 0x09, 0x2b, 0x9a, 0x4b, 0x71, 0xbd, 0xba, 0xb3, 0xf1, 0x24, 0x48, 0x2f, 0x71, 0x76, 0xa6,
	0x3f, 0xb6, 0x67, 0xce, 0xa5, 0x67, 0xbe, 0xba, 0x5f, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x7f,
	0x62, 0xe2, 0xe4, 0x73, 0x08, 0x00, 0x00,
}