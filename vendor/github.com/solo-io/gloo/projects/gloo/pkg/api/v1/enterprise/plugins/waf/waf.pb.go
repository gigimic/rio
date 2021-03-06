// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/waf/waf.proto

package waf

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	waf "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/waf"
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

type Settings struct {
	// Disable waf on this resource (if omitted defaults to false).
	// If a route/virtual host is configured with WAF, you must explicitly disable its WAF,
	// i.e., it will not inherit the disabled status of its parent
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Custom message to display if an intervention occurs.
	CustomInterventionMessage string `protobuf:"bytes,2,opt,name=custom_intervention_message,json=customInterventionMessage,proto3" json:"custom_intervention_message,omitempty"`
	// Add OWASP core rule set
	// if nil will not be added
	CoreRuleSet *CoreRuleSet `protobuf:"bytes,3,opt,name=core_rule_set,json=coreRuleSet,proto3" json:"core_rule_set,omitempty"`
	// Custom rule sets rules to add
	RuleSets             []*waf.RuleSet `protobuf:"bytes,4,rep,name=rule_sets,json=ruleSets,proto3" json:"rule_sets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Settings) Reset()         { *m = Settings{} }
func (m *Settings) String() string { return proto.CompactTextString(m) }
func (*Settings) ProtoMessage()    {}
func (*Settings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f34bd79a50cd0, []int{0}
}
func (m *Settings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Settings.Unmarshal(m, b)
}
func (m *Settings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Settings.Marshal(b, m, deterministic)
}
func (m *Settings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Settings.Merge(m, src)
}
func (m *Settings) XXX_Size() int {
	return xxx_messageInfo_Settings.Size(m)
}
func (m *Settings) XXX_DiscardUnknown() {
	xxx_messageInfo_Settings.DiscardUnknown(m)
}

var xxx_messageInfo_Settings proto.InternalMessageInfo

func (m *Settings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Settings) GetCustomInterventionMessage() string {
	if m != nil {
		return m.CustomInterventionMessage
	}
	return ""
}

func (m *Settings) GetCoreRuleSet() *CoreRuleSet {
	if m != nil {
		return m.CoreRuleSet
	}
	return nil
}

func (m *Settings) GetRuleSets() []*waf.RuleSet {
	if m != nil {
		return m.RuleSets
	}
	return nil
}

type CoreRuleSet struct {
	// Optional custom settings for the OWASP core rule set.
	// For an example on the configuration options see: https://github.com/SpiderLabs/owasp-modsecurity-crs/blob/v3.2/dev/crs-setup.conf.example
	// The same rules apply to these options as do to the `RuleSet`s. The file option is better if possible.
	//
	// Types that are valid to be assigned to CustomSettingsType:
	//	*CoreRuleSet_CustomSettingsString
	//	*CoreRuleSet_CustomSettingsFile
	CustomSettingsType   isCoreRuleSet_CustomSettingsType `protobuf_oneof:"CustomSettingsType"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *CoreRuleSet) Reset()         { *m = CoreRuleSet{} }
func (m *CoreRuleSet) String() string { return proto.CompactTextString(m) }
func (*CoreRuleSet) ProtoMessage()    {}
func (*CoreRuleSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f34bd79a50cd0, []int{1}
}
func (m *CoreRuleSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoreRuleSet.Unmarshal(m, b)
}
func (m *CoreRuleSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoreRuleSet.Marshal(b, m, deterministic)
}
func (m *CoreRuleSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoreRuleSet.Merge(m, src)
}
func (m *CoreRuleSet) XXX_Size() int {
	return xxx_messageInfo_CoreRuleSet.Size(m)
}
func (m *CoreRuleSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CoreRuleSet.DiscardUnknown(m)
}

var xxx_messageInfo_CoreRuleSet proto.InternalMessageInfo

type isCoreRuleSet_CustomSettingsType interface {
	isCoreRuleSet_CustomSettingsType()
	Equal(interface{}) bool
}

