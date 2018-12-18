package repository

import (
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/micro/protobuf/ptypes"
)

const (
	ProjectFieldNameUrlSuccess        = "url_success"
	ProjectFieldNameUrlFail           = "url_fail"
	ProjectFieldNameNotifyEmails      = "notify_emails"
	ProjectFieldNameSecretKey         = "secret_key"
	ProjectFieldNameSendNotifyEmail   = "send_notify_email"
	ProjectFieldNameUrlCheckAccount   = "url_check_account"
	ProjectFieldNameUrlProcessPayment = "url_process_payment"
	ProjectFieldNameCallbackProtocol  = "callback_protocol"
	ProjectFieldNameMerchant          = "merchant"

	MerchantFieldNameExternalId                = "external_id"
	MerchantFieldNameAccountingPeriod          = "accounting_period"
	MerchantFieldNameIsVatEnabled              = "is_vat_enabled"
	MerchantFieldNameIsCommissionToUserEnabled = "is_commission_to_user_enabled"
)

func (r *Repository) getMerchantMap(m *billing.Merchant) map[string]interface{} {
	mm := make(map[string]interface{})

	//mm[FieldNameUnderscoreId] = ByteToObjectId(m.GetId())
	mm[MerchantFieldNameExternalId] = m.GetExternalId()
	mm[FieldNameEmail] = m.GetEmail()
	mm[FieldNameName] = m.GetName()
	mm[FieldNameCountry] = "--------------------------------------------------"
	mm[MerchantFieldNameAccountingPeriod] = m.GetAccountingPeriod()
	mm[FieldNameCurrency] = "-------------------------------------------------"
	mm[MerchantFieldNameIsVatEnabled] = m.GetIsVatEnabled()
	mm[MerchantFieldNameIsCommissionToUserEnabled] = m.GetIsCommissionToUserEnabled()
	mm[FieldNameStatus] = m.GetStatus()

	if m.GetCreatedAt() != nil {
		if v, err := ptypes.Timestamp(m.GetCreatedAt()); err == nil {
			mm[FieldNameCreatedAt] = v
		}
	}

	if m.GetUpdatedAt() != nil {
		if v, err := ptypes.Timestamp(m.GetUpdatedAt()); err == nil {
			mm[FieldNameUpdatedAt] = v
		}
	}

	return mm
}

func (r *Repository) getProjectOrderMap(po *billing.ProjectOrder) map[string]interface{} {
	return map[string]interface{}{
		//FieldNameId:                       ByteToObjectId(po.GetId()),
		FieldNameName:                     po.GetName(),
		ProjectFieldNameUrlSuccess:        po.GetUrlSuccess(),
		ProjectFieldNameUrlFail:           po.GetUrlFail(),
		ProjectFieldNameNotifyEmails:      po.GetNotifyEmails(),
		ProjectFieldNameSecretKey:         po.GetSecretKey(),
		ProjectFieldNameSendNotifyEmail:   po.GetSendNotifyEmail(),
		ProjectFieldNameUrlCheckAccount:   po.GetUrlCheckAccount(),
		ProjectFieldNameUrlProcessPayment: po.GetUrlProcessPayment(),
		ProjectFieldNameCallbackProtocol:  po.GetCallbackProtocol(),
		ProjectFieldNameMerchant:          r.getMerchantMap(po.Merchant),
	}
}
