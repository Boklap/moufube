package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/controller"
)

func InitHealthRoutes(healthRouterGroup *gin.RouterGroup, controller *controller.HealthController) {
	initHealthGetRoute(healthRouterGroup, controller)
}
