package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
)

func (r *Repository) UpdateMerchant(ctx context.Context, req *billing.Merchant, rsp *repository.Result) error {
	m := r.toMap(req)
	id := m[FieldNameUnderscoreId]

	delete(m, FieldNameUnderscoreId)

	err := r.Database.Collection(CollectionMerchant).UpdateId(id, bson.M{"$set": m})

	if err != nil {
		log.Printf(QueryErrorMask, CollectionMerchant, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindMerchantById(ctx context.Context, req *repository.FindByStringValue, rsp *billing.Merchant) error {
	err := r.Database.Collection(CollectionMerchant).Find(bson.M{"external_id": req.Value}).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionMerchant, err.Error())
		return err
	}

	return nil
}