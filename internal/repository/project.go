package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"log"
)

func (r *Repository) FindProjectOrderById(ctx context.Context, req *repository.FindByUnderscoreId, rsp *billing.ProjectOrder) error {
	err := r.Database.Collection(CollectionProject).Find(r.toMap(req)).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionProject, err.Error())
		return err
	}

	return nil
}
