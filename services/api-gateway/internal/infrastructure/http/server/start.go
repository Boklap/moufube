package server

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

func StartHTTP(httpServer *http.Server, appLogger *logger.AppLogger, cfg *config.Config) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			appLogger.Error(err)
		}
	}()

	<-quit

	appLogger.Info("shutdown signal received")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.ShutdownTimeout)*time.Second,
	)

	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		appLogger.Error("graceful shutdown failed", err)
		return err
	}

	appLogger.Info("http server shutdown completed")
	return nil
}
