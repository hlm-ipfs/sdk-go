// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hlm-ipfs/proto/basic/basic.proto

package basic

import (
	fmt "fmt"
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

type DateSpan int32

const (
	DateSpan_DateSpan_Day   DateSpan = 0
	DateSpan_DateSpan_Weak  DateSpan = 1
	DateSpan_DateSpan_Month DateSpan = 2
	DateSpan_DateSpan_Year  DateSpan = 3
)

var DateSpan_name = map[int32]string{
	0: "DateSpan_Day",
	1: "DateSpan_Weak",
	2: "DateSpan_Month",
	3: "DateSpan_Year",
}

var DateSpan_value = map[string]int32{
	"DateSpan_Day":   0,
	"DateSpan_Weak":  1,
	"DateSpan_Month": 2,
	"DateSpan_Year":  3,
}

func (x DateSpan) String() string {
	return proto.EnumName(DateSpan_name, int32(x))
}

func (DateSpan) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{1}
}

func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (m *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(m, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Strings struct {
	Items                []string `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Strings) Reset()         { *m = Strings{} }
func (m *Strings) String() string { return proto.CompactTextString(m) }
func (*Strings) ProtoMessage()    {}
func (*Strings) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{2}
}

func (m *Strings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Strings.Unmarshal(m, b)
}
func (m *Strings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Strings.Marshal(b, m, deterministic)
}
func (m *Strings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Strings.Merge(m, src)
}
func (m *Strings) XXX_Size() int {
	return xxx_messageInfo_Strings.Size(m)
}
func (m *Strings) XXX_DiscardUnknown() {
	xxx_messageInfo_Strings.DiscardUnknown(m)
}

var xxx_messageInfo_Strings proto.InternalMessageInfo

func (m *Strings) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

type Int32 struct {
	Value                int32    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32) Reset()         { *m = Int32{} }
func (m *Int32) String() string { return proto.CompactTextString(m) }
func (*Int32) ProtoMessage()    {}
func (*Int32) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{3}
}

func (m *Int32) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32.Unmarshal(m, b)
}
func (m *Int32) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32.Marshal(b, m, deterministic)
}
func (m *Int32) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32.Merge(m, src)
}
func (m *Int32) XXX_Size() int {
	return xxx_messageInfo_Int32.Size(m)
}
func (m *Int32) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32.DiscardUnknown(m)
}

var xxx_messageInfo_Int32 proto.InternalMessageInfo

func (m *Int32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int32S struct {
	Items                []int32  `protobuf:"varint,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int32S) Reset()         { *m = Int32S{} }
func (m *Int32S) String() string { return proto.CompactTextString(m) }
func (*Int32S) ProtoMessage()    {}
func (*Int32S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{4}
}

func (m *Int32S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int32S.Unmarshal(m, b)
}
func (m *Int32S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int32S.Marshal(b, m, deterministic)
}
func (m *Int32S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int32S.Merge(m, src)
}
func (m *Int32S) XXX_Size() int {
	return xxx_messageInfo_Int32S.Size(m)
}
func (m *Int32S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int32S.DiscardUnknown(m)
}

var xxx_messageInfo_Int32S proto.InternalMessageInfo

func (m *Int32S) GetItems() []int32 {
	if m != nil {
		return m.Items
	}
	return nil
}

type Int64 struct {
	Value                int64    `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64) Reset()         { *m = Int64{} }
func (m *Int64) String() string { return proto.CompactTextString(m) }
func (*Int64) ProtoMessage()    {}
func (*Int64) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{5}
}

func (m *Int64) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64.Unmarshal(m, b)
}
func (m *Int64) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64.Marshal(b, m, deterministic)
}
func (m *Int64) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64.Merge(m, src)
}
func (m *Int64) XXX_Size() int {
	return xxx_messageInfo_Int64.Size(m)
}
func (m *Int64) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64.DiscardUnknown(m)
}

var xxx_messageInfo_Int64 proto.InternalMessageInfo

func (m *Int64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int64S struct {
	Items                []int64  `protobuf:"varint,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64S) Reset()         { *m = Int64S{} }
func (m *Int64S) String() string { return proto.CompactTextString(m) }
func (*Int64S) ProtoMessage()    {}
func (*Int64S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{6}
}

func (m *Int64S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64S.Unmarshal(m, b)
}
func (m *Int64S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64S.Marshal(b, m, deterministic)
}
func (m *Int64S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64S.Merge(m, src)
}
func (m *Int64S) XXX_Size() int {
	return xxx_messageInfo_Int64S.Size(m)
}
func (m *Int64S) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64S.DiscardUnknown(m)
}

var xxx_messageInfo_Int64S proto.InternalMessageInfo

func (m *Int64S) GetItems() []int64 {
	if m != nil {
		return m.Items
	}
	return nil
}

type Float struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Float) Reset()         { *m = Float{} }
func (m *Float) String() string { return proto.CompactTextString(m) }
func (*Float) ProtoMessage()    {}
func (*Float) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{7}
}

func (m *Float) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float.Unmarshal(m, b)
}
func (m *Float) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float.Marshal(b, m, deterministic)
}
func (m *Float) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float.Merge(m, src)
}
func (m *Float) XXX_Size() int {
	return xxx_messageInfo_Float.Size(m)
}
func (m *Float) XXX_DiscardUnknown() {
	xxx_messageInfo_Float.DiscardUnknown(m)
}

