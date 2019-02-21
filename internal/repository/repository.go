package repository

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-recurring-repository/internal/database"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/entity"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/repository"
	"go.uber.org/zap"
)

const (
	QueryErrorMask        = "[PAYSUPER_REPOSITORY] Query to saved cards collection failed"
	AlreadyExistErrorMask = "[PAYSUPER_REPOSITORY] Saved card with specified data already exist"

	CollectionSavedCard = "saved_card"
)

type Repository struct {
	Database *database.Source
}

func (r *Repository) InsertSavedCard(
	ctx context.Context,
	req *repository.SavedCardRequest,
	rsp *repository.Result,
) error {
	var savedCard *entity.SavedCard

	q := bson.M{
		"account":    req.Account,
		"project_id": bson.ObjectIdHex(req.ProjectId),
		"pan":        req.MaskedPan,
	}
	err := r.Database.Collection(CollectionSavedCard).Find(q).One(&savedCard)

	if err != nil && err != mgo.ErrNotFound {
		zap.L().Error(QueryErrorMask, zap.Error(err), zap.Any("filter", q))
		return err
	}

	if savedCard != nil {
		zap.L().Info(AlreadyExistErrorMask, zap.Any("filter", q))
		return nil
	}

	data := &entity.SavedCard{
		Id:          bson.NewObjectId().Hex(),
		Account:     req.Account,
		ProjectId:   req.ProjectId,
		MaskedPan:   req.MaskedPan,
		RecurringId: req.RecurringId,
		Expire:      req.Expire,
		IsActive:    true,
		CreatedAt:   ptypes.TimestampNow(),
	}

	err = r.Database.Collection(CollectionSavedCard).Insert(data)

	if err != nil {
		zap.L().Error(QueryErrorMask, zap.Error(err), zap.Any("set", data))
		return err
	}

	return nil
}

func (r *Repository) DeleteSavedCard(
	ctx context.Context,
	req *repository.FindByStringValue,
	rsp *repository.Result,
) error {
	q := bson.M{"is_active": false}
	err := r.Database.Collection(CollectionSavedCard).UpdateId(bson.ObjectIdHex(req.Value), bson.M{"$set": q})

	if err != nil {
		zap.L().Error(QueryErrorMask, zap.Error(err), zap.String("id", req.Value), zap.Any("set", q))
		return err
	}

	return nil
}

func (r *Repository) FindSavedCards(
	ctx context.Context,
	req *repository.SavedCardRequest,
	rsp *repository.SavedCardList,
) error {
	var c []*entity.SavedCard

	q := bson.M{
		"account":    req.Account,
		"project_id": bson.ObjectIdHex(req.ProjectId),
		"is_active":  true,
	}
	err := r.Database.Collection(CollectionSavedCard).Find(q).All(&c)

	if err != nil {
		zap.L().Error(QueryErrorMask, zap.Error(err), zap.Any("set", q))
	}

	if len(c) > 0 {
		rsp.SavedCards = c
	}

	return nil
}

func (r *Repository) FindSavedCardById(
	ctx context.Context,
	req *repository.FindByStringValue,
	rsp *entity.SavedCard,
) error {
	var c *entity.SavedCard

	err := r.Database.Collection(CollectionSavedCard).Find(bson.M{"_id": bson.ObjectIdHex(req.Value)}).One(&c)

	if err != nil {
		zap.L().Error(QueryErrorMask, zap.Error(err), zap.Any("id", req.Value))
		return err
	}

	rsp.Id = c.Id
	rsp.ProjectId = c.ProjectId
	rsp.Expire = c.Expire
	rsp.MaskedPan = c.MaskedPan
	rsp.RecurringId = c.RecurringId
	rsp.IsActive = c.IsActive
	rsp.CreatedAt = c.CreatedAt

	return nil
}
