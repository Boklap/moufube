package health

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/controller"
	"moufube.com/m/internal/modules/health/router"
)

func InitModule(rootRouterGroup *gin.RouterGroup) {
	rootHealthRoute := rootRouterGroup.Group("health")
	healthController := controller.NewHealthController()
	router.InitHealthRoutes(rootHealthRoute, healthController)
}
