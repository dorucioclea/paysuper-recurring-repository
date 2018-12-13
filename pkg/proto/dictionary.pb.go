// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dictionary.proto

package processing

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Name struct {
	En                   string   `protobuf:"bytes,1,opt,name=en,proto3" json:"en,omitempty"`
	Ru                   string   `protobuf:"bytes,2,opt,name=ru,proto3" json:"ru,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Name) Reset()         { *m = Name{} }
func (m *Name) String() string { return proto.CompactTextString(m) }
func (*Name) ProtoMessage()    {}
func (*Name) Descriptor() ([]byte, []int) {
	return fileDescriptor_dictionary_d208ce3e6f18f83b, []int{0}
}
func (m *Name) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Name.Unmarshal(m, b)
}
func (m *Name) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Name.Marshal(b, m, deterministic)
}
func (dst *Name) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Name.Merge(dst, src)
}
func (m *Name) XXX_Size() int {
	return xxx_messageInfo_Name.Size(m)
}
func (m *Name) XXX_DiscardUnknown() {
	xxx_messageInfo_Name.DiscardUnknown(m)
}

var xxx_messageInfo_Name proto.InternalMessageInfo

func (m *Name) GetEn() string {
	if m != nil {
		return m.En
	}
	return ""
}

func (m *Name) GetRu() string {
	if m != nil {
		return m.Ru
	}
	return ""
}

type Country struct {
	CodeInt              int32                `protobuf:"varint,1,opt,name=code_int,json=codeInt,proto3" json:"code_int,omitempty"`
	CodeA2               string               `protobuf:"bytes,2,opt,name=code_a2,json=codeA2,proto3" json:"code_a2,omitempty"`
	CodeA3               string               `protobuf:"bytes,3,opt,name=code_a3,json=codeA3,proto3" json:"code_a3,omitempty"`
	Name                 *Name                `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	IsActive             bool                 `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Country) Reset()         { *m = Country{} }
func (m *Country) String() string { return proto.CompactTextString(m) }
func (*Country) ProtoMessage()    {}
func (*Country) Descriptor() ([]byte, []int) {
	return fileDescriptor_dictionary_d208ce3e6f18f83b, []int{1}
}
func (m *Country) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Country.Unmarshal(m, b)
}
func (m *Country) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Country.Marshal(b, m, deterministic)
}
func (dst *Country) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Country.Merge(dst, src)
}
func (m *Country) XXX_Size() int {
	return xxx_messageInfo_Country.Size(m)
}
func (m *Country) XXX_DiscardUnknown() {
	xxx_messageInfo_Country.DiscardUnknown(m)
}

var xxx_messageInfo_Country proto.InternalMessageInfo

func (m *Country) GetCodeInt() int32 {
	if m != nil {
		return m.CodeInt
	}
	return 0
}

func (m *Country) GetCodeA2() string {
	if m != nil {
		return m.CodeA2
	}
	return ""
}

func (m *Country) GetCodeA3() string {
	if m != nil {
		return m.CodeA3
	}
	return ""
}

func (m *Country) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Country) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Country) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Country) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type Currency struct {
	CodeInt              int32                `protobuf:"varint,1,opt,name=code_int,json=codeInt,proto3" json:"code_int,omitempty"`
	CodeA3               string               `protobuf:"bytes,2,opt,name=code_a3,json=codeA3,proto3" json:"code_a3,omitempty"`
	Name                 *Name                `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	IsActive             bool                 `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Currency) Reset()         { *m = Currency{} }
func (m *Currency) String() string { return proto.CompactTextString(m) }
func (*Currency) ProtoMessage()    {}
func (*Currency) Descriptor() ([]byte, []int) {
	return fileDescriptor_dictionary_d208ce3e6f18f83b, []int{2}
}
func (m *Currency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Currency.Unmarshal(m, b)
}
func (m *Currency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Currency.Marshal(b, m, deterministic)
}
func (dst *Currency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Currency.Merge(dst, src)
}
func (m *Currency) XXX_Size() int {
	return xxx_messageInfo_Currency.Size(m)
}
func (m *Currency) XXX_DiscardUnknown() {
	xxx_messageInfo_Currency.DiscardUnknown(m)
}

var xxx_messageInfo_Currency proto.InternalMessageInfo

func (m *Currency) GetCodeInt() int32 {
	if m != nil {
		return m.CodeInt
	}
	return 0
}

func (m *Currency) GetCodeA3() string {
	if m != nil {
		return m.CodeA3
	}
	return ""
}

func (m *Currency) GetName() *Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *Currency) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Currency) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Currency) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Name)(nil), "processing.Name")
	proto.RegisterType((*Country)(nil), "processing.Country")
	proto.RegisterType((*Currency)(nil), "processing.Currency")
}

func init() { proto.RegisterFile("dictionary.proto", fileDescriptor_dictionary_d208ce3e6f18f83b) }

var fileDescriptor_dictionary_d208ce3e6f18f83b = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x90, 0x3d, 0x4f, 0xc3, 0x30,
	0x10, 0x86, 0x95, 0xd0, 0x8f, 0xf4, 0x90, 0x50, 0xe5, 0x05, 0x53, 0x06, 0xaa, 0x0a, 0xa1, 0x4e,
	0xae, 0xd4, 0x4c, 0x8c, 0x51, 0x27, 0x16, 0x86, 0x88, 0x3d, 0x72, 0x9d, 0x23, 0xb2, 0x44, 0xec,
	0xc8, 0x39, 0x23, 0xf5, 0x57, 0xf0, 0x73, 0x59, 0x51, 0x9c, 0x44, 0x85, 0x09, 0x56, 0x36, 0xbf,
	0x7a, 0xee, 0xb1, 0xee, 0x3d, 0x58, 0x96, 0x5a, 0x91, 0xb6, 0x46, 0xba, 0x93, 0x68, 0x9c, 0x25,
	0xcb, 0xa0, 0x71, 0x56, 0x61, 0xdb, 0x6a, 0x53, 0xad, 0xee, 0x2a, 0x6b, 0xab, 0x37, 0xdc, 0x05,
	0x72, 0xf4, 0xaf, 0x3b, 0xd2, 0x35, 0xb6, 0x24, 0xeb, 0xa6, 0x1f, 0xde, 0x3c, 0xc0, 0xe4, 0x59,
	0xd6, 0xc8, 0xae, 0x20, 0x46, 0xc3, 0xa3, 0x75, 0xb4, 0x5d, 0xe4, 0x31, 0x9a, 0x2e, 0x3b, 0xcf,
	0xe3, 0x3e, 0x3b, 0xbf, 0xf9, 0x88, 0x61, 0x7e, 0xb0, 0xde, 0x90, 0x3b, 0xb1, 0x1b, 0x48, 0x94,
	0x2d, 0xb1, 0xd0, 0x86, 0x82, 0x31, 0xcd, 0xe7, 0x5d, 0x7e, 0x32, 0xc4, 0xae, 0x21, 0x3c, 0x0b,
	0xb9, 0x1f, 0xdc, 0x59, 0x17, 0xb3, 0xfd, 0x19, 0xa4, 0xfc, 0xe2, 0x1b, 0x48, 0xd9, 0x3d, 0x4c,
	0x8c, 0xac, 0x91, 0x4f, 0xd6, 0xd1, 0xf6, 0x72, 0xbf, 0x14, 0xe7, 0xe5, 0x45, 0xb7, 0x58, 0x1e,
	0x28, 0xbb, 0x85, 0x85, 0x6e, 0x0b, 0xa9, 0x48, 0xbf, 0x23, 0x9f, 0xae, 0xa3, 0x6d, 0x92, 0x27,
	0xba, 0xcd, 0x42, 0x66, 0x8f, 0x00, 0xca, 0xa1, 0x24, 0x2c, 0x0b, 0x49, 0x7c, 0x16, 0x3e, 0x5a,
	0x89, 0xbe, 0xb9, 0x18, 0x9b, 0x8b, 0x97, 0xb1, 0x79, 0xbe, 0x18, 0xa6, 0x33, 0xea, 0x54, 0xdf,
	0x94, 0xa3, 0x3a, 0xff, 0x5d, 0x1d, 0xa6, 0x33, 0xda, 0x7c, 0x46, 0x90, 0x1c, 0xbc, 0x73, 0x68,
	0xd4, 0xdf, 0x4e, 0x92, 0xfe, 0x38, 0xc9, 0xff, 0x6d, 0x7e, 0x9c, 0x05, 0x9c, 0x7e, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x0e, 0xce, 0x7c, 0x9f, 0x7b, 0x02, 0x00, 0x00,
}
