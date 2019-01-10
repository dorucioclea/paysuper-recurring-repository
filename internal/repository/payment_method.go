package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"log"
)

func (r *Repository) FindPaymentMethodsByGroupAndCurrency(ctx context.Context, req *repository.FindByGroupCurrencyRequest, rsp *repository.PaymentMethods) error {
	query := r.toMap(req)
	query["is_active"] = true

	err := r.Database.Collection(CollectionPaymentMethod).Find(query).All(&rsp.PaymentMethods)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionPaymentMethod, err.Error())
		return err
	}

	return nil
}
