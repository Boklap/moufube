package router

import (
	"github.com/gin-gonic/gin"
	v1 "moufube.com/m/internal/interface/router/v1"
)

func Init(gin *gin.Engine, controller *Controller) {
	root := gin.Group("api")

	v1.InitRoutes(root, controller.V1)
}