type CoreRuleSet_CustomSettingsString struct {
	CustomSettingsString string `protobuf:"bytes,2,opt,name=custom_settings_string,json=customSettingsString,proto3,oneof" json:"custom_settings_string,omitempty"`
}
type CoreRuleSet_CustomSettingsFile struct {
	CustomSettingsFile string `protobuf:"bytes,3,opt,name=custom_settings_file,json=customSettingsFile,proto3,oneof" json:"custom_settings_file,omitempty"`
}

func (*CoreRuleSet_CustomSettingsString) isCoreRuleSet_CustomSettingsType() {}
func (*CoreRuleSet_CustomSettingsFile) isCoreRuleSet_CustomSettingsType()   {}

func (m *CoreRuleSet) GetCustomSettingsType() isCoreRuleSet_CustomSettingsType {
	if m != nil {
		return m.CustomSettingsType
	}
	return nil
}

func (m *CoreRuleSet) GetCustomSettingsString() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsString); ok {
		return x.CustomSettingsString
	}
	return ""
}

func (m *CoreRuleSet) GetCustomSettingsFile() string {
	if x, ok := m.GetCustomSettingsType().(*CoreRuleSet_CustomSettingsFile); ok {
		return x.CustomSettingsFile
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CoreRuleSet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CoreRuleSet_CustomSettingsString)(nil),
		(*CoreRuleSet_CustomSettingsFile)(nil),
	}
}

