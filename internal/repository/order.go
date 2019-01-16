package repository

import (
	"context"
	"errors"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
	"strconv"
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

func (r *Repository) InsertOrder(ctx context.Context, req *billing.Order, rsp *repository.Result) error {
	err := r.Database.Collection(CollectionOrder).Insert(req)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindOrderById(ctx context.Context, req *repository.FindByUnderscoreId, rsp *billing.Order) error {
	err := r.Database.Collection(CollectionOrder).Find(req).All(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindOrderByProjectAndOrderId(ctx context.Context, req *repository.FindByProjectOrderId, rsp *billing.Order) error {
	err := r.Database.Collection(CollectionBinData).Find(r.toMap(req)).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindBinData(ctx context.Context, req *repository.FindByStringValue, rsp *billing.BinData) error {
	if len(req.Value) < 6 {
		err := errors.New("number to find by bin base must be greater than or equal 6 digits")

		log.Printf(SomeErrorMask, err.Error())
		return err
	}

	if len(req.Value) > 6 {
		req.Value = req.Value[:6]
	}

	i, err := strconv.ParseInt(req.Value[:6], 10, 32)

	if err != nil {
		log.Printf(SomeErrorMask, err.Error())
		return err
	}

	val := int32(i)

	err = r.Database.Collection(CollectionBinData).Find(bson.M{"card_bin": val}).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionBinData, err.Error())
		return err
	}

	return nil
}