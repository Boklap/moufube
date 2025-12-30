package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"moufube.com/m/internal/config"
)

type AppLogger struct {
	*logrus.Logger
}

func InitAppLogger(config *config.Config) *AppLogger {
	if config.Environment == "dev" {
		return initDevLogrus()
	}

	return initProdLogrus()
}

func initDevLogrus() *AppLogger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &AppLogger{Logger: logger}
}

func initProdLogrus() *AppLogger {
	logger := logrus.New()

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	return &AppLogger{Logger: logger}
}
