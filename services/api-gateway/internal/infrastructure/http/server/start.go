package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

func StartHTTP(httpServer *http.Server, appLogger *logger.AppLogger, cfg *config.Config) {
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			appLogger.Error(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	appLogger.Info("Shutdown signal received... cleaning up")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.ShutdownTimeout)*time.Second,
	)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		appLogger.Fatal("raceful shutdown failed", err)
	}

	appLogger.Info("Server exited properly")
}
