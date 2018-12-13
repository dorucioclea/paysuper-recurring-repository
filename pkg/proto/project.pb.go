// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

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

type AccountingPeriod int32

const (
	AccountingPeriod_day       AccountingPeriod = 0
	AccountingPeriod_week      AccountingPeriod = 1
	AccountingPeriod_two_week  AccountingPeriod = 2
	AccountingPeriod_month     AccountingPeriod = 3
	AccountingPeriod_quarter   AccountingPeriod = 4
	AccountingPeriod_half_year AccountingPeriod = 5
	AccountingPeriod_year      AccountingPeriod = 6
)

var AccountingPeriod_name = map[int32]string{
	0: "day",
	1: "week",
	2: "two_week",
	3: "month",
	4: "quarter",
	5: "half_year",
	6: "year",
}
var AccountingPeriod_value = map[string]int32{
	"day":       0,
	"week":      1,
	"two_week":  2,
	"month":     3,
	"quarter":   4,
	"half_year": 5,
	"year":      6,
}

func (x AccountingPeriod) String() string {
	return proto.EnumName(AccountingPeriod_name, int32(x))
}
func (AccountingPeriod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{0}
}

type Merchant struct {
	Id                        string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExternalId                string               `protobuf:"bytes,2,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Email                     string               `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name                      string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Country                   *Country             `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	AccountingPeriod          AccountingPeriod     `protobuf:"varint,6,opt,name=accounting_period,json=accountingPeriod,proto3,enum=processing.AccountingPeriod" json:"accounting_period,omitempty"`
	Currency                  *Currency            `protobuf:"bytes,7,opt,name=currency,proto3" json:"currency,omitempty"`
	IsVatEnabled              bool                 `protobuf:"varint,8,opt,name=is_vat_enabled,json=isVatEnabled,proto3" json:"is_vat_enabled,omitempty"`
	IsCommissionToUserEnabled bool                 `protobuf:"varint,9,opt,name=is_commission_to_user_enabled,json=isCommissionToUserEnabled,proto3" json:"is_commission_to_user_enabled,omitempty"`
	Status                    int32                `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt                 *timestamp.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                 *timestamp.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}             `json:"-"`
	XXX_unrecognized          []byte               `json:"-"`
	XXX_sizecache             int32                `json:"-"`
}

func (m *Merchant) Reset()         { *m = Merchant{} }
func (m *Merchant) String() string { return proto.CompactTextString(m) }
func (*Merchant) ProtoMessage()    {}
func (*Merchant) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{0}
}
func (m *Merchant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Merchant.Unmarshal(m, b)
}
func (m *Merchant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Merchant.Marshal(b, m, deterministic)
}
func (dst *Merchant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Merchant.Merge(dst, src)
}
func (m *Merchant) XXX_Size() int {
	return xxx_messageInfo_Merchant.Size(m)
}
func (m *Merchant) XXX_DiscardUnknown() {
	xxx_messageInfo_Merchant.DiscardUnknown(m)
}

var xxx_messageInfo_Merchant proto.InternalMessageInfo

func (m *Merchant) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Merchant) GetExternalId() string {
	if m != nil {
		return m.ExternalId
	}
	return ""
}

func (m *Merchant) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Merchant) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Merchant) GetCountry() *Country {
	if m != nil {
		return m.Country
	}
	return nil
}

func (m *Merchant) GetAccountingPeriod() AccountingPeriod {
	if m != nil {
		return m.AccountingPeriod
	}
	return AccountingPeriod_day
}

func (m *Merchant) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

func (m *Merchant) GetIsVatEnabled() bool {
	if m != nil {
		return m.IsVatEnabled
	}
	return false
}

func (m *Merchant) GetIsCommissionToUserEnabled() bool {
	if m != nil {
		return m.IsCommissionToUserEnabled
	}
	return false
}

