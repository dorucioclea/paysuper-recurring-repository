package constant

const (
	PayOneMicroserviceVersion   = "latest"
	PayOneRepositoryServiceName = "p1payrepository"

	PayOneTopicNotifyPaymentName  = "notify-payment"
	PayOneTopicNotifyMerchantName = "notify-merchant"

	TaxjarTransactionsTopicName      = "taxjar-sync-transactions"
	TaxjarTransactionsRetryTopicName = "taxjar-sync-transactions-retry"
	TaxjarRefundsTopicName           = "taxjar-sync-refunds"
	TaxjarRefundsRetryTopicName      = "taxjar-sync-refunds-retry"

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

	OrderPublicStatusCreated    = "created"
	OrderPublicStatusProcessed  = "processed"
	OrderPublicStatusCanceled   = "canceled"
	OrderPublicStatusRejected   = "rejected"
	OrderPublicStatusRefunded   = "refunded"
	OrderPublicStatusChargeback = "chargeback"
	OrderPublicStatusPending    = "pending"

	RegistryKubernetes = "kubernetes"
)
