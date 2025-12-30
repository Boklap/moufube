package bootstrap

import (
	"net/http"

	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

type App struct {
	AppLogger  *logger.AppLogger
	Config     *config.Config
	HTTPServer *http.Server
}