var xxx_messageInfo_Float proto.InternalMessageInfo

func (m *Float) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Floats struct {
	Items                []float32 `protobuf:"fixed32,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Floats) Reset()         { *m = Floats{} }
func (m *Floats) String() string { return proto.CompactTextString(m) }
func (*Floats) ProtoMessage()    {}
func (*Floats) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{8}
}

func (m *Floats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Floats.Unmarshal(m, b)
}
func (m *Floats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Floats.Marshal(b, m, deterministic)
}
func (m *Floats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Floats.Merge(m, src)
}
func (m *Floats) XXX_Size() int {
	return xxx_messageInfo_Floats.Size(m)
}
func (m *Floats) XXX_DiscardUnknown() {
	xxx_messageInfo_Floats.DiscardUnknown(m)
}

var xxx_messageInfo_Floats proto.InternalMessageInfo

func (m *Floats) GetItems() []float32 {
	if m != nil {
		return m.Items
	}
	return nil
}

type Double struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Double) Reset()         { *m = Double{} }
func (m *Double) String() string { return proto.CompactTextString(m) }
func (*Double) ProtoMessage()    {}
func (*Double) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{9}
}

func (m *Double) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Double.Unmarshal(m, b)
}
func (m *Double) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Double.Marshal(b, m, deterministic)
}
func (m *Double) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Double.Merge(m, src)
}
func (m *Double) XXX_Size() int {
	return xxx_messageInfo_Double.Size(m)
}
func (m *Double) XXX_DiscardUnknown() {
	xxx_messageInfo_Double.DiscardUnknown(m)
}

var xxx_messageInfo_Double proto.InternalMessageInfo

func (m *Double) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Doubles struct {
	Items                []float64 `protobuf:"fixed64,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Doubles) Reset()         { *m = Doubles{} }
func (m *Doubles) String() string { return proto.CompactTextString(m) }
func (*Doubles) ProtoMessage()    {}
func (*Doubles) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{10}
}

func (m *Doubles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Doubles.Unmarshal(m, b)
}
func (m *Doubles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Doubles.Marshal(b, m, deterministic)
}
func (m *Doubles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Doubles.Merge(m, src)
}
func (m *Doubles) XXX_Size() int {
	return xxx_messageInfo_Doubles.Size(m)
}
func (m *Doubles) XXX_DiscardUnknown() {
	xxx_messageInfo_Doubles.DiscardUnknown(m)
}

