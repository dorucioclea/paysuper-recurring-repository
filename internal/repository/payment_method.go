package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
)

func (r *Repository) FindPaymentMethodByGroupAndCurrency(ctx context.Context, req *repository.FindByGroupCurrencyRequest, rsp *billing.PaymentMethod) error {
	query := r.toMap(req)

	err := r.Database.Collection(CollectionPaymentMethod).Find(query).Sort("-updated_at").One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionPaymentMethod, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindPaymentMethodsByCurrency(ctx context.Context, req *repository.FindByIntValue, rsp *repository.PaymentMethods) error {
	var pms []*billing.PaymentMethod

	query := bson.M{"currency": req.Value, "is_active": true}
	err := r.Database.Collection(CollectionPaymentMethod).Find(query).All(&pms)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionPaymentMethod, err.Error())
		return err
	}

	rsp.PaymentMethods = pms

	return nil
}
