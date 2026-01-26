package bootstrap

import (
	"os"

	"moufube.com/m/internal/bootstrap/controller"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/cache"
	"moufube.com/m/internal/infrastructure/http/gin"
	"moufube.com/m/internal/infrastructure/http/server"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/interface/router"
)

func Init() *App {
	slog := logger.InitSlog()

	cfg, err := config.Load()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	appLogger := logger.InitAppLogger(cfg)

	ginServer, err := gin.Init(cfg)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	rdb := cache.InitCacheConnection(cfg)

	repo := InitRepository(rdb)
	ctrl := controller.Init()

	router.InitHealth(ginServer, ctrl)
	initGlobalMiddleware(ginServer, cfg, repo.IdentityReader, repo.IdentityWriter)
	router.Init(ginServer, ctrl)

	httpServer := server.InitHTTP(ginServer, cfg)

	return &App{
		AppLogger:  appLogger,
		Config:     cfg,
		HTTPServer: httpServer,
	}
}
