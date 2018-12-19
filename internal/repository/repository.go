package repository

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/tools"
)

const (
	FieldNameUnderscoreId = "_id"

	QueryErrorMask = "Query from table \"%s\" ended with error: %s"

	CollectionOrder        = "order"
	CollectionCurrencyRate = "currency_rate"
)

type Repository struct {
	Database *database.Source
}

func (r *Repository) toMap(obj interface{}) map[string]interface{} {
	return tools.NewStructureConverter(obj).Map()
}
