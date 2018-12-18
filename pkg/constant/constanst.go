package constant

const (
	PayOneMicroserviceVersion = "latest"

	PayOneSubscriberNotifierName = "go.p1.payone.notifier-subscriber"
	PayOneRepositoryServiceName  = "go.p1.payone.repository-server"

	PayOneTopicNotifyPaymentName = "payone.topic.notify-payment"

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
)
