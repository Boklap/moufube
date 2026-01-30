package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/module/health"
)

func RegisterHealthRoutes(rootRoute *gin.RouterGroup, healthModule *health.Module) {
	healthRoute := rootRoute.Group("health")

	healthRoute.GET("check", healthModule.Check.Controller.Check)
}
