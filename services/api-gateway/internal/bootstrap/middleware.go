package bootstrap

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/interface/middleware"
)

func initGlobalMiddleware(ginServer *gin.Engine, cfg *config.Config) {
	ginServer.Use(middleware.IdentityMiddleware(cfg))
}
