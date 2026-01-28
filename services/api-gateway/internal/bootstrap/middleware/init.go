package middleware

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/modules/identity/repository"
)

func InitGlobalMiddleware(
	ginServer *gin.Engine,
	cfg *config.Config,
	identityReader repository.IdentityReader,
	identityWriter repository.IdentityWriter,
) {
	ginServer.Use(identityMiddleware(cfg, identityReader, identityWriter))
}
