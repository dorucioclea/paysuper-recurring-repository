package repository

import (
	"context"
	"encoding/hex"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"log"
)

func (r *Repository) UpdateOrder(ctx context.Context, req *billing.Order, rsp *repository.Response) error {
	log.Print("Received Repository.UpdateOrder request")
	rsp.Msg = "Update order with id " + hex.EncodeToString(req.Id)
	return nil
}
