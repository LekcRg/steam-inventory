package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/LekcRg/steam-inventory/internal/api/handlers"
	"github.com/LekcRg/steam-inventory/internal/api/middlewares"
	response "github.com/LekcRg/steam-inventory/internal/api/responder"
	"github.com/LekcRg/steam-inventory/internal/api/router"
	"github.com/LekcRg/steam-inventory/internal/config"
	"github.com/LekcRg/steam-inventory/internal/logger"
	"github.com/LekcRg/steam-inventory/internal/repository"
	"github.com/LekcRg/steam-inventory/internal/service"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type App struct {
	Log        *zap.Logger
	Config     *config.Config
	server     *http.Server
	repository *repository.Repo
}

func New(ctx context.Context) (*App, error) {
	cfg, err := config.LoadConfig(os.Args[1:])
	if err != nil {
		return nil, err
	}

	log, err := logger.CreateLogger(cfg)
	if err != nil {
		return nil, err
	}

	log.Info("Logger initialized")
	application := &App{
		Config: cfg,
		Log:    log,
	}
	application.printConfig()

	repo, err := repository.New(ctx, &cfg.Postgres, log)
	if err != nil {
		return nil, err
	}

	application.repository = repo

	application.server = application.createServer()

	return application, nil
}

func (a *App) printConfig() {
	const redacted = "[REDACTED]"

	cfg := *a.Config
	cfg.Postgres.Password = redacted

	a.Log.Info("Got config", zap.Any("config", cfg))
}

func (a *App) createServer() *http.Server {
	const (
		readTimeout       = 5 * time.Second
		writeTimeout      = 10 * time.Second
		readHeaderTimeout = 5 * time.Second
		idleTimeout       = 60 * time.Second
	)

	return &http.Server{
		Addr:              a.Config.Addr,
		Handler:           a.createRouter(),
		ReadTimeout:       readTimeout,
		WriteTimeout:      writeTimeout,
		ReadHeaderTimeout: readHeaderTimeout,
		IdleTimeout:       idleTimeout,
		ErrorLog:          zap.NewStdLog(a.Log),
	}
}

func (a *App) createRouter() *chi.Mux {
	svc := service.New(a.Config, a.repository)
	resp := response.New(a.Log)
	handl := handlers.New(a.Log, svc, a.Config, resp)
	middl := middlewares.New(a.Config, a.Log)

	return router.New(handl, middl)
}

func (a *App) startHTTPServer() error {
	a.Log.Info("Starting HTTP server", zap.String("HTTP address", a.server.Addr))

	err := a.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) Start() error {
	return a.startHTTPServer()
}

func (a *App) Shutdown(ctx context.Context) {
	a.Log.Info("Shutting down HTTP server...")

	err := a.server.Shutdown(ctx)
	if err != nil {
		a.Log.Warn("HTTP server shutdown error", zap.Error(err))

		if closeErr := a.server.Close(); closeErr != nil {
			a.Log.Error("HTTP server close error", zap.Error(closeErr))
		}
	} else {
		a.Log.Info("HTTP server gracefully stopped")
	}

	err = a.repository.Close()
	if err != nil {
		a.Log.Error("DB close error", zap.Error(err))
	} else {
		a.Log.Info("DB gracefully stopped")
	}

	if err = a.Log.Sync(); err != nil && !errors.Is(err, syscall.ENOTTY) {
		a.Log.Error("Log sync error", zap.Error(err))
	}
}
