package bootstrap

import (
	"os"

	"moufube.com/m/internal/bootstrap/grpc/stub"
	"moufube.com/m/internal/bootstrap/module"
	"moufube.com/m/internal/bootstrap/router"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/http/gin"
	"moufube.com/m/internal/infrastructure/http/server"
	"moufube.com/m/internal/infrastructure/logger"
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

	// rdb := cache.InitCacheConnection(cfg)
	// repo := InitRepository(rdb)
	grpcStubs := stub.CreateGRPCStubs(cfg, appLogger)
	mod := module.BootstrapingModules(grpcStubs, appLogger)

	router.RegisterRoutes(ginServer, mod)

	httpServer := server.InitHTTP(ginServer, cfg)

	return &App{
		AppLogger:  appLogger,
		Config:     cfg,
		HTTPServer: httpServer,
	}
}