var xxx_messageInfo_Doubles proto.InternalMessageInfo

func (m *Doubles) GetItems() []float64 {
	if m != nil {
		return m.Items
	}
	return nil
}

type Bool struct {
	Value                bool     `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bool) Reset()         { *m = Bool{} }
func (m *Bool) String() string { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()    {}
func (*Bool) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{11}
}

func (m *Bool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bool.Unmarshal(m, b)
}
func (m *Bool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bool.Marshal(b, m, deterministic)
}
func (m *Bool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bool.Merge(m, src)
}
func (m *Bool) XXX_Size() int {
	return xxx_messageInfo_Bool.Size(m)
}
func (m *Bool) XXX_DiscardUnknown() {
	xxx_messageInfo_Bool.DiscardUnknown(m)
}

var xxx_messageInfo_Bool proto.InternalMessageInfo

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type Bools struct {
	Items                []bool   `protobuf:"varint,1,rep,packed,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bools) Reset()         { *m = Bools{} }
func (m *Bools) String() string { return proto.CompactTextString(m) }
func (*Bools) ProtoMessage()    {}
func (*Bools) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{12}
}

func (m *Bools) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bools.Unmarshal(m, b)
}
func (m *Bools) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bools.Marshal(b, m, deterministic)
}
func (m *Bools) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bools.Merge(m, src)
}
func (m *Bools) XXX_Size() int {
	return xxx_messageInfo_Bools.Size(m)
}
func (m *Bools) XXX_DiscardUnknown() {
	xxx_messageInfo_Bools.DiscardUnknown(m)
}

var xxx_messageInfo_Bools proto.InternalMessageInfo

func (m *Bools) GetItems() []bool {
	if m != nil {
		return m.Items
	}
	return nil
}

type Bytes struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bytes) Reset()         { *m = Bytes{} }
func (m *Bytes) String() string { return proto.CompactTextString(m) }
func (*Bytes) ProtoMessage()    {}
func (*Bytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{13}
}

func (m *Bytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bytes.Unmarshal(m, b)
}
func (m *Bytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bytes.Marshal(b, m, deterministic)
}
func (m *Bytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bytes.Merge(m, src)
}
func (m *Bytes) XXX_Size() int {
	return xxx_messageInfo_Bytes.Size(m)
}
func (m *Bytes) XXX_DiscardUnknown() {
	xxx_messageInfo_Bytes.DiscardUnknown(m)
}

var xxx_messageInfo_Bytes proto.InternalMessageInfo

func (m *Bytes) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Byteses struct {
	Items                [][]byte `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Byteses) Reset()         { *m = Byteses{} }
func (m *Byteses) String() string { return proto.CompactTextString(m) }
func (*Byteses) ProtoMessage()    {}
func (*Byteses) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{14}
}

func (m *Byteses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Byteses.Unmarshal(m, b)
}
func (m *Byteses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Byteses.Marshal(b, m, deterministic)
}
func (m *Byteses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Byteses.Merge(m, src)
}
func (m *Byteses) XXX_Size() int {
	return xxx_messageInfo_Byteses.Size(m)
}
func (m *Byteses) XXX_DiscardUnknown() {
	xxx_messageInfo_Byteses.DiscardUnknown(m)
}

var xxx_messageInfo_Byteses proto.InternalMessageInfo

func (m *Byteses) GetItems() [][]byte {
	if m != nil {
		return m.Items
	}
	return nil
}

type KVStringAtString struct {
	Key                  string   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVStringAtString) Reset()         { *m = KVStringAtString{} }
func (m *KVStringAtString) String() string { return proto.CompactTextString(m) }
func (*KVStringAtString) ProtoMessage()    {}
func (*KVStringAtString) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{15}
}

func (m *KVStringAtString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVStringAtString.Unmarshal(m, b)
}
func (m *KVStringAtString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVStringAtString.Marshal(b, m, deterministic)
}
func (m *KVStringAtString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVStringAtString.Merge(m, src)
}
func (m *KVStringAtString) XXX_Size() int {
	return xxx_messageInfo_KVStringAtString.Size(m)
}
func (m *KVStringAtString) XXX_DiscardUnknown() {
	xxx_messageInfo_KVStringAtString.DiscardUnknown(m)
}

var xxx_messageInfo_KVStringAtString proto.InternalMessageInfo

func (m *KVStringAtString) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KVStringAtString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KVInt64AtString struct {
	Key                  int64    `protobuf:"varint,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVInt64AtString) Reset()         { *m = KVInt64AtString{} }
func (m *KVInt64AtString) String() string { return proto.CompactTextString(m) }
func (*KVInt64AtString) ProtoMessage()    {}
func (*KVInt64AtString) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{16}
}

