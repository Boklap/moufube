package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/controller"
)

func initHealthGetRoute(rootHealth *gin.RouterGroup, healthController *controller.HealthController) {
	rootHealth.GET("", healthController.Check)
}
