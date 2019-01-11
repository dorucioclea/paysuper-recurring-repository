package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
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
