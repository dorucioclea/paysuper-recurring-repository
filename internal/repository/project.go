package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"log"
)

func (r *Repository) FindProjectById(ctx context.Context, req *repository.FindByUnderscoreId, rsp *billing.Project) error {
	err := r.Database.Collection(CollectionProject).Find(r.toMap(req)).One(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionProject, err.Error())
		return err
	}

	return nil
}
