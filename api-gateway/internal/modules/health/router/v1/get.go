package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/types"
)

func InitHealthGetRoute(rootHealth *gin.RouterGroup, healthController *types.HealthController) {
	rootHealth.GET("", healthController.Check.Execute)
}
