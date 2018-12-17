package main

import (
	"github.com/ProtocolONE/payone-repository/internal/repository"
	"github.com/ProtocolONE/payone-repository/pkg/constant"
	proto "github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"log"
)

func main() {
	service := grpc.NewService(
		micro.Name("go.p1.srv.repository"),
		micro.Version(constant.PayOneMicroserviceVersion),
	)

	service.Init()

	proto.RegisterRepositoryHandler(service.Server(), new(repository.Repository))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
