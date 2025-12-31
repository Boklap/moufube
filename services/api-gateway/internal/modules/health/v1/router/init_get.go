package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/v1/controller"
)

func InitHealthGetRoute(rootHealth *gin.RouterGroup, healthController *controller.Health) {
	rootHealth.GET("", healthController.Check.Execute)
}
