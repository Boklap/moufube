package bootstrap

import (
	"os"

	"moufube.com/m/internal/appctx"
	"moufube.com/m/internal/bootstrap/controller"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/http/gin"
	"moufube.com/m/internal/infrastructure/http/server"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/interface/middleware"
	"moufube.com/m/internal/interface/router"
)

func Init() *App {
	appCtx := appctx.Init()
	slog := logger.InitSlog()

	cfg, err := config.Load(appCtx.Env, appCtx.Strings)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	appLogger := logger.InitAppLogger(cfg)

	gin, err := gin.Init(cfg)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	ctrl := controller.Init(appCtx.Response)
	middleware.Init(gin)
	router.Init(gin, ctrl)
	httpServer := server.InitHTTP(gin, cfg)

	return &App{
		AppLogger:  appLogger,
		Config:     cfg,
		HTTPServer: httpServer,
	}
}
