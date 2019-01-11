package internal

import (
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/internal/repository"
	"github.com/ProtocolONE/payone-repository/pkg/constant"
	proto "github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-grpc"
	"github.com/micro/go-micro"
	"github.com/oschwald/geoip2-golang"
	"log"
)

type Config struct {
	Host        string `envconfig:"MONGO_HOST" required:"true"`
	Database    string `envconfig:"MONGO_DB" required:"true"`
	User        string `envconfig:"MONGO_USER" required:"true"`
	Password    string `envconfig:"MONGO_PASSWORD" required:"true"`
	GeoIpDbPath string `envconfig:"MAXMIND_GEOIP_DB_PATH" required:"true"`
}

type Application struct {
	service     micro.Service
	Database    *database.Source
	GeoDbReader *geoip2.Reader
}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	cfg := app.initConfig()

	settings := database.Connection{Host: cfg.Host, Database: cfg.Database, User: cfg.User, Password: cfg.Password}
	db, err := database.GetDatabase(settings)

	if err != nil {
		log.Fatalf("Database init failed with error %s\n", err)
	}

	app.Database = db

	geoReader, err := geoip2.Open(cfg.GeoIpDbPath)

	if err != nil {
		log.Fatalf("Geo ip database init failed with error: %s\n", err)
	}

	app.GeoDbReader = geoReader

	app.service = grpc.NewService(
		micro.Name(constant.PayOneRepositoryServiceName),
		micro.Version(constant.PayOneMicroserviceVersion),
	)
	app.service.Init()

	rep := &repository.Repository{Database: app.Database, GeoReader: app.GeoDbReader}
	proto.RegisterRepositoryHandler(app.service.Server(), rep)
}

func (app *Application) initConfig() *Config {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalf("Config init failed with error: %s\n", err)
	}

	return cfg
}

func (app *Application) Run() {
	if err := app.service.Run(); err != nil {
		log.Fatal(err)
	}
}
