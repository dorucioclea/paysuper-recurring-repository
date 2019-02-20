// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity/entity.proto

package entity // import "github.com/ProtocolONE/payone-repository/pkg/proto/entity"

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

type SavedCard struct {
	// @inject_tag: bson:"_id"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	// @inject_tag: bson:"account"
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty" bson:"account"`
	// @inject_tag: bson:"project_id"
	ProjectId string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" bson:"project_id"`
	// @inject_tag: bson:"masked_pan"
	MaskedPan string `protobuf:"bytes,4,opt,name=masked_pan,json=maskedPan,proto3" json:"masked_pan,omitempty" bson:"masked_pan"`
	// @inject_tag: bson:"expire"
	Expire *CardExpire `protobuf:"bytes,5,opt,name=expire,proto3" json:"expire,omitempty" bson:"expire"`
	// @inject_tag: bson:"recurring_id"
	RecurringId string `protobuf:"bytes,6,opt,name=recurring_id,json=recurringId,proto3" json:"recurring_id,omitempty" bson:"recurring_id"`
	// @inject_tag: bson:"is_active"
	IsActive bool `protobuf:"varint,7,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" bson:"is_active"`
	// @inject_tag: bson:"created_at"
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" bson:"created_at"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte               `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32                `json:"-" bson:"-" structure:"-"`
}

func (m *SavedCard) Reset()         { *m = SavedCard{} }
func (m *SavedCard) String() string { return proto.CompactTextString(m) }
func (*SavedCard) ProtoMessage()    {}
func (*SavedCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_2ad515b891205a20, []int{0}
}
func (m *SavedCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SavedCard.Unmarshal(m, b)
}
func (m *SavedCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SavedCard.Marshal(b, m, deterministic)
}
func (dst *SavedCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SavedCard.Merge(dst, src)
}
func (m *SavedCard) XXX_Size() int {
	return xxx_messageInfo_SavedCard.Size(m)
}
func (m *SavedCard) XXX_DiscardUnknown() {
	xxx_messageInfo_SavedCard.DiscardUnknown(m)
}

var xxx_messageInfo_SavedCard proto.InternalMessageInfo

func (m *SavedCard) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SavedCard) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *SavedCard) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *SavedCard) GetMaskedPan() string {
	if m != nil {
		return m.MaskedPan
	}
	return ""
}

func (m *SavedCard) GetExpire() *CardExpire {
	if m != nil {
		return m.Expire
	}
	return nil
}

func (m *SavedCard) GetRecurringId() string {
	if m != nil {
		return m.RecurringId
	}
	return ""
}

func (m *SavedCard) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *SavedCard) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type CardExpire struct {
	// @inject_tag: bson:"month"
	Month string `protobuf:"bytes,1,opt,name=month,proto3" json:"month,omitempty" bson:"month"`
	// @inject_tag: bson:"year"
	Year                 string   `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty" bson:"year"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-" structure:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-" structure:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-" structure:"-"`
}

func (m *CardExpire) Reset()         { *m = CardExpire{} }
func (m *CardExpire) String() string { return proto.CompactTextString(m) }
func (*CardExpire) ProtoMessage()    {}
func (*CardExpire) Descriptor() ([]byte, []int) {
	return fileDescriptor_entity_2ad515b891205a20, []int{1}
}
func (m *CardExpire) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CardExpire.Unmarshal(m, b)
}
func (m *CardExpire) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CardExpire.Marshal(b, m, deterministic)
}
func (dst *CardExpire) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CardExpire.Merge(dst, src)
}
func (m *CardExpire) XXX_Size() int {
	return xxx_messageInfo_CardExpire.Size(m)
}
func (m *CardExpire) XXX_DiscardUnknown() {
	xxx_messageInfo_CardExpire.DiscardUnknown(m)
}

var xxx_messageInfo_CardExpire proto.InternalMessageInfo

func (m *CardExpire) GetMonth() string {
	if m != nil {
		return m.Month
	}
	return ""
}

func (m *CardExpire) GetYear() string {
	if m != nil {
		return m.Year
	}
	return ""
}

func init() {
	proto.RegisterType((*SavedCard)(nil), "entity.SavedCard")
	proto.RegisterType((*CardExpire)(nil), "entity.CardExpire")
}

func init() { proto.RegisterFile("entity/entity.proto", fileDescriptor_entity_2ad515b891205a20) }

