package repository

import (
	"context"
	"errors"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
	"time"
)

func (r *Repository) CalculateCommission(ctx context.Context, req *repository.CommissionRequest, rsp *repository.CommissionResponse) error {
	var c *billing.Commission

	q := bson.M{
		"pm_id":      bson.ObjectIdHex(req.PaymentMethodId),
		"project_id": bson.ObjectIdHex(req.ProjectId),
		"start_date": bson.M{"$lte": time.Now()},
	}

	err := r.Database.Collection(CollectionCommission).Find(q).Sort("-start_date").Limit(1).One(&c)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionCommission, err.Error())
		return err
	}

	if c == nil {
		log.Println(CommissionNotFoundError)
		return errors.New(CommissionNotFoundError)
	}

	rsp.PaymentMethod = req.Amount * (c.PaymentMethodCommission / 100)
	rsp.Psp = req.Amount * (c.PspCommission / 100)
	rsp.ToUser = req.Amount * (c.TotalCommissionToUser / 100)

	return nil
}
