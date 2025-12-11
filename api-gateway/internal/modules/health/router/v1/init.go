package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/types"
)

func InitHealthRoute(rootV1 *gin.RouterGroup, healthController *types.HealthController) {
	rootHealth := rootV1.Group("health")

	InitHealthGetRoute(rootHealth, healthController)
}
