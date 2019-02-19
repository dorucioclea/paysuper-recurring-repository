package repository

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/tools"
)

const (
	FieldNameUnderscoreId = "_id"

	QueryErrorMask           = "[PAYONE ERROR] Query from table \"%s\" ended with error: %s\n"
	NotFoundGeneralErrorMask = "[PAYONE ERROR] %s not found\n"
	NotFoundErrorMask        = "[PAYONE ERROR] %s not found for specified project and payment method\n"
	SomeErrorMask            = "[PAYONE ERROR] Something is wrong. Operation ended with error: %s\n"
	AlreadyExistErrorMask    = "[PAYONE ERROR] %s record with specified data already exist"

	CollectionOrder         = "order"
	CollectionMerchant      = "merchant"
	CollectionCurrencyRate  = "currency_rate"
	CollectionSavedCard     = "saved_card"
)

type Repository struct {
	Database *database.Source
}

func (r *Repository) toMap(obj interface{}) map[string]interface{} {
	return tools.NewStructureConverter(obj).Map()
}
