package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"github.com/micro/protobuf/ptypes"
	"log"
)

const (
	OrderFieldNameProject                = "project"
	OrderFieldNameProjectOrderId         = "project_order_id"
	OrderFieldNameProjectAccount         = "project_account"
	OrderFieldNameDescription            = "description"
	OrderFieldNameProjectIncomeAmount    = "project_income_amount"
	OrderFieldNameProjectIncomeCurrency  = "project_income_currency"
	OrderFieldNameProjectOutcomeAmount   = "project_outcome_amount"
	OrderFieldNameProjectOutcomeCurrency = "project_outcome_currency"
	OrderFieldNameProjectLastRequestedAt = "project_last_requested_at"
	OrderFieldNameProjectParams          = "project_params"
	OrderFieldNamePayerData              = "payer_data"
	OrderFieldNamePaymentMethod          = "payment_method"
	OrderFieldNamePmTerminalId           = "pm_terminal_id"
	OrderFieldNamePmOrderId              = "pm_order_id"
	OrderFieldNamePmOutcomeAmount        = "pm_outcome_amount"
	OrderFieldNamePmOutcomeCurrency      = "pm_outcome_currency"
	OrderFieldNamePmIncomeAmount         = "pm_income_amount"
	OrderFieldNamePmIncomeCurrency       = "pm_income_currency"
	OrderFieldNamePmOrderCloseDate       = "pm_order_close_date"
	OrderFieldNameStatus                 = "status"
	OrderFieldNameCreatedByJson          = "created_by_json"
	OrderFieldNameAmountPspAc            = "amount_psp_ac"
	OrderFieldNameAmountInMerchantAc     = "amount_in_merchant_ac"
	OrderFieldNameAmountOutMerchantAc    = "amount_out_merchant_ac"
	OrderFieldNameAmountPsAc             = "amount_ps_ac"
	OrderFieldNamePmAccount              = "pm_account"
	OrderFieldNamePmTxnParams            = "pm_txn_params"
	OrderFieldNameFixedPackage           = "fixed_package"
	OrderFieldNamePaymentRequisites      = "payment_requisites"
	OrderFieldNamePspFeeAmount           = "psp_fee_amount"
	OrderFieldNameProjectFeeAmount       = "project_fee_amount"
	OrderFieldNameToPayerFeeAmount       = "to_payer_fee_amount"
	OrderFieldNameVatAmount              = "vat_amount"
	OrderFieldNamePsFeeAmount            = "ps_fee_amount"
)

func (r *Repository) UpdateOrder(ctx context.Context, req *billing.Order, rsp *repository.Result) error {
	m := r.toMap(req)
	id := m[FieldNameUnderscoreId]

	delete(m, FieldNameUnderscoreId)

	err := r.Database.Collection(CollectionOrder).UpdateId(id, bson.M{"$set": m})

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}

func (r *Repository) getOrderMap(o *billing.Order) map[string]interface{} {
	om := make(map[string]interface{})

	//om[FieldNameUnderscoreId] = ByteToObjectId(o.Id)
	om[OrderFieldNameProject] = r.getProjectOrderMap(o.GetProject())
	om[OrderFieldNameProjectOrderId] = o.GetProjectOrderId()
	om[OrderFieldNameProjectAccount] = o.GetProjectAccount()
	om[OrderFieldNameDescription] = o.GetDescription()
	om[OrderFieldNameProjectIncomeAmount] = o.GetProjectIncomeAmount()
	om[OrderFieldNameProjectIncomeCurrency] = "------------------------------------------"
	om[OrderFieldNameProjectOutcomeAmount] = o.GetProjectOutcomeAmount()
	om[OrderFieldNameProjectOutcomeCurrency] = "------------------------------------------"
	om[OrderFieldNameProjectParams] = o.GetProjectParams()
	om[OrderFieldNamePayerData] = o.GetPayerData()
	om[OrderFieldNamePaymentMethod] = o.GetPaymentMethod()
	om[OrderFieldNamePmTerminalId] = o.GetPaymentMethodTerminalId()
	om[OrderFieldNamePmOrderId] = o.GetPaymentMethodOrderId()
	om[OrderFieldNamePmOutcomeAmount] = o.GetPaymentMethodOutcomeAmount()
	om[OrderFieldNamePmOutcomeCurrency] = "------------------------------------------"
	om[OrderFieldNamePmIncomeAmount] = o.GetPaymentMethodIncomeAmount()
	om[OrderFieldNamePmIncomeCurrency] = "------------------------------------------"
	om[OrderFieldNameStatus] = o.GetStatus()
	om[OrderFieldNameCreatedByJson] = o.GetIsJsonRequest()
	om[OrderFieldNameAmountPspAc] = o.GetAmountInPspAccountingCurrency()
	om[OrderFieldNameAmountInMerchantAc] = o.GetAmountInMerchantAccountingCurrency()
	om[OrderFieldNameAmountOutMerchantAc] = o.GetAmountOutMerchantAccountingCurrency()
	om[OrderFieldNameAmountPsAc] = o.GetAmountInPaymentSystemAccountingCurrency()
	om[OrderFieldNamePmAccount] = o.GetPaymentMethodPayerAccount()
	om[OrderFieldNamePmTxnParams] = o.GetPaymentMethodTxnParams()
	om[OrderFieldNameFixedPackage] = o.GetFixedPackage()
	om[OrderFieldNamePaymentRequisites] = o.GetPaymentRequisites()
	om[OrderFieldNamePspFeeAmount] = "------------------------------------------"
	om[OrderFieldNameProjectFeeAmount] = "------------------------------------------"
	om[OrderFieldNameToPayerFeeAmount] = "------------------------------------------"
	om[OrderFieldNameVatAmount] = o.GetVatAmount()
	om[OrderFieldNamePsFeeAmount] = "------------------------------------------"

	if o.GetProjectLastRequestedAt() != nil {
		if v, err := ptypes.Timestamp(o.GetProjectLastRequestedAt()); err == nil {
			om[OrderFieldNameProjectLastRequestedAt] = v
		}
	}

	if o.GetCreatedAt() != nil {
		if v, err := ptypes.Timestamp(o.GetCreatedAt()); err == nil {
			om[FieldNameCreatedAt] = v
		}
	}

	if o.GetUpdatedAt() != nil {
		if v, err := ptypes.Timestamp(o.GetUpdatedAt()); err == nil {
			om[FieldNameUpdatedAt] = v
		}
	}

	if o.GetPaymentMethodOrderClosedAt() != nil {
		if v, err := ptypes.Timestamp(o.GetPaymentMethodOrderClosedAt()); err == nil {
			om[OrderFieldNamePmOrderCloseDate] = v
		}
	}

	return om
}
