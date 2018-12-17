package main

import (
	"context"
	"github.com/ProtocolONE/payone-repository/pkg/constant"
	"github.com/ProtocolONE/payone-repository/pkg/proto/billing"
	"github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"log"
)

type Repository struct{}

func (r *Repository) UpdateOrder(ctx context.Context, req *billing.Order, rsp *repository.Response) error {
	log.Print("Received Repository.UpdateOrder request")
	rsp.Msg = "Update order with id " + req.Id
	return nil
}

func main() {
	service := grpc.NewService(
		micro.Name("go.p1.srv.repository"),
		micro.Version(constant.PayOneMicroserviceVersion),
	)

	service.Init()

	repository.RegisterRepositoryHandler(service.Server(), new(Repository))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
