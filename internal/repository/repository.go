package repository

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-billing-server/pkg"
	mongodb "github.com/paysuper/paysuper-database-mongo"
	"github.com/paysuper/paysuper-recurring-repository/pkg/constant"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/entity"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/repository"
	"go.uber.org/zap"
	"time"
)

const (
	savedCardQueryErrorFind   = "Query to find saved card failed"
	savedCardQueryErrorCreate = "Query to create saved card failed"

	errorUnknown  = "unknown error"
	errorNotFound = "saved card for customer not found"
)

type Repository struct {
	db *mongodb.Source
}

func NewRepositoryService(db *mongodb.Source) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertSavedCard(
	ctx context.Context,
	req *repository.SavedCardRequest,
	rsp *repository.Result,
) error {
	var card *entity.SavedCard

	query := bson.M{"token": req.Token, "masked_pan": req.MaskedPan}
	err := r.db.Collection(constant.CollectionSavedCard).Find(query).One(&card)

	if err != nil && err != mgo.ErrNotFound {
		r.logError(savedCardQueryErrorFind, []interface{}{"error", err.Error(), "query", query})
		return constant.ErrDatabase
	}

	if card != nil {
		if card.RecurringId == req.RecurringId {
			return nil
		} else {
			card.RecurringId = req.RecurringId
			card.UpdatedAt = ptypes.TimestampNow()

			err = r.db.Collection(constant.CollectionSavedCard).UpdateId(bson.ObjectIdHex(card.Id), card)
		}
	} else {
		card = &entity.SavedCard{
			Id:          bson.NewObjectId().Hex(),
			Token:       req.Token,
			ProjectId:   req.ProjectId,
			MerchantId:  req.MerchantId,
			MaskedPan:   req.MaskedPan,
			CardHolder:  req.CardHolder,
			RecurringId: req.RecurringId,
			Expire:      req.Expire,
			IsActive:    true,
			CreatedAt:   ptypes.TimestampNow(),
			UpdatedAt:   ptypes.TimestampNow(),
		}

		err = r.db.Collection(constant.CollectionSavedCard).Insert(card)
	}

	if err != nil {
		r.logError(savedCardQueryErrorCreate, []interface{}{"error", err.Error(), "data", card})
		return constant.ErrDatabase
	}

	return nil
}

func (r *Repository) DeleteSavedCard(
	ctx context.Context,
	req *repository.DeleteSavedCardRequest,
	rsp *repository.DeleteSavedCardResponse,
) error {
	query := bson.M{
		"_id":   bson.ObjectIdHex(req.Id),
		"token": req.Token,
	}
	set := bson.M{
		"is_active":  false,
		"updated_at": time.Now(),
	}
	err := r.db.Collection(constant.CollectionSavedCard).Update(query, bson.M{"$set": set})

	if err != nil {
		rsp.Status = pkg.ResponseStatusSystemError
		rsp.Message = errorUnknown

		if err != mgo.ErrNotFound {
			rsp.Status = pkg.ResponseStatusNotFound
			rsp.Message = errorNotFound
			return nil
		}

		zap.L().Error(
			pkg.ErrorDatabaseQueryFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, constant.CollectionSavedCard),
			zap.Any(pkg.ErrorDatabaseFieldQuery, query),
			zap.Any(pkg.ErrorDatabaseFieldSet, set),
		)

		return nil
	}

	rsp.Status = pkg.ResponseStatusOk

	return nil
}

func (r *Repository) FindSavedCards(
	ctx context.Context,
	req *repository.SavedCardRequest,
	rsp *repository.SavedCardList,
) error {
	var cards []*entity.SavedCard

	query := bson.M{"token": req.Token, "is_active": true}
	err := r.db.Collection(constant.CollectionSavedCard).Find(query).All(&cards)

	if err != nil && err != mgo.ErrNotFound {
		r.logError(savedCardQueryErrorFind, []interface{}{"error", err.Error(), "query", query})
		return nil
	}

	if len(cards) > 0 {
		rsp.SavedCards = cards
	}

	return nil
}

func (r *Repository) FindSavedCardById(
	ctx context.Context,
	req *repository.FindByStringValue,
	rsp *entity.SavedCard,
) error {
	var card *entity.SavedCard
	err := r.db.Collection(constant.CollectionSavedCard).FindId(bson.ObjectIdHex(req.Value)).One(&card)

	if err != nil {
		if err == mgo.ErrNotFound {
			return constant.ErrNotFound
		}

		r.logError(savedCardQueryErrorFind, []interface{}{"error", err.Error(), "_id", req.Value})
		return constant.ErrDatabase
	}

	rsp.Id = card.Id
	rsp.Token = card.Token
	rsp.ProjectId = card.ProjectId
	rsp.MerchantId = card.MerchantId
	rsp.Expire = card.Expire
	rsp.MaskedPan = card.MaskedPan
	rsp.CardHolder = card.CardHolder
	rsp.RecurringId = card.RecurringId
	rsp.IsActive = card.IsActive
	rsp.CreatedAt = card.CreatedAt
	rsp.UpdatedAt = card.UpdatedAt

	return nil
}

func (r *Repository) logError(msg string, data []interface{}) {
	zap.S().Errorw(fmt.Sprintf("[PAYSUPER_REPOSITORY] %s", msg), data...)
}
