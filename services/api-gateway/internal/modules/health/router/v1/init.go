package v1

import (
	"github.com/gin-gonic/gin"
	v1 "moufube.com/m/internal/modules/health/controller/v1"
)

func InitHealthRoutes(root *gin.RouterGroup, controller *v1.HealthController) {
	healthGroup := root.Group("health")

	initHealthGetRoute(healthGroup, controller)
}
