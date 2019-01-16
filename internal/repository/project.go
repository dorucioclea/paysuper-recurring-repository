package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/globalsign/mgo/bson"
	"log"
	"time"
)

func (r *Repository) InsertProject(ctx context.Context, req *billing.Project, rsp *repository.Result) error {
	err := r.Database.Collection(CollectionProject).Insert(r.toMap(req))

	if err != nil {
		log.Printf(QueryErrorMask, CollectionProject, err.Error())
		return err
	}

	return nil
}

func (r *Repository) UpdateProject(ctx context.Context, req *billing.Project, rsp *repository.Result) error {
	m := r.toMap(req)

	id := m[FieldNameUnderscoreId]
	delete(m, FieldNameUnderscoreId)

	err := r.Database.Collection(CollectionProject).UpdateId(id, bson.M{"$set": m})

	if err != nil {
		log.Printf(QueryErrorMask, CollectionProject, err.Error())
		return err
	}

	return nil
}

func (r *Repository) FindProjectById(ctx context.Context, req *repository.FindByUnderscoreId, rsp *billing.Project) error {
	err := r.Database.Collection(CollectionProject).Find(r.toMap(req)).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionProject, err.Error())
		return err
	}

	return nil
}

func (r *Repository) ConvertProjectToProjectOrder(ctx context.Context, req *billing.Project, rsp *billing.ProjectOrder) error {
	rsp.Id = req.GetId()
	rsp.Name = req.GetName()
	rsp.UrlSuccess = req.GetUrlRedirectSuccess()
	rsp.UrlFail = req.GetUrlRedirectFail()
	rsp.SendNotifyEmail = req.GetSendNotifyEmail()
	rsp.NotifyEmails = req.GetNotifyEmails()
	rsp.SecretKey = req.GetSecretKey()
	rsp.UrlCheckAccount = req.GetUrlCheckAccount()
	rsp.UrlProcessPayment = req.GetUrlProcessPayment()
	rsp.CallbackProtocol = req.GetCallbackProtocol()
	rsp.Merchant = req.GetMerchant()

	return nil
}

func (r *Repository) DeleteProject(ctx context.Context, req *repository.FindByUnderscoreId, rsp *repository.Result) error {
	m := r.toMap(req)
	q := bson.M{"$set": bson.M{"is_active": false, "updated_at": time.Now()}}

	err := r.Database.Collection(CollectionProject).UpdateId(m[FieldNameUnderscoreId], q)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionProject, err.Error())
		return err
	}

	return nil
}
