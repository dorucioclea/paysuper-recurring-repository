package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
)

func (r *Repository) FindCurrencyByCodeA3(ctx context.Context, req *repository.FindByStringValue, rsp *billing.Currency) error {
	err := r.Database.Collection(CollectionCurrency).Find(bson.M{"code_a3": req.GetValue(), "is_active": true}).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionCurrency, err.Error())
		return err
	}

	return nil
}
