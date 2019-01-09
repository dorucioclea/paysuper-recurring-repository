package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
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

func (r *Repository) FindOrderBy(ctx context.Context, req *repository.FindByRequest, rsp *billing.Order) error {
	err := r.Database.Collection(CollectionOrder).Find(req.Query).Limit(int(req.Limit)).Skip(int(req.Offset)).All(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionOrder, err.Error())
		return err
	}

	return nil
}
