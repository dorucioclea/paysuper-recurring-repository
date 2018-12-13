// Code generated by protoc-gen-go. DO NOT EDIT.
// source: payment_method.proto

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

type PaymentSystem struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Country              *Country             `protobuf:"bytes,3,opt,name=Country,proto3" json:"Country,omitempty"`
	AccountingCurrency   *Currency            `protobuf:"bytes,4,opt,name=accounting_currency,json=accountingCurrency,proto3" json:"accounting_currency,omitempty"`
	AccountingPeriod     string               `protobuf:"bytes,5,opt,name=accounting_period,json=accountingPeriod,proto3" json:"accounting_period,omitempty"`
	IsActive             bool                 `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PaymentSystem) Reset()         { *m = PaymentSystem{} }
func (m *PaymentSystem) String() string { return proto.CompactTextString(m) }
func (*PaymentSystem) ProtoMessage()    {}
func (*PaymentSystem) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_method_c3f32286f022385a, []int{0}
}
func (m *PaymentSystem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentSystem.Unmarshal(m, b)
}
func (m *PaymentSystem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentSystem.Marshal(b, m, deterministic)
}
func (dst *PaymentSystem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentSystem.Merge(dst, src)
}
func (m *PaymentSystem) XXX_Size() int {
	return xxx_messageInfo_PaymentSystem.Size(m)
}
func (m *PaymentSystem) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentSystem.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentSystem proto.InternalMessageInfo

func (m *PaymentSystem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentSystem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PaymentSystem) GetCountry() *Country {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *PaymentSystem) GetAccountingCurrency() *Currency {
	if m != nil {
		return m.AccountingCurrency
	}
	return nil
}

func (m *PaymentSystem) GetAccountingPeriod() string {
	if m != nil {
		return m.AccountingPeriod
	}
	return ""
}

func (m *PaymentSystem) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *PaymentSystem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PaymentSystem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type PaymentMethodParams struct {
	Handler              string            `protobuf:"bytes,1,opt,name=handler,proto3" json:"handler,omitempty"`
	Terminal             string            `protobuf:"bytes,2,opt,name=terminal,proto3" json:"terminal,omitempty"`
	ExternalId           string            `protobuf:"bytes,3,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Other                map[string]string `protobuf:"bytes,4,rep,name=other,proto3" json:"other,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PaymentMethodParams) Reset()         { *m = PaymentMethodParams{} }
func (m *PaymentMethodParams) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodParams) ProtoMessage()    {}
func (*PaymentMethodParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_method_c3f32286f022385a, []int{1}
}
func (m *PaymentMethodParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodParams.Unmarshal(m, b)
}
func (m *PaymentMethodParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodParams.Marshal(b, m, deterministic)
}
func (dst *PaymentMethodParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodParams.Merge(dst, src)
}
func (m *PaymentMethodParams) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodParams.Size(m)
}
func (m *PaymentMethodParams) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodParams.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodParams proto.InternalMessageInfo

func (m *PaymentMethodParams) GetHandler() string {
	if m != nil {
		return m.Handler
	}
	return ""
}

func (m *PaymentMethodParams) GetTerminal() string {
	if m != nil {
		return m.Terminal
	}
	return ""
}

func (m *PaymentMethodParams) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *PaymentMethodParams) GetOther() map[string]string {
	if m != nil {
		return m.Other
	}
	return nil
}

