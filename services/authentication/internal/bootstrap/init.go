package bootstrap

import (
	"os"

	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/database"
	"moufube.com/m/internal/infrastructure/logger"
)

func InitApp() {
	slog := logger.InitSlog()

	cfg, err := config.Load()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	appLogger := logger.InitAppLogger(cfg)

	db, err := database.InitConnection(cfg)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	gormDB, err := database.InitOrm(db)
	if err != nil {
		appLogger.Fatal(err.Error())
	}

	writer := InitWriter(db, gormDB)
	reader := InitReader(db, gormDB)
	_ = InitUseCase(writer, reader)
}