func (m *KVInt64AtString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVInt64AtString.Unmarshal(m, b)
}
func (m *KVInt64AtString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVInt64AtString.Marshal(b, m, deterministic)
}
func (m *KVInt64AtString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVInt64AtString.Merge(m, src)
}
func (m *KVInt64AtString) XXX_Size() int {
	return xxx_messageInfo_KVInt64AtString.Size(m)
}
func (m *KVInt64AtString) XXX_DiscardUnknown() {
	xxx_messageInfo_KVInt64AtString.DiscardUnknown(m)
}

var xxx_messageInfo_KVInt64AtString proto.InternalMessageInfo

func (m *KVInt64AtString) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *KVInt64AtString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KVInt64AtInt64S struct {
	Key                  int64    `protobuf:"varint,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Values               []int64  `protobuf:"varint,2,rep,packed,name=Values,proto3" json:"Values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KVInt64AtInt64S) Reset()         { *m = KVInt64AtInt64S{} }
func (m *KVInt64AtInt64S) String() string { return proto.CompactTextString(m) }
func (*KVInt64AtInt64S) ProtoMessage()    {}
func (*KVInt64AtInt64S) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{17}
}

func (m *KVInt64AtInt64S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KVInt64AtInt64S.Unmarshal(m, b)
}
func (m *KVInt64AtInt64S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KVInt64AtInt64S.Marshal(b, m, deterministic)
}
func (m *KVInt64AtInt64S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KVInt64AtInt64S.Merge(m, src)
}
func (m *KVInt64AtInt64S) XXX_Size() int {
	return xxx_messageInfo_KVInt64AtInt64S.Size(m)
}
func (m *KVInt64AtInt64S) XXX_DiscardUnknown() {
	xxx_messageInfo_KVInt64AtInt64S.DiscardUnknown(m)
}

var xxx_messageInfo_KVInt64AtInt64S proto.InternalMessageInfo

func (m *KVInt64AtInt64S) GetKey() int64 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *KVInt64AtInt64S) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type File struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	DataStr              string   `protobuf:"bytes,3,opt,name=DataStr,proto3" json:"DataStr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *File) Reset()         { *m = File{} }
func (m *File) String() string { return proto.CompactTextString(m) }
func (*File) ProtoMessage()    {}
func (*File) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{18}
}

func (m *File) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_File.Unmarshal(m, b)
}
func (m *File) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_File.Marshal(b, m, deterministic)
}
func (m *File) XXX_Merge(src proto.Message) {
	xxx_messageInfo_File.Merge(m, src)
}
func (m *File) XXX_Size() int {
	return xxx_messageInfo_File.Size(m)
}
func (m *File) XXX_DiscardUnknown() {
	xxx_messageInfo_File.DiscardUnknown(m)
}

var xxx_messageInfo_File proto.InternalMessageInfo

func (m *File) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *File) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *File) GetDataStr() string {
	if m != nil {
		return m.DataStr
	}
	return ""
}

type Files struct {
	Items                []*File  `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Files) Reset()         { *m = Files{} }