func (m *Merchant) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Merchant) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Merchant) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type FixedPackage struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CurrencyInt          int32                `protobuf:"varint,3,opt,name=currency_int,json=currencyInt,proto3" json:"currency_int,omitempty"`
	Price                float64              `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	IsActive             bool                 `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Currency             *Currency            `protobuf:"bytes,8,opt,name=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FixedPackage) Reset()         { *m = FixedPackage{} }
func (m *FixedPackage) String() string { return proto.CompactTextString(m) }
func (*FixedPackage) ProtoMessage()    {}
func (*FixedPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{1}
}
func (m *FixedPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedPackage.Unmarshal(m, b)
}
func (m *FixedPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedPackage.Marshal(b, m, deterministic)
}
func (dst *FixedPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedPackage.Merge(dst, src)
}
func (m *FixedPackage) XXX_Size() int {
	return xxx_messageInfo_FixedPackage.Size(m)
}
func (m *FixedPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedPackage.DiscardUnknown(m)
}

var xxx_messageInfo_FixedPackage proto.InternalMessageInfo

func (m *FixedPackage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *FixedPackage) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *FixedPackage) GetCurrencyInt() int32 {
	if m != nil {
		return m.CurrencyInt
	}
	return 0
}

func (m *FixedPackage) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *FixedPackage) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *FixedPackage) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *FixedPackage) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *FixedPackage) GetCurrency() *Currency {
	if m != nil {
		return m.Currency
	}
	return nil
}

type FixedPackages struct {
	FixedPackage         []*FixedPackage `protobuf:"bytes,1,rep,name=fixed_package,json=fixedPackage,proto3" json:"fixed_package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *FixedPackages) Reset()         { *m = FixedPackages{} }
func (m *FixedPackages) String() string { return proto.CompactTextString(m) }
func (*FixedPackages) ProtoMessage()    {}
func (*FixedPackages) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{2}
}
func (m *FixedPackages) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FixedPackages.Unmarshal(m, b)
}
func (m *FixedPackages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FixedPackages.Marshal(b, m, deterministic)
}
func (dst *FixedPackages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FixedPackages.Merge(dst, src)
}
func (m *FixedPackages) XXX_Size() int {
	return xxx_messageInfo_FixedPackages.Size(m)
}
func (m *FixedPackages) XXX_DiscardUnknown() {
	xxx_messageInfo_FixedPackages.DiscardUnknown(m)
}

var xxx_messageInfo_FixedPackages proto.InternalMessageInfo

func (m *FixedPackages) GetFixedPackage() []*FixedPackage {
	if m != nil {
		return m.FixedPackage
	}
	return nil
}

type ProjectPaymentMethod struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AddedAAt             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=added_aAt,json=addedAAt,proto3" json:"added_aAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ProjectPaymentMethod) Reset()         { *m = ProjectPaymentMethod{} }
func (m *ProjectPaymentMethod) String() string { return proto.CompactTextString(m) }
func (*ProjectPaymentMethod) ProtoMessage()    {}
func (*ProjectPaymentMethod) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{3}
}
func (m *ProjectPaymentMethod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectPaymentMethod.Unmarshal(m, b)
}
func (m *ProjectPaymentMethod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectPaymentMethod.Marshal(b, m, deterministic)
}
func (dst *ProjectPaymentMethod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectPaymentMethod.Merge(dst, src)
}
func (m *ProjectPaymentMethod) XXX_Size() int {
	return xxx_messageInfo_ProjectPaymentMethod.Size(m)
}
func (m *ProjectPaymentMethod) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectPaymentMethod.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectPaymentMethod proto.InternalMessageInfo

func (m *ProjectPaymentMethod) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProjectPaymentMethod) GetAddedAAt() *timestamp.Timestamp {
	if m != nil {
		return m.AddedAAt
	}
	return nil
}

