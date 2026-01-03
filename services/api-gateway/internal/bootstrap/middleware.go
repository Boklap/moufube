package bootstrap

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/interface/middleware"
	"moufube.com/m/internal/modules/identity/repository"
)

func initGlobalMiddleware(
	ginServer *gin.Engine,
	cfg *config.Config,
	identityReader repository.IdentityReader,
	identityWriter repository.IdentityWriter,
) {
	ginServer.Use(middleware.IdentityMiddleware(cfg, identityReader, identityWriter))
}
