// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hlm-ipfs/proto/market/bucket.proto

package market

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "hlm-ipfs/sdk/proto/basic"
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

func init() { proto.RegisterFile("hlm-ipfs/proto/market/bucket.proto", fileDescriptor_75a3fe955f3bae55) }

var fileDescriptor_75a3fe955f3bae55 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xca, 0xc8, 0xc9, 0xd5,
	0xcd, 0x2c, 0x48, 0x2b, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0xca, 0x4e,
	0x2d, 0xd1, 0x4f, 0x2a, 0x4d, 0xce, 0x4e, 0x2d, 0xd1, 0x03, 0x8b, 0x09, 0xb1, 0x41, 0x04, 0xa5,
	0x14, 0xd0, 0xd4, 0x26, 0x25, 0x16, 0x67, 0x26, 0x43, 0x48, 0x88, 0x4a, 0x29, 0x0d, 0xec, 0xa6,
	0xe5, 0xe6, 0xa7, 0xa4, 0xe6, 0xe8, 0x21, 0x9b, 0x69, 0xf4, 0x8f, 0x91, 0x8b, 0xcd, 0x09, 0x2c,
	0x20, 0x64, 0xc4, 0xc5, 0xea, 0x97, 0x9f, 0x92, 0x5a, 0x2c, 0xc4, 0xa3, 0x07, 0x31, 0xcb, 0x33,
	0xaf, 0xc4, 0xcc, 0x44, 0x4a, 0x5a, 0x0f, 0xa2, 0x5b, 0x0f, 0xa2, 0x0c, 0xac, 0x24, 0x28, 0xb5,
	0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x80, 0x8b, 0xc5, 0x27, 0xb3, 0xb8, 0x04, 0x4d, 0x8b,
	0x14, 0xaa, 0x16, 0x90, 0x0a, 0xb8, 0x0e, 0x63, 0x2e, 0x36, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0x54,
	0x21, 0x34, 0x83, 0x21, 0xa2, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x52, 0x28, 0x06, 0x82,
	0x34, 0x85, 0x16, 0xa4, 0x60, 0xd1, 0x04, 0x11, 0x45, 0xd7, 0xe4, 0x9a, 0x5b, 0x50, 0x52, 0x29,
	0xa4, 0xc2, 0xc5, 0x16, 0x94, 0x9a, 0x9b, 0x5f, 0x96, 0x8a, 0xe6, 0x3a, 0x14, 0x55, 0x4e, 0xe2,
	0x51, 0xa2, 0x58, 0x03, 0x2b, 0x89, 0x0d, 0xcc, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x7b, 0xcd, 0xb5, 0x9a, 0x01, 0x00, 0x00,
}