package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/router"
)

func InitHealth(gin *gin.Engine, controller *Controller) {
	root := gin.Group("api")

	router.InitHealthRoutes(root, controller.V1.Health)
}