// TODO(kdorosh) delete this once we stop supporting opaque configuration under extensions
//
// Deprecated: Do not use.
type VhostSettings struct {
	// Disable waf on this virtual host
	Disabled             bool      `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Settings             *Settings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *VhostSettings) Reset()         { *m = VhostSettings{} }
func (m *VhostSettings) String() string { return proto.CompactTextString(m) }
func (*VhostSettings) ProtoMessage()    {}
func (*VhostSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f34bd79a50cd0, []int{2}
}
func (m *VhostSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VhostSettings.Unmarshal(m, b)
}
func (m *VhostSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VhostSettings.Marshal(b, m, deterministic)
}
func (m *VhostSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VhostSettings.Merge(m, src)
}
func (m *VhostSettings) XXX_Size() int {
	return xxx_messageInfo_VhostSettings.Size(m)
}
func (m *VhostSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_VhostSettings.DiscardUnknown(m)
}

var xxx_messageInfo_VhostSettings proto.InternalMessageInfo

func (m *VhostSettings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *VhostSettings) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

// TODO(kdorosh) delete this once we stop supporting opaque configuration under extensions
//
// Deprecated: Do not use.
type RouteSettings struct {
	// Disable waf on this route
	Disabled             bool      `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Settings             *Settings `protobuf:"bytes,2,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RouteSettings) Reset()         { *m = RouteSettings{} }
func (m *RouteSettings) String() string { return proto.CompactTextString(m) }
func (*RouteSettings) ProtoMessage()    {}
func (*RouteSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f15f34bd79a50cd0, []int{3}
}
func (m *RouteSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteSettings.Unmarshal(m, b)
}
func (m *RouteSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteSettings.Marshal(b, m, deterministic)
}
func (m *RouteSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteSettings.Merge(m, src)
}
func (m *RouteSettings) XXX_Size() int {
	return xxx_messageInfo_RouteSettings.Size(m)
}
func (m *RouteSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RouteSettings proto.InternalMessageInfo

func (m *RouteSettings) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *RouteSettings) GetSettings() *Settings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*Settings)(nil), "waf.plugins.gloo.solo.io.Settings")
	proto.RegisterType((*CoreRuleSet)(nil), "waf.plugins.gloo.solo.io.CoreRuleSet")
	proto.RegisterType((*VhostSettings)(nil), "waf.plugins.gloo.solo.io.VhostSettings")
	proto.RegisterType((*RouteSettings)(nil), "waf.plugins.gloo.solo.io.RouteSettings")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/plugins/waf/waf.proto", fileDescriptor_f15f34bd79a50cd0)
}

var fileDescriptor_f15f34bd79a50cd0 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x2d, 0x42, 0xee, 0x5a, 0xbd, 0xac, 0x22, 0x64, 0x82, 0x84, 0xa2, 0x48, 0x48, 0xb9,
	0xb0, 0x0b, 0x46, 0xe2, 0xc0, 0x81, 0x43, 0x2b, 0x01, 0x3d, 0xe4, 0x80, 0x83, 0x38, 0x70, 0xb1,
	0x1c, 0x67, 0xb2, 0x59, 0xd8, 0xec, 0x58, 0x3b, 0xe3, 0x40, 0x6e, 0xfc, 0x01, 0xbf, 0xc1, 0x77,
	0xf1, 0x25, 0xc8, 0x76, 0x5d, 0xda, 0x4a, 0x45, 0x39, 0x71, 0xb0, 0x34, 0xb3, 0x33, 0xef, 0x3d,
	0xbf, 0xe7, 0xb5, 0x98, 0x1b, 0xcb, 0x9b, 0x66, 0xa9, 0x2a, 0xdc, 0x6a, 0x42, 0x87, 0xcf, 0x2c,
	0x6a, 0xe3, 0x10, 0x75, 0x1d, 0xf0, 0x0b, 0x54, 0x4c, 0x7d, 0x57, 0xd6, 0x56, 0xef, 0x5e, 0x68,
	0xf0, 0x0c, 0xa1, 0x0e, 0x96, 0x40, 0xd7, 0xae, 0x31, 0xd6, 0x93, 0xfe, 0x56, 0xae, 0xdb, 0x47,
	0xd5, 0x01, 0x19, 0x65, 0xda, 0x95, 0xfd, 0x48, 0xb5, 0x48, 0xd5, 0x92, 0x2a, 0x8b, 0xe3, 0x0f,
	0x87, 0x0b, 0xc1, 0x77, 0x86, 0xe0, 0x4b, 0xa7, 0xc1, 0xef, 0x70, 0xdf, 0xb5, 0x9e, 0x2c, 0xde,
	0x16, 0x1b, 0x8f, 0x0c, 0x1a, 0xec, 0x4a, 0xdd, 0x56, 0xfd, 0xe9, 0xf4, 0xc7, 0x91, 0x88, 0x17,
	0xc0, 0x6c, 0xbd, 0x21, 0x39, 0x16, 0xf1, 0xca, 0x52, 0xb9, 0x74, 0xb0, 0x4a, 0xa3, 0x49, 0x34,
	0x8b, 0xf3, 0xab, 0x5e, 0xbe, 0x11, 0x8f, 0xab, 0x86, 0x18, 0xb7, 0x85, 0x6d, 0x6d, 0xed, 0xc0,
	0xb3, 0x45, 0x5f, 0x6c, 0x81, 0xa8, 0x34, 0x90, 0x1e, 0x4d, 0xa2, 0xd9, 0x49, 0xfe, 0xa8, 0x5f,
	0xb9, 0xb8, 0xb6, 0x31, 0xef, 0x17, 0xe4, 0x85, 0x38, 0xad, 0x30, 0x40, 0x11, 0x1a, 0x07, 0x05,
	0x01, 0xa7, 0xc7, 0x93, 0x68, 0x96, 0x64, 0x4f, 0xd5, 0x5d, 0x19, 0xa8, 0x73, 0x0c, 0x90, 0x37,
	0x0e, 0x16, 0xc0, 0x79, 0x52, 0xfd, 0x6d, 0xe4, 0x5c, 0x9c, 0x0c, 0x2c, 0x94, 0xde, 0x9f, 0x1c,
	0xcf, 0x92, 0xec, 0xb9, 0xea, 0xdc, 0xab, 0x0a, 0xfd, 0xda, 0x1a, 0xb5, 0xb6, 0x8e, 0x21, 0xa8,
	0x0d, 0x73, 0xad, 0xb6, 0xb8, 0x22, 0xa8, 0x9a, 0x60, 0x79, 0xaf, 0x76, 0x99, 0x1a, 0x18, 0xe3,
	0xd0, 0x17, 0x34, 0xfd, 0x19, 0x89, 0xe4, 0x9a, 0x96, 0x7c, 0x25, 0x1e, 0x5e, 0x3a, 0xa5, 0xcb,
	0x60, 0x0a, 0xe2, 0x60, 0xbd, 0xe9, 0x4d, 0xbe, 0xbf, 0x97, 0x8f, 0xfa, 0xf9, 0x90, 0xdb, 0xa2,
	0x9b, 0xca, 0x4c, 0x8c, 0x6e, 0xe3, 0xd6, 0xd6, 0x41, 0x67, 0xb4, 0x45, 0xc9, 0x9b, 0xa8, 0xb7,
	0xd6, 0xc1, 0xd9, 0x48, 0xc8, 0xf3, 0x1b, 0xa7, 0x1f, 0xf7, 0x35, 0x4c, 0x51, 0x9c, 0x7e, 0xda,
	0x20, 0xf1, 0x81, 0x1f, 0x26, 0x1e, 0xf4, 0xba, 0x17, 0x4c, 0xb2, 0xe9, 0xdd, 0x99, 0x0e, 0x8c,
	0xf9, 0x15, 0xe6, 0xf5, 0x51, 0x1a, 0xb5, 0x82, 0x39, 0x36, 0x0c, 0xff, 0x4b, 0xf0, 0x6c, 0xfe,
	0xeb, 0xf7, 0x93, 0xe8, 0xf3, 0xbb, 0xc3, 0x6e, 0x79, 0xfd, 0xd5, 0xfc, 0xfb, 0x97, 0x5a, 0x3e,
	0xe8, 0x2e, 0xf3, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x9d, 0x1b, 0xc4, 0xa0, 0x03,
	0x00, 0x00,
}

func (this *Settings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Settings)
	if !ok {
		that2, ok := that.(Settings)
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
	if this.Disabled != that1.Disabled {
		return false
	}
	if this.CustomInterventionMessage != that1.CustomInterventionMessage {
		return false
	}
	if !this.CoreRuleSet.Equal(that1.CoreRuleSet) {
		return false
	}
	if len(this.RuleSets) != len(that1.RuleSets) {
		return false
	}
	for i := range this.RuleSets {
		if !this.RuleSets[i].Equal(that1.RuleSets[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet)
	if !ok {
		that2, ok := that.(CoreRuleSet)
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
	if that1.CustomSettingsType == nil {
		if this.CustomSettingsType != nil {
			return false
		}
	} else if this.CustomSettingsType == nil {
		return false
	} else if !this.CustomSettingsType.Equal(that1.CustomSettingsType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsString) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsString)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsString)
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
	if this.CustomSettingsString != that1.CustomSettingsString {
		return false
	}
	return true
}
func (this *CoreRuleSet_CustomSettingsFile) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CoreRuleSet_CustomSettingsFile)
	if !ok {
		that2, ok := that.(CoreRuleSet_CustomSettingsFile)
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
	if this.CustomSettingsFile != that1.CustomSettingsFile {
		return false
	}
	return true
}
func (this *VhostSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*VhostSettings)
	if !ok {
		that2, ok := that.(VhostSettings)
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
	if this.Disabled != that1.Disabled {
		return false
	}
	if !this.Settings.Equal(that1.Settings) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RouteSettings) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteSettings)
	if !ok {
		that2, ok := that.(RouteSettings)
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
	if this.Disabled != that1.Disabled {
		return false
	}
	if !this.Settings.Equal(that1.Settings) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
