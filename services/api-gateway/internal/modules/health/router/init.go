package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/controller"
)

func InitHealthRoutes(root *gin.RouterGroup, controller *controller.HealthController) {
	healthGroup := root.Group("health")

	initHealthGetRoute(healthGroup, controller)
}
