package module

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/module/authentication"
	"moufube.com/m/internal/bootstrap/module/health"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

func InitModule(
	rootRouterGroup *gin.RouterGroup,
	cfg *config.Config,
	logger *logger.AppLogger,
) {
	authentication.InitModule(rootRouterGroup, cfg, logger)
	health.InitModule(rootRouterGroup)
}
