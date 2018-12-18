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
	Database *database.Source
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Run() {
	if err := app.service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (app *Application) InitDatabase(conf *Config) {
	var err error

	settings := database.Connection{
		Host:     conf.Host,
		Database: conf.Database,
		User:     conf.User,
		Password: conf.Password,
	}

	if app.Database, err = database.GetDatabase(settings); err != nil {
		log.Fatalln(err)
	}
}

func (app *Application) InitService() {
	app.service = grpc.NewService(
		micro.Name(constant.PayOneRepositoryServiceName),
		micro.Version(constant.PayOneMicroserviceVersion),
	)
	app.service.Init()

	rep := &repository.Repository{Database: app.Database}
	proto.RegisterRepositoryHandler(app.service.Server(), rep)
}
