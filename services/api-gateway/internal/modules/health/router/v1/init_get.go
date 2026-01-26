package v1

import (
	"github.com/gin-gonic/gin"
	v1 "moufube.com/m/internal/modules/health/controller/v1"
)

func initHealthGetRoute(rootHealth *gin.RouterGroup, healthController *v1.HealthController) {
	rootHealth.GET("", healthController.Check)
}