type ProjectPaymentMethods struct {
	PaymentMethods       []*ProjectPaymentMethod `protobuf:"bytes,1,rep,name=payment_methods,json=paymentMethods,proto3" json:"payment_methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ProjectPaymentMethods) Reset()         { *m = ProjectPaymentMethods{} }
func (m *ProjectPaymentMethods) String() string { return proto.CompactTextString(m) }
func (*ProjectPaymentMethods) ProtoMessage()    {}
func (*ProjectPaymentMethods) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{4}
}
func (m *ProjectPaymentMethods) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectPaymentMethods.Unmarshal(m, b)
}
func (m *ProjectPaymentMethods) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectPaymentMethods.Marshal(b, m, deterministic)
}
func (dst *ProjectPaymentMethods) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectPaymentMethods.Merge(dst, src)
}
func (m *ProjectPaymentMethods) XXX_Size() int {
	return xxx_messageInfo_ProjectPaymentMethods.Size(m)
}
func (m *ProjectPaymentMethods) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectPaymentMethods.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectPaymentMethods proto.InternalMessageInfo

func (m *ProjectPaymentMethods) GetPaymentMethods() []*ProjectPaymentMethod {
	if m != nil {
		return m.PaymentMethods
	}
	return nil
}

type Project struct {
	Id                         string                            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CallbackCurrency           *Currency                         `protobuf:"bytes,2,opt,name=callback_currency,json=callbackCurrency,proto3" json:"callback_currency,omitempty"`
	CallbackProtocol           string                            `protobuf:"bytes,3,opt,name=callback_protocol,json=callbackProtocol,proto3" json:"callback_protocol,omitempty"`
	CreateInvoiceAllowedUrls   []string                          `protobuf:"bytes,4,rep,name=create_invoice_allowed_urls,json=createInvoiceAllowedUrls,proto3" json:"create_invoice_allowed_urls,omitempty"`
	Merchant                   *Merchant                         `protobuf:"bytes,5,opt,name=merchant,proto3" json:"merchant,omitempty"`
	IsAllowDynamicNotifyUrls   bool                              `protobuf:"varint,6,opt,name=is_allow_dynamic_notify_urls,json=isAllowDynamicNotifyUrls,proto3" json:"is_allow_dynamic_notify_urls,omitempty"`
	IsAllowDynamicRedirectUrls bool                              `protobuf:"varint,7,opt,name=is_allow_dynamic_redirect_urls,json=isAllowDynamicRedirectUrls,proto3" json:"is_allow_dynamic_redirect_urls,omitempty"`
	LimitsCurrency             *Currency                         `protobuf:"bytes,8,opt,name=limits_currency,json=limitsCurrency,proto3" json:"limits_currency,omitempty"`
	MaxPaymentAmount           float64                           `protobuf:"fixed64,9,opt,name=MaxPaymentAmount,proto3" json:"MaxPaymentAmount,omitempty"`
	MinPaymentAmount           float64                           `protobuf:"fixed64,10,opt,name=MinPaymentAmount,proto3" json:"MinPaymentAmount,omitempty"`
	Name                       string                            `protobuf:"bytes,11,opt,name=Name,proto3" json:"Name,omitempty"`
	NotifyEmails               []string                          `protobuf:"bytes,12,rep,name=NotifyEmails,proto3" json:"NotifyEmails,omitempty"`
	OnlyFixedAmounts           bool                              `protobuf:"varint,13,opt,name=OnlyFixedAmounts,proto3" json:"OnlyFixedAmounts,omitempty"`
	SecretKey                  string                            `protobuf:"bytes,14,opt,name=SecretKey,proto3" json:"SecretKey,omitempty"`
	SendNotifyEmail            bool                              `protobuf:"varint,15,opt,name=SendNotifyEmail,proto3" json:"SendNotifyEmail,omitempty"`
	URLCheckAccount            string                            `protobuf:"bytes,16,opt,name=URLCheckAccount,proto3" json:"URLCheckAccount,omitempty"`
	URLProcessPayment          string                            `protobuf:"bytes,17,opt,name=URLProcessPayment,proto3" json:"URLProcessPayment,omitempty"`
	URLRedirectFail            string                            `protobuf:"bytes,18,opt,name=URLRedirectFail,proto3" json:"URLRedirectFail,omitempty"`
	URLRedirectSuccess         string                            `protobuf:"bytes,19,opt,name=URLRedirectSuccess,proto3" json:"URLRedirectSuccess,omitempty"`
	FixedPackage               map[string]*FixedPackages         `protobuf:"bytes,23,rep,name=fixed_package,json=fixedPackage,proto3" json:"fixed_package,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PaymentMethods             map[string]*ProjectPaymentMethods `protobuf:"bytes,24,rep,name=payment_methods,json=paymentMethods,proto3" json:"payment_methods,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IsActive                   bool                              `protobuf:"varint,20,opt,name=IsActive,proto3" json:"IsActive,omitempty"`
	CreatedAt                  *timestamp.Timestamp              `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt                  *timestamp.Timestamp              `protobuf:"bytes,22,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                          `json:"-"`
	XXX_unrecognized           []byte                            `json:"-"`
	XXX_sizecache              int32                             `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{5}
}
func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (dst *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(dst, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Project) GetCallbackCurrency() *Currency {
	if m != nil {
		return m.CallbackCurrency
	}
	return nil
}

func (m *Project) GetCallbackProtocol() string {
	if m != nil {
		return m.CallbackProtocol
	}
	return ""
}

func (m *Project) GetCreateInvoiceAllowedUrls() []string {
	if m != nil {
		return m.CreateInvoiceAllowedUrls
	}
	return nil
}

func (m *Project) GetMerchant() *Merchant {
	if m != nil {
		return m.Merchant
	}
	return nil
}

func (m *Project) GetIsAllowDynamicNotifyUrls() bool {
	if m != nil {
		return m.IsAllowDynamicNotifyUrls
	}
	return false
}

func (m *Project) GetIsAllowDynamicRedirectUrls() bool {
	if m != nil {
		return m.IsAllowDynamicRedirectUrls
	}
	return false
}

func (m *Project) GetLimitsCurrency() *Currency {
	if m != nil {
		return m.LimitsCurrency
	}
	return nil
}

func (m *Project) GetMaxPaymentAmount() float64 {
	if m != nil {
		return m.MaxPaymentAmount
	}
	return 0
}

func (m *Project) GetMinPaymentAmount() float64 {
	if m != nil {
		return m.MinPaymentAmount
	}
	return 0
}

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetNotifyEmails() []string {
	if m != nil {
		return m.NotifyEmails
	}
	return nil
}

func (m *Project) GetOnlyFixedAmounts() bool {
	if m != nil {
		return m.OnlyFixedAmounts
	}
	return false
}

func (m *Project) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *Project) GetSendNotifyEmail() bool {
	if m != nil {
		return m.SendNotifyEmail
	}
	return false
}

func (m *Project) GetURLCheckAccount() string {
	if m != nil {
		return m.URLCheckAccount
	}
	return ""
}

func (m *Project) GetURLProcessPayment() string {
	if m != nil {
		return m.URLProcessPayment
	}
	return ""
}

func (m *Project) GetURLRedirectFail() string {
	if m != nil {
		return m.URLRedirectFail
	}
	return ""
}

func (m *Project) GetURLRedirectSuccess() string {
	if m != nil {
		return m.URLRedirectSuccess
	}
	return ""
}

func (m *Project) GetFixedPackage() map[string]*FixedPackages {
	if m != nil {
		return m.FixedPackage
	}
	return nil
}

func (m *Project) GetPaymentMethods() map[string]*ProjectPaymentMethods {
	if m != nil {
		return m.PaymentMethods
	}
	return nil
}

func (m *Project) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Project) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Project) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type ProjectOrder struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	UrlSuccess           string    `protobuf:"bytes,3,opt,name=url_success,json=urlSuccess,proto3" json:"url_success,omitempty"`
	UrlFail              string    `protobuf:"bytes,4,opt,name=url_fail,json=urlFail,proto3" json:"url_fail,omitempty"`
	NotifyEmails         []string  `protobuf:"bytes,5,rep,name=notify_emails,json=notifyEmails,proto3" json:"notify_emails,omitempty"`
	SecretKey            string    `protobuf:"bytes,6,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	SendNotifyEmail      bool      `protobuf:"varint,7,opt,name=send_notify_email,json=sendNotifyEmail,proto3" json:"send_notify_email,omitempty"`
	UrlCheckAccount      string    `protobuf:"bytes,8,opt,name=url_check_account,json=urlCheckAccount,proto3" json:"url_check_account,omitempty"`
	UrlProcessPayment    string    `protobuf:"bytes,9,opt,name=url_process_payment,json=urlProcessPayment,proto3" json:"url_process_payment,omitempty"`
	CallbackProtocol     string    `protobuf:"bytes,10,opt,name=callback_protocol,json=callbackProtocol,proto3" json:"callback_protocol,omitempty"`
	Merchant             *Merchant `protobuf:"bytes,11,opt,name=merchant,proto3" json:"merchant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ProjectOrder) Reset()         { *m = ProjectOrder{} }
func (m *ProjectOrder) String() string { return proto.CompactTextString(m) }
func (*ProjectOrder) ProtoMessage()    {}
func (*ProjectOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_project_daff29af18dd3507, []int{6}
}
func (m *ProjectOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectOrder.Unmarshal(m, b)
}
func (m *ProjectOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectOrder.Marshal(b, m, deterministic)
}
func (dst *ProjectOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectOrder.Merge(dst, src)
}
func (m *ProjectOrder) XXX_Size() int {
	return xxx_messageInfo_ProjectOrder.Size(m)
}
func (m *ProjectOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectOrder.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectOrder proto.InternalMessageInfo

func (m *ProjectOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProjectOrder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProjectOrder) GetUrlSuccess() string {
	if m != nil {
		return m.UrlSuccess
	}
	return ""
}

func (m *ProjectOrder) GetUrlFail() string {
	if m != nil {
		return m.UrlFail
	}
	return ""
}

func (m *ProjectOrder) GetNotifyEmails() []string {
	if m != nil {
		return m.NotifyEmails
	}
	return nil
}

func (m *ProjectOrder) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *ProjectOrder) GetSendNotifyEmail() bool {
	if m != nil {
		return m.SendNotifyEmail
	}
	return false
}

func (m *ProjectOrder) GetUrlCheckAccount() string {
	if m != nil {
		return m.UrlCheckAccount
	}
	return ""
}

func (m *ProjectOrder) GetUrlProcessPayment() string {
	if m != nil {
		return m.UrlProcessPayment
	}
	return ""
}

func (m *ProjectOrder) GetCallbackProtocol() string {
	if m != nil {
		return m.CallbackProtocol
	}
	return ""
}

func (m *ProjectOrder) GetMerchant() *Merchant {
	if m != nil {
		return m.Merchant
	}
	return nil
}

func init() {
	proto.RegisterType((*Merchant)(nil), "processing.Merchant")
	proto.RegisterType((*FixedPackage)(nil), "processing.FixedPackage")
	proto.RegisterType((*FixedPackages)(nil), "processing.FixedPackages")
	proto.RegisterType((*ProjectPaymentMethod)(nil), "processing.ProjectPaymentMethod")
	proto.RegisterType((*ProjectPaymentMethods)(nil), "processing.ProjectPaymentMethods")
	proto.RegisterType((*Project)(nil), "processing.Project")
	proto.RegisterMapType((map[string]*FixedPackages)(nil), "processing.Project.FixedPackageEntry")
	proto.RegisterMapType((map[string]*ProjectPaymentMethods)(nil), "processing.Project.PaymentMethodsEntry")
	proto.RegisterType((*ProjectOrder)(nil), "processing.ProjectOrder")
	proto.RegisterEnum("processing.AccountingPeriod", AccountingPeriod_name, AccountingPeriod_value)
}

func init() { proto.RegisterFile("project.proto", fileDescriptor_project_daff29af18dd3507) }

var fileDescriptor_project_daff29af18dd3507 = []byte{
	// 1197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdb, 0x72, 0x1b, 0x35,
	0x18, 0xc6, 0x76, 0x7c, 0xfa, 0xed, 0x24, 0x6b, 0x25, 0x2d, 0xaa, 0xdb, 0x52, 0xd7, 0xc0, 0xe0,
	0x09, 0xe0, 0x32, 0xe5, 0xa2, 0xc0, 0x4c, 0x19, 0x4c, 0x69, 0x67, 0x0c, 0x3d, 0x78, 0xb6, 0x0d,
	0x17, 0xdc, 0x68, 0x94, 0x5d, 0x39, 0x11, 0xd9, 0x83, 0x91, 0xb4, 0x69, 0xfd, 0x04, 0x3c, 0x03,
	0x6f, 0xc1, 0x4b, 0xf1, 0x1e, 0x8c, 0x0e, 0xeb, 0xee, 0xda, 0x6e, 0x9b, 0xf6, 0x6e, 0xf5, 0xe9,
	0xd3, 0x27, 0xe9, 0x3f, 0x7c, 0x5a, 0xd8, 0x5d, 0x88, 0xf4, 0x4f, 0x16, 0xa8, 0xf1, 0x42, 0xa4,
	0x2a, 0x45, 0xb0, 0x10, 0x69, 0xc0, 0xa4, 0xe4, 0xc9, 0x69, 0xdf, 0x0b, 0x79, 0xa0, 0x78, 0x9a,
	0x50, 0xb1, 0xb4, 0xb3, 0xfd, 0x5b, 0xa7, 0x69, 0x7a, 0x1a, 0xb1, 0x3b, 0x66, 0x74, 0x92, 0xcd,
	0xef, 0x28, 0x1e, 0x33, 0xa9, 0x68, 0xbc, 0xb0, 0x84, 0xe1, 0xdf, 0x3b, 0xd0, 0x7a, 0xc2, 0x44,
	0x70, 0x46, 0x13, 0x85, 0xf6, 0xa0, 0xca, 0x43, 0x5c, 0x19, 0x54, 0x46, 0x6d, 0xbf, 0xca, 0x43,
	0x74, 0x0b, 0x3a, 0xec, 0x95, 0x62, 0x22, 0xa1, 0x11, 0xe1, 0x21, 0xae, 0x9a, 0x09, 0xc8, 0xa1,
	0x69, 0x88, 0x0e, 0xa1, 0xce, 0x62, 0xca, 0x23, 0x5c, 0x33, 0x53, 0x76, 0x80, 0x10, 0xec, 0x24,
	0x34, 0x66, 0x78, 0xc7, 0x80, 0xe6, 0x1b, 0x7d, 0x0d, 0xcd, 0x20, 0xcd, 0x12, 0x25, 0x96, 0xb8,
	0x3e, 0xa8, 0x8c, 0x3a, 0x77, 0x0f, 0xc6, 0xaf, 0x0f, 0x3e, 0x7e, 0x60, 0xa7, 0xfc, 0x9c, 0x83,
	0xa6, 0xd0, 0xa3, 0x81, 0x19, 0xf0, 0xe4, 0x94, 0x2c, 0x98, 0xe0, 0x69, 0x88, 0x1b, 0x83, 0xca,
	0x68, 0xef, 0xee, 0x8d, 0xe2, 0xc2, 0xc9, 0x8a, 0x34, 0x33, 0x1c, 0xdf, 0xa3, 0x6b, 0x08, 0xfa,
	0x06, 0x5a, 0x41, 0x26, 0x04, 0x4b, 0x82, 0x25, 0x6e, 0x9a, 0xad, 0x0f, 0x4b, 0x5b, 0xbb, 0x39,
	0x7f, 0xc5, 0x42, 0x9f, 0xc1, 0x1e, 0x97, 0xe4, 0x82, 0x2a, 0xc2, 0x12, 0x7a, 0x12, 0xb1, 0x10,
	0xb7, 0x06, 0x95, 0x51, 0xcb, 0xef, 0x72, 0xf9, 0x3b, 0x55, 0x0f, 0x2d, 0x86, 0x7e, 0x82, 0x9b,
	0x5c, 0x92, 0x20, 0x8d, 0x63, 0x2e, 0x25, 0x4f, 0x13, 0xa2, 0x52, 0x92, 0x49, 0x26, 0x56, 0x8b,
	0xda, 0x66, 0xd1, 0x35, 0x2e, 0x1f, 0xac, 0x38, 0x2f, 0xd2, 0x63, 0xc9, 0x44, 0xae, 0x70, 0x15,
	0x1a, 0x52, 0x51, 0x95, 0x49, 0x0c, 0x83, 0xca, 0xa8, 0xee, 0xbb, 0x11, 0xfa, 0x1e, 0x20, 0x10,
	0x8c, 0x2a, 0x16, 0x12, 0xaa, 0x70, 0xc7, 0x9c, 0xb9, 0x3f, 0xb6, 0x99, 0x1c, 0xe7, 0x99, 0x1c,
	0xbf, 0xc8, 0x33, 0xe9, 0xb7, 0x1d, 0x7b, 0xa2, 0xf4, 0xd2, 0x6c, 0x11, 0xe6, 0x4b, 0xbb, 0xef,
	0x5e, 0xea, 0xd8, 0x13, 0x35, 0xfc, 0xb7, 0x0a, 0xdd, 0x47, 0xfc, 0x15, 0x0b, 0x67, 0x34, 0x38,
	0xa7, 0xa7, 0x6c, 0xa3, 0x1a, 0xf2, 0xb4, 0x56, 0x0b, 0x69, 0xbd, 0x0d, 0xdd, 0x3c, 0x6c, 0x84,
	0x27, 0xca, 0xd4, 0x41, 0xdd, 0xef, 0xe4, 0xd8, 0x34, 0x51, 0xba, 0x46, 0x16, 0x82, 0x07, 0xb6,
	0x1c, 0x2a, 0xbe, 0x1d, 0xa0, 0xeb, 0xd0, 0xe6, 0x92, 0xd0, 0x40, 0xf1, 0x0b, 0x66, 0x2a, 0xa2,
	0xe5, 0xb7, 0xb8, 0x9c, 0x98, 0xf1, 0x5a, 0x00, 0x1a, 0x1f, 0x1e, 0x80, 0xe6, 0x7b, 0x04, 0xa0,
	0x54, 0x28, 0xad, 0xcb, 0x14, 0xca, 0xf0, 0x29, 0xec, 0x16, 0x23, 0x26, 0xd1, 0x7d, 0xd8, 0x9d,
	0x6b, 0x80, 0x2c, 0x2c, 0x82, 0x2b, 0x83, 0xda, 0xa8, 0x73, 0x17, 0x17, 0x75, 0x8a, 0x2b, 0xfc,
	0xee, 0xbc, 0x30, 0x1a, 0x12, 0x38, 0x9c, 0xd9, 0xe6, 0x9e, 0xd1, 0x65, 0xcc, 0x12, 0xf5, 0x84,
	0xa9, 0xb3, 0x34, 0xdc, 0xc8, 0xc4, 0x3d, 0x68, 0xd3, 0x30, 0xd4, 0x57, 0x9c, 0x28, 0x93, 0x8e,
	0xb7, 0xdf, 0xb1, 0x65, 0xc8, 0x93, 0x89, 0x1a, 0x9e, 0xc0, 0x95, 0x6d, 0x1b, 0x48, 0x34, 0x85,
	0xfd, 0x85, 0x45, 0x48, 0x6c, 0x21, 0x77, 0xf4, 0x41, 0xf1, 0xe8, 0xdb, 0xd6, 0xfa, 0x7b, 0x8b,
	0x92, 0xd4, 0xf0, 0x3f, 0x80, 0xa6, 0x23, 0x6e, 0x1c, 0x7c, 0x02, 0xbd, 0x80, 0x46, 0xd1, 0x09,
	0x0d, 0xce, 0xc9, 0x2a, 0xd6, 0xd5, 0xb7, 0xc4, 0xda, 0xcb, 0xe9, 0x39, 0x82, 0xbe, 0x2c, 0x48,
	0x98, 0xbb, 0x06, 0x69, 0x6e, 0x3f, 0x2b, 0xf2, 0xcc, 0xe1, 0xe8, 0x3e, 0x5c, 0xb7, 0xa5, 0x41,
	0x78, 0x72, 0x91, 0xf2, 0x80, 0x11, 0x1a, 0x45, 0xe9, 0x4b, 0x16, 0x92, 0x4c, 0x44, 0x12, 0xef,
	0x0c, 0x6a, 0xa3, 0xb6, 0x8f, 0x2d, 0x65, 0x6a, 0x19, 0x13, 0x4b, 0x38, 0x16, 0x91, 0xd4, 0x15,
	0x11, 0x3b, 0x6f, 0x74, 0xae, 0x55, 0x3a, 0x65, 0xee, 0x9b, 0xfe, 0x8a, 0x85, 0x7e, 0x84, 0x1b,
	0xba, 0xac, 0xb5, 0x06, 0x09, 0x97, 0x09, 0x8d, 0x79, 0x40, 0x92, 0x54, 0xf1, 0xf9, 0xd2, 0xee,
	0xd8, 0x30, 0x95, 0x8e, 0xb9, 0x34, 0xdb, 0xfc, 0x62, 0x19, 0x4f, 0x0d, 0xc1, 0xec, 0xf8, 0x33,
	0x7c, 0xb2, 0xb1, 0x5e, 0xb0, 0x90, 0x0b, 0x16, 0x28, 0xab, 0xd0, 0x34, 0x0a, 0xfd, 0xb2, 0x82,
	0xef, 0x28, 0x46, 0xe3, 0x3e, 0xec, 0x47, 0x3c, 0xe6, 0x4a, 0x92, 0x4b, 0x95, 0xf3, 0x9e, 0x25,
	0xaf, 0x02, 0x7c, 0x04, 0xde, 0x13, 0xfa, 0xca, 0xe5, 0x78, 0x12, 0x6b, 0x33, 0x35, 0x56, 0x56,
	0xf1, 0x37, 0x70, 0xc3, 0xe5, 0x49, 0x99, 0x0b, 0x8e, 0xbb, 0x86, 0x6b, 0xfb, 0x78, 0xaa, 0xed,
	0xa3, 0x63, 0xed, 0x43, 0x7f, 0xa3, 0x21, 0x74, 0xed, 0xe5, 0x1f, 0xea, 0x87, 0x43, 0xe2, 0xae,
	0x49, 0x48, 0x09, 0xd3, 0x7b, 0x3c, 0x4b, 0xa2, 0xa5, 0x69, 0x1b, 0x2b, 0x25, 0xf1, 0xae, 0x09,
	0xc2, 0x06, 0x8e, 0x6e, 0x40, 0xfb, 0x39, 0x0b, 0x04, 0x53, 0xbf, 0xb1, 0x25, 0xde, 0x33, 0x1b,
	0xbd, 0x06, 0xd0, 0x08, 0xf6, 0x9f, 0xb3, 0x24, 0x2c, 0xa8, 0xe3, 0x7d, 0x23, 0xb4, 0x0e, 0x6b,
	0xe6, 0xb1, 0xff, 0xf8, 0xc1, 0x19, 0x0b, 0xce, 0xdd, 0x0b, 0x83, 0x3d, 0xa3, 0xb6, 0x0e, 0xa3,
	0xaf, 0xa0, 0x77, 0xec, 0x3f, 0x9e, 0xd9, 0xb8, 0xba, 0x0b, 0xe3, 0x9e, 0xe1, 0x6e, 0x4e, 0x38,
	0xdd, 0x3c, 0x5b, 0x8f, 0xf4, 0x09, 0xd0, 0x4a, 0xb7, 0x08, 0xa3, 0x31, 0xa0, 0x02, 0xf4, 0x3c,
	0x0b, 0xb4, 0x0c, 0x3e, 0x30, 0xe4, 0x2d, 0x33, 0xe8, 0xd7, 0x75, 0xe7, 0xf9, 0xd8, 0xb4, 0xef,
	0xe7, 0x5b, 0xda, 0xb7, 0xe4, 0x40, 0x0f, 0xcd, 0xbb, 0x5b, 0xb2, 0x21, 0x34, 0xdb, 0x34, 0x03,
	0x6c, 0xd4, 0xbe, 0xd8, 0xa6, 0x56, 0x76, 0x12, 0xab, 0xb7, 0xe6, 0x09, 0xa8, 0x0f, 0xad, 0xa9,
	0x33, 0x77, 0x7c, 0x68, 0xcd, 0x7e, 0xba, 0xdd, 0xec, 0xaf, 0x7c, 0xb8, 0xd9, 0x5f, 0x7d, 0x0f,
	0xb3, 0xef, 0xff, 0x01, 0xbd, 0x8d, 0x30, 0x20, 0x0f, 0x6a, 0xe7, 0x6c, 0xe9, 0xfc, 0x4a, 0x7f,
	0xa2, 0x3b, 0x50, 0xbf, 0xa0, 0x51, 0xc6, 0x9c, 0x49, 0x5d, 0x7b, 0x93, 0x91, 0x4b, 0xdf, 0xf2,
	0x7e, 0xa8, 0x7e, 0x57, 0xe9, 0x87, 0x70, 0xb0, 0x25, 0x28, 0x5b, 0xd4, 0xef, 0x95, 0xd5, 0x6f,
	0xbf, 0xcb, 0x6b, 0x8b, 0xbb, 0x0c, 0xff, 0xa9, 0x41, 0xd7, 0x91, 0x9e, 0x89, 0x90, 0x89, 0x4b,
	0xbd, 0xd7, 0xb7, 0xa0, 0x93, 0x89, 0x88, 0x48, 0x57, 0x4f, 0xd6, 0x37, 0x21, 0x13, 0x51, 0x5e,
	0x47, 0xd7, 0xa0, 0xa5, 0x09, 0x73, 0x5d, 0x9a, 0xf6, 0xff, 0xad, 0x99, 0x89, 0xc8, 0x94, 0xe4,
	0xa7, 0xb0, 0xeb, 0xac, 0x8c, 0xd9, 0x6e, 0xad, 0xdb, 0x6e, 0x4d, 0x8a, 0xdd, 0x7a, 0x13, 0x40,
	0x9a, 0x86, 0x23, 0xfa, 0xae, 0x0d, 0xdb, 0x82, 0x72, 0xd5, 0x82, 0x47, 0xd0, 0x93, 0x2c, 0x09,
	0x49, 0x51, 0xc8, 0x59, 0xda, 0xbe, 0x5c, 0x6b, 0xc2, 0x23, 0xe8, 0xe9, 0xa3, 0x04, 0xba, 0xdd,
	0x88, 0xfb, 0xad, 0x33, 0x4e, 0xd6, 0xf6, 0xf7, 0x33, 0x11, 0x95, 0xda, 0x70, 0x0c, 0x07, 0x9a,
	0xeb, 0xe2, 0x47, 0x5c, 0xf9, 0x19, 0xdf, 0x6a, 0xfb, 0x5a, 0x66, 0xad, 0x11, 0xb7, 0xbe, 0x22,
	0xf0, 0x86, 0x57, 0xa4, 0xf8, 0x0c, 0x74, 0x2e, 0xf3, 0x0c, 0x1c, 0x85, 0xe0, 0xad, 0xff, 0x99,
	0xa2, 0x26, 0xd4, 0x42, 0xba, 0xf4, 0x3e, 0x42, 0x2d, 0xd8, 0x79, 0xc9, 0xd8, 0xb9, 0x57, 0x41,
	0x5d, 0x68, 0xa9, 0x97, 0x29, 0x31, 0xa3, 0x2a, 0x6a, 0x43, 0x3d, 0x4e, 0x13, 0x75, 0xe6, 0xd5,
	0x50, 0x07, 0x9a, 0x7f, 0x65, 0x54, 0x28, 0x26, 0xbc, 0x1d, 0xb4, 0x0b, 0xed, 0x33, 0x1a, 0xcd,
	0xc9, 0x92, 0x51, 0xe1, 0xd5, 0xf5, 0x72, 0xf3, 0xd5, 0x38, 0x69, 0x98, 0x93, 0x7f, 0xfb, 0x7f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x71, 0xda, 0xc2, 0x12, 0x0c, 0x00, 0x00,
}