func (m *Files) String() string { return proto.CompactTextString(m) }
func (*Files) ProtoMessage()    {}
func (*Files) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d3506d33fff2b32, []int{19}
}

func (m *Files) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Files.Unmarshal(m, b)
}
func (m *Files) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Files.Marshal(b, m, deterministic)
}
func (m *Files) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Files.Merge(m, src)
}
func (m *Files) XXX_Size() int {
	return xxx_messageInfo_Files.Size(m)
}
func (m *Files) XXX_DiscardUnknown() {
	xxx_messageInfo_Files.DiscardUnknown(m)
}

var xxx_messageInfo_Files proto.InternalMessageInfo

func (m *Files) GetItems() []*File {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterEnum("basic.DateSpan", DateSpan_name, DateSpan_value)
	proto.RegisterType((*Empty)(nil), "basic.Empty")
	proto.RegisterType((*String)(nil), "basic.String")
	proto.RegisterType((*Strings)(nil), "basic.Strings")
	proto.RegisterType((*Int32)(nil), "basic.Int32")
	proto.RegisterType((*Int32S)(nil), "basic.Int32s")
	proto.RegisterType((*Int64)(nil), "basic.Int64")
	proto.RegisterType((*Int64S)(nil), "basic.Int64s")
	proto.RegisterType((*Float)(nil), "basic.Float")
	proto.RegisterType((*Floats)(nil), "basic.Floats")
	proto.RegisterType((*Double)(nil), "basic.Double")
	proto.RegisterType((*Doubles)(nil), "basic.Doubles")
	proto.RegisterType((*Bool)(nil), "basic.Bool")
	proto.RegisterType((*Bools)(nil), "basic.Bools")
	proto.RegisterType((*Bytes)(nil), "basic.Bytes")
	proto.RegisterType((*Byteses)(nil), "basic.Byteses")
	proto.RegisterType((*KVStringAtString)(nil), "basic.KVStringAtString")
	proto.RegisterType((*KVInt64AtString)(nil), "basic.KVInt64AtString")
	proto.RegisterType((*KVInt64AtInt64S)(nil), "basic.KVInt64AtInt64s")
	proto.RegisterType((*File)(nil), "basic.File")
	proto.RegisterType((*Files)(nil), "basic.Files")
}

func init() { proto.RegisterFile("hlm-ipfs/proto/basic/basic.proto", fileDescriptor_3d3506d33fff2b32) }

