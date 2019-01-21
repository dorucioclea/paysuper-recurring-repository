package repository

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/tools"
)

const (
	FieldNameUnderscoreId = "_id"

	QueryErrorMask    = "[PAYONE ERROR] Query from table \"%s\" ended with error: %s\n"
	NotFoundErrorMask = "[PAYONE ERROR] %s not found for specified project and payment method\n"
	SomeErrorMask     = "[PAYONE ERROR] Something is wrong. Operation ended with error: %s\n"

	CollectionOrder         = "order"
	CollectionMerchant      = "merchant"
	CollectionCurrencyRate  = "currency_rate"
	CollectionProject       = "project"
	CollectionPaymentMethod = "payment_method"
	CollectionCurrency      = "currency"
	CollectionCommission    = "commission"
	CollectionVat           = "vat"
	CollectionBinData       = "bank_bin"
	CollectionSavedCard     = "saved_card"
)

var VatBySubdivisionCountries = map[string]bool{"US": true, "CA": true}

type Repository struct {
	Database *database.Source
}

func (r *Repository) toMap(obj interface{}) map[string]interface{} {
	return tools.NewStructureConverter(obj).Map()
}
