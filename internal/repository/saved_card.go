package repository

import (
	"context"
	"fmt"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/ProtocolONE/payone-repository/tools"
	"github.com/globalsign/mgo/bson"
	"github.com/golang/protobuf/ptypes"
	"log"
)

func (r *Repository) AddSavedCard(ctx context.Context, req *repository.AddSavedCardRequest, rsp *repository.Result) error {
	var c *billing.SavedCard

	q := bson.M{
		"account": req.Account,
		"project_id": tools.ByteToObjectId(req.ProjectId),
		"is_active": true,
	}
	err := r.Database.Collection(CollectionSavedCard).Find(q).One(&c)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
	}

	isNew := c == nil

	cm := &billing.SavedCardMasked{
		Pan: tools.MaskBankCardNumber(req.Pan),
		Expire: req.Expire,
		CreatedAt: ptypes.TimestampNow(),
	}
	c.SavedCardMasked = append(c.SavedCardMasked, cm)

	if isNew {
		c.SavedCardFull = make(map[string]*billing.SavedCardFull)
		c.Account = req.Account
		c.ProjectId = req.ProjectId
		c.CreatedAt = ptypes.TimestampNow()
		c.IsActive = true
	}

	c.SavedCardFull[fmt.Sprintf("%s:%s_%s", cm.Pan, cm.Expire.Month, cm.Expire.Year)] = &billing.SavedCardFull{
		Pan: req.Pan,
		Expire: req.Expire,
		CardHolder: req.CardHolder,
		CreatedAt: ptypes.TimestampNow(),
	}
	c.UpdatedAt = ptypes.TimestampNow()

	m := r.toMap(c)

	if isNew {
		err = r.Database.Collection(CollectionSavedCard).Insert(c)
	} else {
		id := m[FieldNameUnderscoreId]
		delete(m, FieldNameUnderscoreId)

		err = r.Database.Collection(CollectionSavedCard).UpdateId(id, bson.M{"$set": m})
	}

	if err != nil {
		log.Printf(QueryErrorMask, CollectionSavedCard, err.Error())
		return err
	}

	return nil
}

func (r *Repository) DeleteSavedCard(ctx context.Context, req *billing.SavedCardMasked, rsp *repository.Result) error {
	return nil
}

func (r *Repository) FindSavedCardsByAccount(ctx context.Context, req *repository.FindByStringValue, rsp *billing.SavedCard) error {
	return nil
}
