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
	OrderFieldNameProject = "project"
	OrderFieldNameProjectOrderId = "project_order_id"
	OrderFieldNameProjectAccount = "project_account"
	OrderFieldNameDescription = "description"
	OrderFieldNameProjectIncomeAmount = "project_income_amount"
	OrderFieldNameProjectIncomeCurrency = "project_income_currency"
	OrderFieldNameProjectOutcomeAmount = "project_outcome_amount"
	OrderFieldNameProjectOutcomeCurrency = "project_outcome_currency"
	OrderFieldNameProjectLastRequestedAt = "project_last_requested_at"
	OrderFieldNameProjectParams = "project_params"
	OrderFieldNamePayerData = "payer_data"
	OrderFieldNamePaymentMethod = "payment_method"
	OrderFieldNamePmTerminalId = "pm_terminal_id"
	OrderFieldNamePmOrderId = "pm_order_id"
	OrderFieldNamePmOutcomeAmount = "pm_outcome_amount"
	OrderFieldNamePmOutcomeCurrency = "pm_outcome_currency"
	OrderFieldNamePmIncomeAmount = "pm_income_amount"
	OrderFieldNamePmIncomeCurrency = "pm_income_currency"
	OrderFieldNamePmOrderCloseDate = "pm_order_close_date"
	OrderFieldNameStatus = "status"
	OrderFieldNameCreatedAt = "created_at"
	OrderFieldNameUpdatedAt = "updated_at"
	OrderFieldNameCreatedByJson = "created_by_json"
	OrderFieldNameAmountPspAc = "amount_psp_ac"
	OrderFieldNameAmountInMerchantAc = "amount_in_merchant_ac"
	OrderFieldNameAmountOutMerchantAc = "amount_out_merchant_ac"
	OrderFieldNameAmountPsAc = "amount_ps_ac"
	OrderFieldNamePmAccount = "pm_account"
	OrderFieldNamePmTxnParams = "pm_txn_params"
	OrderFieldNameFixedPackage = "fixed_package"
	OrderFieldNamePaymentRequisites = "payment_requisites"
	OrderFieldNamePspFeeAmount = "psp_fee_amount"
	OrderFieldNameProjectFeeAmount = "project_fee_amount"
	OrderFieldNameToPayerFeeAmount = "to_payer_fee_amount"
	OrderFieldNameVatAmount = "vat_amount"
	OrderFieldNamePsFeeAmount = "ps_fee_amount"
)

func (r *Repository) UpdateOrder(ctx context.Context, req *billing.Order, rsp *repository.Result) error {
	changes := bson.M{
		OrderFieldNameProject: req.GetProject(),
		OrderFieldNameProjectOrderId: req.GetProjectOrderId(),
		OrderFieldNameProjectAccount: req.GetProjectAccount(),
		OrderFieldNameDescription: req.GetDescription(),
		OrderFieldNameProjectIncomeAmount: req.GetProjectIncomeAmount(),
		OrderFieldNameProjectIncomeCurrency: req.GetProjectIncomeCurrency(),
		OrderFieldNameProjectOutcomeAmount: req.GetProjectOutcomeAmount(),
		OrderFieldNameProjectOutcomeCurrency: req.GetProjectOutcomeCurrency(),
		OrderFieldNameProjectParams: req.GetProjectParams(),
		OrderFieldNamePayerData: req.GetPayerData(),
		OrderFieldNamePaymentMethod: req.GetPaymentMethod(),
		OrderFieldNamePmTerminalId: req.GetPaymentMethodTerminalId(),
		OrderFieldNamePmOrderId: req.GetPaymentMethodOrderId(),
		OrderFieldNamePmOutcomeAmount: req.GetPaymentMethodOutcomeAmount(),
		OrderFieldNamePmOutcomeCurrency: req.GetPaymentMethodOutcomeCurrency(),
		OrderFieldNamePmIncomeAmount: req.GetPaymentMethodIncomeAmount(),
		OrderFieldNamePmIncomeCurrency: req.GetPaymentMethodIncomeCurrency(),
		OrderFieldNameStatus: req.GetStatus(),
		OrderFieldNameCreatedByJson: req.GetIsJsonRequest(),
		OrderFieldNameAmountPspAc: req.GetAmountInPspAccountingCurrency(),
		OrderFieldNameAmountInMerchantAc: req.GetAmountInMerchantAccountingCurrency(),
		OrderFieldNameAmountOutMerchantAc: req.GetAmountOutMerchantAccountingCurrency(),
		OrderFieldNameAmountPsAc: req.GetAmountInPaymentSystemAccountingCurrency(),
		OrderFieldNamePmAccount: req.GetPaymentMethodPayerAccount(),
		OrderFieldNamePmTxnParams: req.GetPaymentMethodTxnParams(),
		OrderFieldNameFixedPackage: req.GetFixedPackage(),
		OrderFieldNamePaymentRequisites: req.GetPaymentRequisites(),
		OrderFieldNamePspFeeAmount: req.GetPspFeeAmount(),
		OrderFieldNameProjectFeeAmount: req.GetProjectFeeAmount(),
		OrderFieldNameToPayerFeeAmount: req.GetToPayerFeeAmount(),
		OrderFieldNameVatAmount: req.GetVatAmount(),
		OrderFieldNamePsFeeAmount: req.GetPaymentSystemFeeAmount(),
	}

	if req.GetProjectLastRequestedAt() != nil {
		if v, err := ptypes.Timestamp(req.GetProjectLastRequestedAt()); err == nil {
			changes[OrderFieldNameProjectLastRequestedAt] = v
		}
	}

	if req.GetCreatedAt() != nil {
		if v, err := ptypes.Timestamp(req.GetCreatedAt()); err == nil {
			changes[OrderFieldNameCreatedAt] = v
		}
	}

	if req.GetUpdatedAt() != nil {
		if v, err := ptypes.Timestamp(req.GetUpdatedAt()); err == nil {
			changes[OrderFieldNameUpdatedAt] = v
		}
	}

	if req.GetPaymentMethodOrderClosedAt() != nil {
		if v, err := ptypes.Timestamp(req.GetPaymentMethodOrderClosedAt()); err == nil {
			changes[OrderFieldNamePmOrderCloseDate] = v
		}
	}

	err := r.Database.Collection(CollectionOrder).UpdateId(ByteToObjectId(req.Id), bson.M{"$set": changes})

	if err != nil {
		rsp.HasError = true
		rsp.Comment = err.Error()

		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
	}

	return nil
}
