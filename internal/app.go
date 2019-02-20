package internal

import (
	"context"
	"github.com/InVisionApp/go-health"
	"github.com/InVisionApp/go-health/handlers"
	metrics "github.com/ProtocolONE/go-micro-plugins/wrapper/monitoring/prometheus"
	"github.com/ProtocolONE/paysuper-recurring-repository/internal/database"
	"github.com/ProtocolONE/paysuper-recurring-repository/internal/repository"
	"github.com/ProtocolONE/paysuper-recurring-repository/pkg/constant"
	proto "github.com/ProtocolONE/paysuper-recurring-repository/pkg/proto/repository"
	"github.com/kelseyhightower/envconfig"
	"github.com/micro/go-micro"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Host        string `envconfig:"MONGO_HOST" required:"true"`
	Database    string `envconfig:"MONGO_DB" required:"true"`
	User        string `envconfig:"MONGO_USER" required:"true"`
	Password    string `envconfig:"MONGO_PASSWORD" required:"true"`
	MetricsPort string `envconfig:"METRICS_PORT" required:"false" default:"8085"`
}

type Application struct {
	cfg        *Config
	log        *zap.Logger
	db         *database.Source
	service    micro.Service
	httpServer *http.Server
	router     *http.ServeMux
}

type appHealthCheck struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Init() {
	app.initLogger()
	app.initConfig()

	settings := database.Connection{
		Host:     app.cfg.Host,
		Database: app.cfg.Database,
		User:     app.cfg.User,
		Password: app.cfg.Password,
	}
	db, err := database.GetDatabase(settings)

	if err != nil {
		app.log.Fatal(
			"[PAYSUPER_REPOSITORY] Database init failed",
			zap.Error(err),
			zap.String("connection_string", settings.String()),
		)
	}

	app.db = db

	app.service = micro.NewService(
		micro.Name(constant.PayOneRepositoryServiceName),
		micro.Version(constant.PayOneMicroserviceVersion),
		micro.WrapHandler(metrics.NewHandlerWrapper()),
		micro.AfterStop(func() error {
			app.log.Info("Micro service stopped")
			return nil
		}),
	)
	app.log.Info("[PAYSUPER_REPOSITORY] Initialize micro service")

	app.service.Init()

	rep := &repository.Repository{Database: app.db}
	err = proto.RegisterRepositoryHandler(app.service.Server(), rep)

	if err != nil {
		app.log.Fatal("[PAYSUPER_REPOSITORY] Repository init failed", zap.Error(err))
	}

	app.router = http.NewServeMux()
	app.initHealth()
	app.initMetrics()
}

func (app *Application) initConfig() {
	cfg := &Config{}

	if err := envconfig.Process("", cfg); err != nil {
		app.log.Fatal("[PAYSUPER_REPOSITORY] Config init failed", zap.Error(err))
	}

	app.cfg = cfg
}

func (app *Application) initLogger() {
	var err error

	app.log, err = zap.NewProduction()

	if err != nil {
		log.Fatalf("[PAYSUPER_REPOSITORY] Application logger initialization failed with error: %s\n", err)
	}
	zap.ReplaceGlobals(app.log)
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
		app.log.Fatal("[PAYSUPER_REPOSITORY] Health check register failed", zap.Error(err))
	}

	app.log.Info("[PAYSUPER_REPOSITORY] Health check listening ...", zap.String("port", app.cfg.MetricsPort))

	if err = h.Start(); err != nil {
		app.log.Fatal("[PAYSUPER_REPOSITORY] Health check start failed", zap.Error(err))
	}

	app.router.HandleFunc("/health", handlers.NewJSONHandlerFunc(h, nil))
}

func (app *Application) initMetrics() {
	app.router.Handle("/metrics", promhttp.Handler())
}

func (app *Application) Run() {
	app.httpServer = &http.Server{
		Addr:    ":" + app.cfg.MetricsPort,
		Handler: app.router,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.log.Fatal("[PAYSUPER_REPOSITORY] Http server start failed", zap.Error(err))
		}
	}()

	if err := app.service.Run(); err != nil {
		log.Fatal(err)
	}
}

func (app *Application) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.httpServer.Shutdown(ctx); err != nil {
		app.log.Fatal("Http server shutdown failed", zap.Error(err))
	}
	app.log.Info("Http server stopped")

	app.db.Close()
	app.log.Info("Database connection closed")

	func() {
		if err := app.log.Sync(); err != nil {
			app.log.Fatal("Logger sync failed", zap.Error(err))
		} else {
			app.log.Info("Logger synced")
		}
	}()
}

func (c *appHealthCheck) Status() (interface{}, error) {
	return "ok", nil
}
