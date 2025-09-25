package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/LekcRg/steam-inventory/internal/app"
	"go.uber.org/zap"
)

// @title           Steam inventory API
// @version         1.0
// @description     Steam inventory HTTP API
// @BasePath        /

// @securityDefinitions.apikey CookieAuth
// @in                cookie
// @name              sestoken

func main() {
	const ctxTimeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)

	application, err := app.New(ctx)
	if err != nil {
		panic(err)
	}

	defer cancel()

	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()
		exitSignals(application)
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := application.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func exitSignals(a *app.App) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	sig := <-sigChan
	a.Log.Info("Received shutdown signal", zap.String("signal", sig.String()))

	const shutdownTimeout = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

	defer cancel()
	a.Shutdown(ctx)
}