var fileDescriptor_entity_2ad515b891205a20 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0xcf, 0x6b, 0xdb, 0x30,
	0x18, 0xc5, 0x5e, 0xe2, 0xc4, 0x5f, 0xc6, 0x0e, 0xda, 0x0e, 0x22, 0x63, 0x2c, 0xcb, 0x29, 0x0c,
	0x6a, 0x43, 0x0b, 0x85, 0xd0, 0x53, 0x5a, 0x72, 0xc8, 0xa5, 0x0d, 0x6e, 0x4f, 0xbd, 0x18, 0x45,
	0xfa, 0xea, 0xa8, 0x89, 0x2d, 0x21, 0xcb, 0xa1, 0xfe, 0x57, 0xfa, 0xd7, 0x16, 0x4b, 0x4e, 0x7b,
	0xb2, 0xdf, 0x0f, 0xf4, 0x1e, 0xef, 0x83, 0x9f, 0x58, 0x59, 0x69, 0xdb, 0xd4, 0x7f, 0x12, 0x6d,
	0x94, 0x55, 0x24, 0xf2, 0x68, 0xfa, 0xb7, 0x50, 0xaa, 0x38, 0x62, 0xea, 0xd8, 0x5d, 0xf3, 0x92,
	0x5a, 0x59, 0x62, 0x6d, 0x59, 0xa9, 0xbd, 0x71, 0xfe, 0x1e, 0x42, 0xfc, 0xc8, 0x4e, 0x28, 0xee,
	0x98, 0x11, 0xe4, 0x07, 0x84, 0x52, 0xd0, 0x60, 0x16, 0x2c, 0xe2, 0x2c, 0x94, 0x82, 0x50, 0x18,
	0x31, 0xce, 0x55, 0x53, 0x59, 0x1a, 0x3a, 0xf2, 0x0c, 0xc9, 0x1f, 0x00, 0x6d, 0xd4, 0x2b, 0x72,
	0x9b, 0x4b, 0x41, 0xbf, 0x39, 0x31, 0xee, 0x99, 0x8d, 0xe8, 0xe4, 0x92, 0xd5, 0x07, 0x14, 0xb9,
	0x66, 0x15, 0x1d, 0x78, 0xd9, 0x33, 0x5b, 0x56, 0x91, 0xff, 0x10, 0xe1, 0x9b, 0x96, 0x06, 0xe9,
	0x70, 0x16, 0x2c, 0x26, 0x97, 0x24, 0xe9, 0xdb, 0x77, 0x2d, 0xd6, 0x4e, 0xc9, 0x7a, 0x07, 0xf9,
	0x07, 0xdf, 0x0d, 0xf2, 0xc6, 0x18, 0x59, 0x15, 0x5d, 0x56, 0xe4, 0x1e, 0x9b, 0x7c, 0x72, 0x1b,
	0x41, 0x7e, 0x43, 0x2c, 0xeb, 0x9c, 0x71, 0x2b, 0x4f, 0x48, 0x47, 0xb3, 0x60, 0x31, 0xce, 0xc6,
	0xb2, 0x5e, 0x39, 0x4c, 0x96, 0x00, 0xdc, 0x20, 0xb3, 0x28, 0x72, 0x66, 0xe9, 0xd8, 0xe5, 0x4d,
	0x13, 0xbf, 0x4b, 0x72, 0xde, 0x25, 0x79, 0x3a, 0xef, 0x92, 0xc5, 0xbd, 0x7b, 0x65, 0xe7, 0xd7,
	0x00, 0x5f, 0x85, 0xc8, 0x2f, 0x18, 0x96, 0xaa, 0xb2, 0xfb, 0x7e, 0x1f, 0x0f, 0x08, 0x81, 0x41,
	0x8b, 0xcc, 0xf4, 0xfb, 0xb8, 0xff, 0xdb, 0x9b, 0xe7, 0x65, 0x21, 0xed, 0xbe, 0xd9, 0x25, 0x5c,
	0x95, 0xe9, 0xb6, 0xcb, 0xe0, 0xea, 0xf8, 0x70, 0xbf, 0x4e, 0x35, 0x6b, 0x55, 0x85, 0x17, 0x06,
	0xb5, 0xaa, 0xa5, 0x55, 0xa6, 0x4d, 0xf5, 0xa1, 0xf0, 0xd7, 0xe9, 0x0f, 0xb8, 0x8b, 0x1c, 0xba,
	0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0x40, 0x84, 0xab, 0x06, 0xd8, 0x01, 0x00, 0x00,
}