var fileDescriptor_3d3506d33fff2b32 = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x4f, 0xcf, 0xd2, 0x40,
	0x10, 0xc6, 0x6d, 0xcb, 0x16, 0x1c, 0xaa, 0xd6, 0x0d, 0x21, 0x3d, 0x08, 0x62, 0x4f, 0x86, 0x44,
	0x48, 0x80, 0x90, 0xa8, 0x27, 0x49, 0x25, 0x12, 0xa2, 0x07, 0x48, 0x6a, 0xf4, 0x62, 0x16, 0xb3,
	0x4a, 0x63, 0xff, 0xa5, 0x5d, 0x0e, 0xfd, 0xf6, 0x66, 0xb7, 0x5d, 0xd8, 0xed, 0x4b, 0xde, 0x0b,
	0xcc, 0xcc, 0x33, 0xbf, 0xd9, 0x76, 0x9e, 0x2d, 0x4c, 0xce, 0x71, 0xf2, 0x2e, 0xca, 0xff, 0x94,
	0xf3, 0xbc, 0xc8, 0x58, 0x36, 0x3f, 0x91, 0x32, 0xfa, 0x5d, 0xff, 0xce, 0x44, 0x05, 0x23, 0x91,
	0xf8, 0x5d, 0x40, 0x9f, 0x93, 0x9c, 0x55, 0xfe, 0x18, 0xec, 0x23, 0x2b, 0xa2, 0xf4, 0x2f, 0x1e,
	0x00, 0x0a, 0x49, 0x7c, 0xa1, 0x9e, 0x31, 0x31, 0xde, 0x3e, 0x3d, 0xd4, 0x89, 0xff, 0x1a, 0xba,
	0xb5, 0x5e, 0xf2, 0x86, 0x1d, 0xa3, 0x49, 0xe9, 0x19, 0x13, 0x8b, 0x37, 0x88, 0xc4, 0x1f, 0x01,
	0xda, 0xa5, 0x6c, 0xb9, 0xd0, 0x79, 0x24, 0xf9, 0x31, 0xd8, 0x42, 0x6e, 0xe1, 0x48, 0xc7, 0xd7,
	0x2b, 0x1d, 0xb7, 0x74, 0x7c, 0xbd, 0x6a, 0xe1, 0x96, 0x82, 0x6f, 0xe3, 0x8c, 0x30, 0x1d, 0x37,
	0x15, 0x5c, 0xc8, 0x2d, 0xdc, 0x94, 0xf8, 0x18, 0xec, 0x20, 0xbb, 0x9c, 0x62, 0xaa, 0xf3, 0x86,
	0xf2, 0xf6, 0xb5, 0xde, 0x1a, 0x60, 0xc8, 0x01, 0xaf, 0xa0, 0xb3, 0xc9, 0xb2, 0x58, 0xc7, 0x7b,
	0x12, 0x1f, 0x01, 0xe2, 0x6a, 0x0b, 0xee, 0x29, 0x0f, 0xbf, 0xa9, 0x58, 0x3d, 0xfb, 0x46, 0x3b,
	0xca, 0xe1, 0x42, 0x6e, 0x1f, 0xee, 0x48, 0xfe, 0x03, 0xb8, 0xfb, 0xb0, 0x76, 0xe7, 0x13, 0x6b,
	0x5c, 0x74, 0xc1, 0xda, 0xd3, 0xaa, 0xf1, 0x90, 0x87, 0xb7, 0xe1, 0xa6, 0xea, 0xeb, 0x7b, 0x78,
	0xb1, 0x0f, 0xc5, 0x6a, 0xef, 0xa1, 0xd6, 0x63, 0xe8, 0x47, 0x05, 0x6d, 0xcc, 0x79, 0x88, 0x0e,
	0xc1, 0x16, 0xdd, 0xa5, 0x67, 0x0a, 0xbf, 0x9a, 0xcc, 0xff, 0x02, 0x9d, 0x6d, 0x14, 0x53, 0x8c,
	0xa1, 0xf3, 0x8d, 0x24, 0xf2, 0xb2, 0x89, 0x98, 0xd7, 0x02, 0xc2, 0x88, 0x38, 0xcd, 0x39, 0x88,
	0x18, 0x7b, 0xd0, 0xe5, 0xff, 0x47, 0x56, 0x78, 0x96, 0x68, 0x95, 0xa9, 0x3f, 0x05, 0xc4, 0x27,
	0x95, 0xf8, 0x8d, 0xba, 0x9c, 0xfe, 0xa2, 0x3f, 0xab, 0xef, 0x3b, 0x17, 0x9b, 0x4d, 0x4d, 0x43,
	0xe8, 0x05, 0x84, 0xd1, 0x63, 0x4e, 0x52, 0xec, 0x82, 0x23, 0xe3, 0x5f, 0x01, 0xa9, 0xdc, 0x27,
	0xf8, 0x25, 0x3c, 0xbb, 0x56, 0xbe, 0x53, 0xf2, 0xcf, 0x35, 0x30, 0x86, 0xe7, 0xd7, 0xd2, 0xd7,
	0x2c, 0x65, 0x67, 0xd7, 0xd4, 0xda, 0x7e, 0x50, 0x52, 0xb8, 0xd6, 0x66, 0xf8, 0x73, 0x70, 0xef,
	0x8b, 0x3b, 0xd9, 0x22, 0x59, 0xfe, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x33, 0x8e, 0xfa, 0x93, 0x90,
	0x03, 0x00, 0x00,
}