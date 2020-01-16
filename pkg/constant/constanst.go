package constant

import "errors"

const (
	PayOneTopicNotifyPaymentName  = "notify-payment"
	PayOneTopicNotifyMerchantName = "notify-merchant"

	TaxjarTransactionsTopicName      = "taxjar-sync-transactions"
	TaxjarTransactionsRetryTopicName = "taxjar-sync-transactions-retry"
	TaxjarRefundsTopicName           = "taxjar-sync-refunds"
	TaxjarRefundsRetryTopicName      = "taxjar-sync-refunds-retry"
	TaxjarNotificationStatusPayment  = "taxjar-payment"
	TaxjarNotificationStatusRefund   = "taxjar-refund"

	PaymentSystemCardPayDateFormat  = "2006-01-02T15:04:05Z"
	PaymentSystemGroupAliasBankCard = "BANKCARD"
	PaymentSystemGroupAliasWebMoney = "WEBMONEY"
	PaymentSystemGroupAliasQiwi     = "QIWI"
	PaymentSystemGroupAliasNeteller = "NETELLER"
	PaymentSystemGroupAliasAlipay   = "ALIPAY"
	PaymentSystemGroupAliasBitcoin  = "BITCOIN"

	OrderStatusNew                         = 0
	OrderStatusPaymentSystemCreate         = 1
	OrderStatusPaymentSystemRejectOnCreate = 2
	OrderStatusPaymentSystemReject         = 3
	OrderStatusPaymentSystemComplete       = 4
	OrderStatusProjectInProgress           = 5
	OrderStatusProjectComplete             = 6
	OrderStatusProjectPending              = 7
	OrderStatusProjectReject               = 8
	OrderStatusRefund                      = 9
	OrderStatusChargeback                  = 10
	OrderStatusPaymentSystemDeclined       = 11
	OrderStatusPaymentSystemCanceled       = 12
	OrderStatusItemReplaced                = 13

	RegistryKubernetes = "kubernetes"

	CollectionSavedCard = "saved_card"

	SavedCardErrorNotFound = "saved card with specified identifier not found"
	SavedCardErrorDatabase = "saved cards database query failed"

	ErrorDatabaseQueryFailed          = "Query to database collection failed"
	ErrorQueryCursorExecutionFailed   = "Execute result from query cursor failed"
	ErrorDatabaseFieldCollection      = "collection"
	ErrorDatabaseFieldDocumentId      = "document_id"
	ErrorDatabaseFieldQuery           = "query"
	ErrorDatabaseFieldSet             = "set"
	ErrorDatabaseFieldSorts           = "sorts"
	ErrorDatabaseFieldLimit           = "limit"
	ErrorDatabaseFieldOffset          = "offset"
	ErrorDatabaseFieldOperation       = "operation"
	ErrorDatabaseFieldOperationCount  = "count"
	ErrorDatabaseFieldOperationInsert = "insert"
	ErrorDatabaseFieldOperationUpdate = "update"
	ErrorDatabaseFieldOperationUpsert = "upsert"
	ErrorDatabaseFieldDocument        = "document"
)

var (
	ErrNotFound = errors.New(SavedCardErrorNotFound)
	ErrDatabase = errors.New(SavedCardErrorDatabase)
)
