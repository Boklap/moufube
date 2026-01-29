package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/module"
	v1 "moufube.com/m/internal/bootstrap/router/v1"
)

func RegisterRoutes(gin *gin.Engine, mod *module.Module) {
	rootRoute := gin.Group("api")

	v1.RegisterRoutes(rootRoute, mod)
}
