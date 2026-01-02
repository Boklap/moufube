package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "moufube.com/m/internal/modules/health/router/v1"
)

func InitRoutes(root *gin.RouterGroup, controller *Controller) {
	v1Group := root.Group("v1")

	v1Group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.InitHealthRoutes(v1Group, controller.Health)
}
