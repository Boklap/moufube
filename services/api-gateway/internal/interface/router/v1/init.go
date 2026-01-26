package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(root *gin.RouterGroup, _ *Controller) {
	v1Group := root.Group("v1")

	v1Group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
