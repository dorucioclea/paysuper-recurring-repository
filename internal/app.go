package internal

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/internal/repository"
	"github.com/ProtocolONE/payone-repository/pkg/constant"
	proto "github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"log"
)

type Application struct {
	service  micro.Service
	database *database.Source
}

func (app *Application) Init() {
	conf, err := NewConfig()

	if err != nil {
		log.Fatalln(err)
	}

	app.database, err = app.initDatabase(conf)

	if err != nil {
		log.Fatalln(err)
	}

	app.initService()
}

func (app *Application) Run() {
	if err := app.service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (app *Application) initDatabase(conf *Config) (*database.Source, error) {
	settings := database.Connection{
		Host:     conf.Host,
		Database: conf.Database,
		User:     conf.User,
		Password: conf.Password,
	}

	return database.GetDatabase(settings)
}

func (app *Application) initService() {
	app.service = grpc.NewService(
		micro.Name(constant.PayOneRepositoryServiceName),
		micro.Version(constant.PayOneMicroserviceVersion),
	)
	app.service.Init()

	proto.RegisterRepositoryHandler(app.service.Server(), new(repository.Repository))
}
