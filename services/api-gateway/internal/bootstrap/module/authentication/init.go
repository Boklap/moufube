package authentication

import (
	"github.com/gin-gonic/gin"
	v1 "moufube.com/m/internal/bootstrap/module/authentication/v1"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

func InitModule(
	rootRouterGroup *gin.RouterGroup,
	cfg *config.Config,
	logger *logger.AppLogger,
) {
	authenticationRouterGroup := rootRouterGroup.Group("authentication")

	v1.InitAuthenticationModule(authenticationRouterGroup, cfg, logger)
}