type PaymentMethodOrder struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Params               *PaymentMethodParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	PaymentSystem        *PaymentSystem       `protobuf:"bytes,4,opt,name=payment_system,json=paymentSystem,proto3" json:"payment_system,omitempty"`
	GroupAlias           string               `protobuf:"bytes,5,opt,name=group_alias,json=groupAlias,proto3" json:"group_alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PaymentMethodOrder) Reset()         { *m = PaymentMethodOrder{} }
func (m *PaymentMethodOrder) String() string { return proto.CompactTextString(m) }
func (*PaymentMethodOrder) ProtoMessage()    {}
func (*PaymentMethodOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_payment_method_c3f32286f022385a, []int{2}
}
func (m *PaymentMethodOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PaymentMethodOrder.Unmarshal(m, b)
}
func (m *PaymentMethodOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PaymentMethodOrder.Marshal(b, m, deterministic)
}
func (dst *PaymentMethodOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PaymentMethodOrder.Merge(dst, src)
}
func (m *PaymentMethodOrder) XXX_Size() int {
	return xxx_messageInfo_PaymentMethodOrder.Size(m)
}
func (m *PaymentMethodOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_PaymentMethodOrder.DiscardUnknown(m)
}

var xxx_messageInfo_PaymentMethodOrder proto.InternalMessageInfo

func (m *PaymentMethodOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PaymentMethodOrder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PaymentMethodOrder) GetParams() *PaymentMethodParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *PaymentMethodOrder) GetPaymentSystem() *PaymentSystem {
	if m != nil {
		return m.PaymentSystem
	}
	return nil
}

func (m *PaymentMethodOrder) GetGroupAlias() string {
	if m != nil {
		return m.GroupAlias
	}
	return ""
}

func init() {
	proto.RegisterType((*PaymentSystem)(nil), "processing.PaymentSystem")
	proto.RegisterType((*PaymentMethodParams)(nil), "processing.PaymentMethodParams")
	proto.RegisterMapType((map[string]string)(nil), "processing.PaymentMethodParams.OtherEntry")
	proto.RegisterType((*PaymentMethodOrder)(nil), "processing.PaymentMethodOrder")
}

func init() {
	proto.RegisterFile("payment_method.proto", fileDescriptor_payment_method_c3f32286f022385a)
}

var fileDescriptor_payment_method_c3f32286f022385a = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xdd, 0x7c, 0x4e, 0xd4, 0x2a, 0x4c, 0x72, 0x30, 0xe6, 0x90, 0x28, 0xa7, 0x08, 0x84,
	0x2b, 0x95, 0x03, 0x85, 0x53, 0x23, 0xd4, 0x03, 0x07, 0xd4, 0xc8, 0x70, 0xb7, 0xb6, 0xde, 0xc1,
	0x59, 0x61, 0xef, 0x5a, 0xeb, 0x75, 0x85, 0xff, 0x27, 0x3f, 0x02, 0xfe, 0x05, 0xf2, 0x7a, 0x4d,
	0x1a, 0xa9, 0x52, 0xb9, 0xed, 0xbc, 0x99, 0x37, 0x7e, 0x7a, 0xf3, 0x0c, 0xcb, 0x92, 0x35, 0x05,
	0x49, 0x93, 0x14, 0x64, 0x0e, 0x8a, 0x47, 0xa5, 0x56, 0x46, 0x21, 0x94, 0x5a, 0xa5, 0x54, 0x55,
	0x42, 0x66, 0xe1, 0x9c, 0x8b, 0xd4, 0x08, 0x25, 0x99, 0x6e, 0xba, 0x6e, 0xb8, 0xca, 0x94, 0xca,
	0x72, 0xba, 0xb4, 0xd5, 0x7d, 0xfd, 0xfd, 0xd2, 0x88, 0x82, 0x2a, 0xc3, 0x8a, 0xb2, 0x1b, 0xd8,
	0xfc, 0xf1, 0xe1, 0x7c, 0xdf, 0xed, 0xfd, 0xda, 0x54, 0x86, 0x0a, 0xbc, 0x00, 0x5f, 0xf0, 0xc0,
	0x5b, 0x7b, 0xdb, 0x69, 0xec, 0x0b, 0x8e, 0x08, 0x03, 0xc9, 0x0a, 0x0a, 0x7c, 0x8b, 0xd8, 0x37,
	0xbe, 0x85, 0xf1, 0x27, 0x55, 0x4b, 0xa3, 0x9b, 0xe0, 0x6c, 0xed, 0x6d, 0x67, 0x57, 0x8b, 0xe8,
	0x28, 0x23, 0x72, 0xad, 0xb8, 0x9f, 0xc1, 0x5b, 0x58, 0xb0, 0x34, 0x6d, 0x0b, 0x21, 0xb3, 0x24,
	0xad, 0xb5, 0x26, 0x99, 0x36, 0xc1, 0xc0, 0x52, 0x97, 0x27, 0x54, 0xd7, 0x8b, 0xf1, 0x48, 0xe8,
	0x31, 0x7c, 0x03, 0x2f, 0x1e, 0xad, 0x29, 0x49, 0x0b, 0xc5, 0x83, 0xa1, 0x95, 0x35, 0x3f, 0x36,
	0xf6, 0x16, 0xc7, 0x57, 0x30, 0x15, 0x55, 0xc2, 0x52, 0x23, 0x1e, 0x28, 0x18, 0xad, 0xbd, 0xed,
	0x24, 0x9e, 0x88, 0x6a, 0x67, 0x6b, 0xfc, 0x00, 0x90, 0x6a, 0x62, 0x86, 0x78, 0xc2, 0x4c, 0x30,
	0xb6, 0x3a, 0xc2, 0xa8, 0xf3, 0x2a, 0xea, 0xbd, 0x8a, 0xbe, 0xf5, 0x5e, 0xc5, 0x53, 0x37, 0xbd,
	0x33, 0x2d, 0xb5, 0x2e, 0x79, 0x4f, 0x9d, 0x3c, 0x4f, 0x75, 0xd3, 0x3b, 0xb3, 0xf9, 0xed, 0xc1,
	0xc2, 0x79, 0xfd, 0xc5, 0x9e, 0x70, 0xcf, 0x34, 0x2b, 0x2a, 0x0c, 0x60, 0x7c, 0x60, 0x92, 0xe7,
	0xa4, 0x9d, 0xed, 0x7d, 0x89, 0x21, 0x4c, 0x0c, 0xe9, 0x42, 0x48, 0x96, 0x3b, 0xff, 0xff, 0xd5,
	0xb8, 0x82, 0x19, 0xfd, 0x34, 0xa4, 0x25, 0xcb, 0x13, 0xc1, 0xed, 0x1d, 0xa6, 0x31, 0xf4, 0xd0,
	0x67, 0x8e, 0x37, 0x30, 0x54, 0xe6, 0x40, 0x3a, 0x18, 0xac, 0xcf, 0xb6, 0xb3, 0xab, 0xd7, 0x8f,
	0x7d, 0x7e, 0x42, 0x46, 0x74, 0xd7, 0x0e, 0xdf, 0xda, 0xcb, 0x75, 0xc4, 0xf0, 0x1a, 0xe0, 0x08,
	0xe2, 0x1c, 0xce, 0x7e, 0x50, 0xe3, 0x24, 0xb6, 0x4f, 0x5c, 0xc2, 0xf0, 0x81, 0xe5, 0x75, 0x9f,
	0x8d, 0xae, 0xf8, 0xe8, 0x5f, 0x7b, 0x9b, 0x5f, 0x1e, 0xe0, 0xc9, 0x37, 0xee, 0x34, 0x27, 0xfd,
	0x5f, 0xd9, 0x7a, 0x0f, 0xa3, 0xd2, 0x0a, 0x72, 0xd1, 0x5a, 0x3d, 0xa3, 0x3b, 0x76, 0xe3, 0x78,
	0x03, 0x17, 0xfd, 0x1f, 0x52, 0xd9, 0x28, 0xbb, 0x80, 0xbd, 0x7c, 0x62, 0x41, 0x97, 0xf5, 0xf8,
	0xbc, 0x3c, 0x89, 0xfe, 0x0a, 0x66, 0x99, 0x56, 0x75, 0x99, 0xb0, 0x5c, 0xb0, 0xca, 0x45, 0x0b,
	0x2c, 0xb4, 0x6b, 0x91, 0xfb, 0x91, 0x3d, 0xf0, 0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x07,
	0xaa, 0xf5, 0xda, 0x8b, 0x03, 0x00, 0x00,
}