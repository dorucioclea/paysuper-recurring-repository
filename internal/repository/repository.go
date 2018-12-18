package repository

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/tools"
)

const (
	FieldNameUnderscoreId = "_id"
	FieldNameId           = "id"
	FieldNameName         = "name"
	FieldNameEmail        = "email"
	FieldNameCountry      = "country"
	FieldNameCurrency     = "currency"
	FieldNameCreatedAt    = "created_at"
	FieldNameUpdatedAt    = "updated_at"
	FieldNameStatus       = "status"

	QueryErrorMask = "Query from table \"%s\" ended with error: %s"

	CollectionOrder = "order"
)

type Repository struct {
	Database *database.Source
}

func (r *Repository) toMap(obj interface{}) map[string]interface{} {
	return tools.NewStructureConverter(obj).Map()
}
