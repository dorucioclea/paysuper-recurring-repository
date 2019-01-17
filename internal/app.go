package internal

import (
	"github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/handlers"
	"github.com/ProtocolONE/payone-repository/internal/database"
	"github.com/ProtocolONE/payone-repository/internal/repository"
	"github.com/ProtocolONE/payone-repository/pkg/constant"
	proto "github.com/ProtocolONE/payone-repository/pkg/proto/repository"
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Host           string `envconfig:"MONGO_HOST" required:"true"`
	Database       string `envconfig:"MONGO_DB" required:"true"`
	User           string `envconfig:"MONGO_USER" required:"true"`
	Password       string `envconfig:"MONGO_PASSWORD" required:"true"`
	MetricsPort    string `envconfig:"METRICS_PORT" required:"false" default:"8085"`
	KubernetesHost string `envconfig:"KUBERNETES_SERVICE_HOST" required:"false"`
}

type Application struct {
	cfg        *Config
	service    micro.Service
	Database   *database.Source
	httpServer *http.Server
	router     *http.ServeMux
}

type appHealthCheck struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	app.initConfig()

	settings := database.Connection{
		Host:     app.cfg.Host,
		Database: app.cfg.Database,
		User:     app.cfg.User,
		Password: app.cfg.Password,
	}
	db, err := database.GetDatabase(settings)

	if err != nil {
		log.Fatalf("Database init failed with error %s\n", err)
	}

	app.Database = db

	options := []micro.Option{
		micro.Name(constant.PayOneRepositoryServiceName),
		micro.Version(constant.PayOneMicroserviceVersion),
	}

	if app.cfg.KubernetesHost == "" {
		app.service = micro.NewService(options...)
		log.Println("Initialize micro service")
	} else {
		app.service = k8s.NewService(options...)
		log.Println("Initialize k8s service")
	}

	app.service.Init()

	rep := &repository.Repository{Database: app.Database}
	err = proto.RegisterRepositoryHandler(app.service.Server(), rep)

	if err != nil {
		log.Fatalf("Repository init failed with error %s\n", err)
	}

	app.router = http.NewServeMux()
	app.initHealth()
}

func (app *Application) initConfig() {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		log.Fatalf("Config init failed with error: %s\n", err)
	}

	app.cfg = cfg
}

func (app *Application) initHealth() {
	h := health.New()
	err := h.AddChecks([]*health.Config{
		{
			Name:     "health-check",
			Checker:  &appHealthCheck{},
			Interval: time.Duration(1) * time.Second,
			Fatal:    true,
		},
	})

	if err != nil {
		log.Fatal("Health check register failed")
	}

	log.Printf("Health check listening on :%s", app.cfg.MetricsPort)

	if err = h.Start(); err != nil {
		log.Fatal("Health check start failed")
	}

	app.router.HandleFunc("/health", handlers.NewJSONHandlerFunc(h, nil))
}

func (app *Application) Run() {
	app.httpServer = &http.Server{
		Addr:    ":" + app.cfg.MetricsPort,
		Handler: app.router,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	if err := app.service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (c *appHealthCheck) Status() (interface{}, error) {
	return "ok", nil
}
