package repository

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/tools"
)

const (
	FieldNameUnderscoreId = "_id"

	QueryErrorMask = "[PAYONE ERROR] Query from table \"%s\" ended with error: %s\n"
	CommissionNotFoundError = "[PAYONE ERROR] Commission not found for specified project and payment method"

	CollectionOrder         = "order"
	CollectionMerchant      = "merchant"
	CollectionCurrencyRate  = "currency_rate"
	CollectionProject       = "project"
	CollectionPaymentMethod = "payment_method"
	CollectionCurrency      = "currency"
	CollectionCommission    = "commission"
)

type Repository struct {
	Database  *database.Source
}

func (r *Repository) toMap(obj interface{}) map[string]interface{} {
	return tools.NewStructureConverter(obj).Map()
}
