package repository

import (
	"context"
	"fmt"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
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
	if err := r.FindSavedCard(ctx, req, &billing.SavedCard{}); err == nil {
		msg := fmt.Sprintf(
			savedCardAlreadyExistMessage,
			req.Account,
			tools.ByteToObjectId(req.ProjectId).Hex(),
			tools.MaskBankCardNumber(req.Pan),
		)
		log.Printf(AlreadyExistErrorMask, msg)
		return nil
	}

	data := &billing.SavedCard{
		Id:         tools.ObjectIdToByte(bson.NewObjectId()),
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
	var c []*billing.SavedCard

	q := bson.M{"account": req.Account, "project_id": tools.ByteToObjectId(req.ProjectId), "is_active": true}
	err := r.Database.Collection(CollectionSavedCard).Find(q).All(&c)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
	}

	if len(c) > 0 {
		rsp.SavedCards = c
	}

	return nil
}

func (r *Repository) FindSavedCard(ctx context.Context, req *repository.SavedCardRequest, rsp *billing.SavedCard) error {
	var c *billing.SavedCard

	q := bson.M{
		"account":    req.Account,
		"project_id": tools.ByteToObjectId(req.ProjectId),
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
