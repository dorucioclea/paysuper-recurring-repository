package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
)

const (
	OrderStatusNew                         = 0
	OrderStatusPaymentSystemCreate         = 1
	OrderStatusPaymentSystemRejectOnCreate = 2
	OrderStatusPaymentSystemReject         = 3
	OrderStatusPaymentSystemComplete       = 4
	OrderStatusProjectInProgress           = 5
	OrderStatusProjectComplete             = 6
	OrderStatusProjectPending              = 7
	OrderStatusProjectReject               = 8
	OrderStatusRefund                      = 9
	OrderStatusChargeback                  = 10
	OrderStatusPaymentSystemDeclined       = 11
	OrderStatusPaymentSystemCanceled       = 12
)

func (r *Repository) UpdateOrder(ctx context.Context, req *billing.Order, rsp *repository.Result) error {
	m := r.toMap(req)
	id := m[FieldNameUnderscoreId]

	delete(m, FieldNameUnderscoreId)

	err := r.Database.Collection(CollectionOrder).UpdateId(id, bson.M{"$set": m})

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}
