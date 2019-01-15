package repository

import (
	"context"
	"errors"
	"fmt"
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
		mes := fmt.Sprintf(NotFoundErrorMask, "Commission")

		log.Printf(mes)
		return errors.New(mes)
	}

	rsp.PaymentMethod = req.Amount * (c.PaymentMethodCommission / 100)
	rsp.Psp = req.Amount * (c.PspCommission / 100)
	rsp.ToUser = req.Amount * (c.TotalCommissionToUser / 100)

	return nil
}

func (r *Repository) CalculateVat(ctx context.Context, req *repository.CalculateVatRequest, rsp *repository.FloatValue) error {
	var vat *billing.Vat

	q := bson.M{"country.code_a2": req.CountryCodeA2}

	vsFlag, ok := VatBySubdivisionCountries[req.CountryCodeA2]

	if ok && vsFlag == true && len(req.Subdivision) > 0 {
		q["subdivision_code"] = req.Subdivision
	}

	err := r.Database.Collection(CollectionVat).Find(q).Sort("-created_at").Limit(1).One(&vat)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionVat, err.Error())
		return err
	}

	if vat == nil {
		mes := fmt.Sprintf(NotFoundErrorMask, "Vat")

		log.Printf(mes)
		return errors.New(mes)
	}

	rsp.Value = req.Amount * (vat.Vat / 100)

	return nil
}
