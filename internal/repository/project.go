package repository

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"log"
)

func (r *Repository) FindProjectBy(ctx context.Context, req *repository.FindByRequest, rsp *repository.Projects) error {
	err := r.Database.Collection(CollectionProject).Find(req.Query).Limit(int(req.Limit)).Skip(int(req.Offset)).All(&rsp)

	if err != nil {
		log.Printf(QueryErrorMask, CollectionProject, err.Error())
		return err
	}

	return nil
}
