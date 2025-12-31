package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/v1/controller"
)

func InitHealthRoute(rootV1 *gin.RouterGroup, healthController *controller.Health) {
	rootHealth := rootV1.Group("health")

	InitHealthGetRoute(rootHealth, healthController)
}
