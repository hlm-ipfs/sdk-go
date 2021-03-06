// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hlm-ipfs/proto/market/model.user.proto

package market

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	basic "hlm-ipfs/sdk/proto/basic"
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

type SettingRequest struct {
	LocalAddr            *basic.String `protobuf:"bytes,1,opt,name=LocalAddr,proto3" json:"LocalAddr,omitempty"`
	RemoteAddr           *basic.String `protobuf:"bytes,2,opt,name=RemoteAddr,proto3" json:"RemoteAddr,omitempty"`
	SecretEnable         *basic.Bool   `protobuf:"bytes,3,opt,name=SecretEnable,proto3" json:"SecretEnable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SettingRequest) Reset()         { *m = SettingRequest{} }
func (m *SettingRequest) String() string { return proto.CompactTextString(m) }
func (*SettingRequest) ProtoMessage()    {}
func (*SettingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1b6157669fe8ae, []int{0}
}

func (m *SettingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingRequest.Unmarshal(m, b)
}
func (m *SettingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingRequest.Marshal(b, m, deterministic)
}
func (m *SettingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingRequest.Merge(m, src)
}
func (m *SettingRequest) XXX_Size() int {
	return xxx_messageInfo_SettingRequest.Size(m)
}
func (m *SettingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SettingRequest proto.InternalMessageInfo

func (m *SettingRequest) GetLocalAddr() *basic.String {
	if m != nil {
		return m.LocalAddr
	}
	return nil
}

func (m *SettingRequest) GetRemoteAddr() *basic.String {
	if m != nil {
		return m.RemoteAddr
	}
	return nil
}

func (m *SettingRequest) GetSecretEnable() *basic.Bool {
	if m != nil {
		return m.SecretEnable
	}
	return nil
}

type ProfileInfo struct {
	LocalAddr            string   `protobuf:"bytes,1,opt,name=LocalAddr,proto3" json:"LocalAddr,omitempty"`
	RemoteAddr           string   `protobuf:"bytes,2,opt,name=RemoteAddr,proto3" json:"RemoteAddr,omitempty"`
	SecretEnable         bool     `protobuf:"varint,3,opt,name=SecretEnable,proto3" json:"SecretEnable,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProfileInfo) Reset()         { *m = ProfileInfo{} }
func (m *ProfileInfo) String() string { return proto.CompactTextString(m) }
func (*ProfileInfo) ProtoMessage()    {}
func (*ProfileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e1b6157669fe8ae, []int{1}
}

func (m *ProfileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProfileInfo.Unmarshal(m, b)
}
func (m *ProfileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProfileInfo.Marshal(b, m, deterministic)
}
func (m *ProfileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProfileInfo.Merge(m, src)
}
func (m *ProfileInfo) XXX_Size() int {
	return xxx_messageInfo_ProfileInfo.Size(m)
}
func (m *ProfileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProfileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProfileInfo proto.InternalMessageInfo

func (m *ProfileInfo) GetLocalAddr() string {
	if m != nil {
		return m.LocalAddr
	}
	return ""
}

func (m *ProfileInfo) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *ProfileInfo) GetSecretEnable() bool {
	if m != nil {
		return m.SecretEnable
	}
	return false
}

func init() {
	proto.RegisterType((*SettingRequest)(nil), "market.SettingRequest")
	proto.RegisterType((*ProfileInfo)(nil), "market.ProfileInfo")
}

func init() {
	proto.RegisterFile("hlm-ipfs/proto/market/model.user.proto", fileDescriptor_7e1b6157669fe8ae)
}

var fileDescriptor_7e1b6157669fe8ae = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0xc8, 0xc9, 0xd5,
	0xcd, 0x2c, 0x48, 0x2b, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0xca, 0x4e,
	0x2d, 0xd1, 0xcf, 0xcd, 0x4f, 0x49, 0xcd, 0xd1, 0x2b, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x8b, 0x0b,
	0xb1, 0x41, 0x24, 0xa4, 0x14, 0xd0, 0xd4, 0x27, 0x25, 0x16, 0x67, 0x26, 0x43, 0x48, 0x88, 0x4a,
	0xa5, 0xb9, 0x8c, 0x5c, 0x7c, 0xc1, 0xa9, 0x25, 0x25, 0x99, 0x79, 0xe9, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0xda, 0x5c, 0x9c, 0x3e, 0xf9, 0xc9, 0x89, 0x39, 0x8e, 0x29, 0x29, 0x45,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0xbc, 0x7a, 0x10, 0x3d, 0xc1, 0x25, 0x45, 0x20, 0x85,
	0x08, 0x79, 0x21, 0x5d, 0x2e, 0xae, 0xa0, 0xd4, 0xdc, 0xfc, 0x92, 0x54, 0xb0, 0x6a, 0x26, 0x6c,
	0xaa, 0x91, 0x14, 0x08, 0xe9, 0x73, 0xf1, 0x04, 0xa7, 0x26, 0x17, 0xa5, 0x96, 0xb8, 0xe6, 0x25,
	0x26, 0xe5, 0xa4, 0x4a, 0x30, 0x83, 0x35, 0x70, 0x43, 0x35, 0x38, 0xe5, 0xe7, 0xe7, 0x04, 0xa1,
	0x28, 0x50, 0xca, 0xe7, 0xe2, 0x0e, 0x28, 0xca, 0x4f, 0xcb, 0xcc, 0x49, 0xf5, 0xcc, 0x4b, 0xcb,
	0x17, 0x92, 0x41, 0x77, 0x1b, 0x27, 0xb2, 0x63, 0xe4, 0x30, 0x1c, 0xc3, 0x89, 0x62, 0xbb, 0x12,
	0x16, 0xdb, 0x39, 0x50, 0x2d, 0x74, 0x12, 0x8f, 0x12, 0xc5, 0x1a, 0xc8, 0x49, 0x6c, 0x60, 0x9e,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x4d, 0xd8, 0x2b, 0x7f, 0x84, 0x01, 0x00, 0x00,
}
