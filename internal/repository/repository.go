package repository

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/paysuper/paysuper-billing-server/pkg"
	"github.com/paysuper/paysuper-recurring-repository/pkg/constant"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/entity"
	"github.com/paysuper/paysuper-recurring-repository/pkg/proto/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	mongodb "gopkg.in/paysuper/paysuper-database-mongo.v1"
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
	_ *repository.Result,
) error {
	var card *entity.SavedCard

	query := bson.M{
		"token":      req.Token,
		"masked_pan": req.MaskedPan,
	}
	err := r.db.Collection(constant.CollectionSavedCard).FindOne(ctx, query).Decode(&card)

	if err != nil && err != mongo.ErrNoDocuments {
		r.logError(savedCardQueryErrorFind, []interface{}{"error", err.Error(), "query", query})
		return constant.ErrDatabase
	}

	if card != nil {
		if card.RecurringId != req.RecurringId {
			card.RecurringId = req.RecurringId
		}

		card.UpdatedAt = ptypes.TimestampNow()

		if !card.IsActive {
			card.IsActive = true
		}

		oid, _ := primitive.ObjectIDFromHex(card.Id)
		filter := bson.M{"_id": oid}
		opts := options.Replace().SetUpsert(true)
		_, err = r.db.Collection(constant.CollectionSavedCard).ReplaceOne(ctx, filter, card, opts)
	} else {
		card = &entity.SavedCard{
			Id:          primitive.NewObjectID().Hex(),
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

		_, err = r.db.Collection(constant.CollectionSavedCard).InsertOne(ctx, card)
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
	oid, _ := primitive.ObjectIDFromHex(req.Id)
	query := bson.M{
		"_id":   oid,
		"token": req.Token,
	}
	set := bson.M{
		"is_active":  false,
		"updated_at": time.Now(),
	}
	_, err := r.db.Collection(constant.CollectionSavedCard).UpdateOne(ctx, query, bson.M{"$set": set})

	if err != nil {
		rsp.Status = pkg.ResponseStatusSystemError
		rsp.Message = errorUnknown

		if err != mongo.ErrNoDocuments {
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
	query := bson.M{"token": req.Token, "is_active": true}
	cursor, err := r.db.Collection(constant.CollectionSavedCard).Find(ctx, query)

	if err != nil {
		if err != mongo.ErrNoDocuments {
			r.logError(savedCardQueryErrorFind, []interface{}{"error", err.Error(), "query", query})
		}
		return nil
	}

	var cards []*entity.SavedCard
	err = cursor.All(ctx, &cards)

	if err != nil {
		zap.L().Error(
			pkg.ErrorQueryCursorExecutionFailed,
			zap.Error(err),
			zap.String(pkg.ErrorDatabaseFieldCollection, constant.CollectionSavedCard),
			zap.Any(pkg.ErrorDatabaseFieldQuery, query),
		)
		return err
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
	oid, _ := primitive.ObjectIDFromHex(req.Value)
	filter := bson.M{"_id": oid}
	err := r.db.Collection(constant.CollectionSavedCard).FindOne(ctx, filter).Decode(&card)

	if err != nil {
		if err == mongo.ErrNoDocuments {
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
