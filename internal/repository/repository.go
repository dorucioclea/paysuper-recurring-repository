package repository

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/tools"
	"github.com/oschwald/geoip2-golang"
)

const (
	FieldNameUnderscoreId = "_id"

	QueryErrorMask = "[PAYONE ERROR] Query from table \"%s\" ended with error: %s\n"
	GeoIpErrorMask = "[PAYONE ERROR] Getting geo information about ip %s ending with error: %s\n"

	CollectionOrder         = "order"
	CollectionMerchant      = "merchant"
	CollectionCurrencyRate  = "currency_rate"
	CollectionProject       = "project"
	CollectionPaymentMethod = "payment_method"
	CollectionCurrency      = "currency"
)

type Repository struct {
	Database  *database.Source
	GeoReader *geoip2.Reader
}

func (r *Repository) toMap(obj interface{}) map[string]interface{} {
	return tools.NewStructureConverter(obj).Map()
}
