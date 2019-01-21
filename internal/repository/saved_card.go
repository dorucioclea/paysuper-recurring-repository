package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/ProtocolONE/payone-repository/tools"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"log"
)

func (r *Repository) InsertSavedCard(ctx context.Context, req *repository.SavedCardRequest, rsp *repository.Result) error {
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
	return nil
}

func (r *Repository) FindSavedCardsByAccount(ctx context.Context, req *repository.FindByStringValue, rsp *billing.SavedCard) error {
	return nil
}
