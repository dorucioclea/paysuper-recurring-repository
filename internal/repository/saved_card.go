package repository

import (
	"context"
	"fmt"
	"github.com/ProtocolONE/payone-repository/pkg/proto/entity"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/ProtocolONE/payone-repository/tools"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"github.com/pkg/errors"
	"log"
)

const (
	savedCardAlreadyExistMessage = "Saved card (account: %s; project_id: %s; pan: %s)"
)

func (r *Repository) InsertSavedCard(ctx context.Context, req *repository.SavedCardRequest, rsp *repository.Result) error {
	if err := r.FindSavedCard(ctx, req, &entity.SavedCard{}); err == nil {
		msg := fmt.Sprintf(
			savedCardAlreadyExistMessage,
			req.Account,
			req.ProjectId,
			tools.MaskBankCardNumber(req.Pan),
		)
		log.Printf(AlreadyExistErrorMask, msg)
		return nil
	}

	data := &entity.SavedCard{
		Id:         bson.NewObjectId().Hex(),
		Account:    req.Account,
		ProjectId:  req.ProjectId,
		MaskedPan:  tools.MaskBankCardNumber(req.Pan),
		Pan:        req.Pan,
		CardHolder: req.CardHolder,
		Expire:     req.Expire,
		IsActive:   true,
		CreatedAt:  ptypes.TimestampNow(),
	}

	err := r.Database.Collection(CollectionSavedCard).Insert(data)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
		return err
	}

	return nil
}

func (r *Repository) DeleteSavedCard(ctx context.Context, req *repository.FindByStringValue, rsp *repository.Result) error {
	q := bson.M{"is_active": false}
	err := r.Database.Collection(CollectionSavedCard).UpdateId(bson.ObjectIdHex(req.Value), bson.M{"$set": q})

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindSavedCards(ctx context.Context, req *repository.SavedCardRequest, rsp *repository.SavedCardList) error {
	var c []*entity.SavedCard

	q := bson.M{"account": req.Account, "project_id": req.ProjectId, "is_active": true}
	err := r.Database.Collection(CollectionSavedCard).Find(q).All(&c)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
	}

	if len(c) > 0 {
		rsp.SavedCards = c
	}

	return nil
}

func (r *Repository) FindSavedCard(ctx context.Context, req *repository.SavedCardRequest, rsp *entity.SavedCard) error {
	var c *entity.SavedCard

	q := bson.M{
		"account":    req.Account,
		"project_id": req.ProjectId,
		"pan":        req.Pan,
	}
	err := r.Database.Collection(CollectionSavedCard).Find(q).One(&c)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
	}

	if c == nil {
		return errors.New(fmt.Sprintf(NotFoundGeneralErrorMask, "Saved card"))
	}

	rsp = c

	return nil
}

func (r *Repository) FindSavedCardById(ctx context.Context, req *repository.FindByStringValue, rsp *entity.SavedCard) error {
	var c *entity.SavedCard
	err := r.Database.Collection(CollectionSavedCard).Find(bson.M{"_id": bson.ObjectIdHex(req.Value)}).One(&c)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
		return err
	}

	rsp.Id = c.Id
	rsp.ProjectId = c.ProjectId
	rsp.Pan = c.Pan
	rsp.Expire = c.Expire
	rsp.CardHolder = c.CardHolder
	rsp.MaskedPan = c.MaskedPan
	rsp.IsActive = c.IsActive
	rsp.CreatedAt = c.CreatedAt

	return nil
}
