// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hlm-ipfs/proto/market/storage.proto

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

func init() {
	proto.RegisterFile("hlm-ipfs/proto/market/storage.proto", fileDescriptor_c6d1a0488f2a7096)
}

var fileDescriptor_c6d1a0488f2a7096 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xce, 0xc8, 0xc9, 0xd5,
	0xcd, 0x2c, 0x48, 0x2b, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0x4d, 0x2c, 0xca, 0x4e,
	0x2d, 0xd1, 0x2f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0xd5, 0x03, 0x0b, 0x0a, 0xb1, 0x41, 0x44,
	0xa5, 0x14, 0xd0, 0x14, 0x27, 0x25, 0x16, 0x67, 0x26, 0x43, 0x48, 0x88, 0x4a, 0x29, 0x4d, 0xec,
	0xc6, 0xe5, 0xe6, 0xa7, 0xa4, 0xe6, 0xe8, 0xa1, 0x18, 0x6a, 0xd4, 0xc5, 0xc4, 0xc5, 0x1e, 0x0c,
	0x11, 0x11, 0xd2, 0xe2, 0x62, 0x71, 0xcb, 0xcc, 0x4b, 0x11, 0xe2, 0xd5, 0x83, 0x18, 0x16, 0x5c,
	0x52, 0x94, 0x99, 0x97, 0x2e, 0x25, 0xac, 0x07, 0xd1, 0xaf, 0x07, 0x55, 0xe7, 0x99, 0x97, 0x96,
	0x2f, 0x64, 0xcf, 0xc5, 0xe2, 0x93, 0x59, 0x5c, 0x22, 0x24, 0x85, 0x26, 0x09, 0x12, 0x0c, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x91, 0x92, 0xc6, 0x2a, 0x57, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a,
	0x64, 0xca, 0xc5, 0xe6, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x2a, 0x24, 0x83, 0xa6, 0x0c, 0x22, 0x0c,
	0x33, 0x04, 0xd5, 0x31, 0x42, 0x26, 0x5c, 0x6c, 0xa1, 0x05, 0x29, 0xd8, 0xb4, 0x41, 0x84, 0x61,
	0xda, 0x78, 0xa0, 0xda, 0x5c, 0x73, 0x0b, 0x4a, 0x2a, 0x85, 0x54, 0xb9, 0xd8, 0x82, 0x52, 0x73,
	0xf3, 0xcb, 0x52, 0xd1, 0xfd, 0x86, 0xa2, 0xcc, 0x49, 0x3c, 0x4a, 0x14, 0x6b, 0xc8, 0x25, 0xb1,
	0x81, 0x79, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xae, 0x2e, 0x58, 0xa8, 0x01, 0x00,
	0x00,
}