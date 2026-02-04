package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/middleware"
	"moufube.com/m/internal/bootstrap/module"
	v1 "moufube.com/m/internal/bootstrap/router/v1"
	"moufube.com/m/internal/config"
)

func RegisterRoutes(gin *gin.Engine, cfg *config.Config, mod *module.Module) {
	apiRoute := gin.Group("api")

	registerHealthRoutes(apiRoute, mod.Health)

	apiProtectedRoute := apiRoute.Group("")
	apiProtectedRoute.Use(middleware.Identity(cfg, mod.Identity))
	apiProtectedRoute.Use(middleware.RateLimit(mod.Identity))

	v1.RegisterRoutes(apiProtectedRoute, mod)
}
