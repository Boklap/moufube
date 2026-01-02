package router

import (
	"github.com/gin-gonic/gin"
	v1 "moufube.com/m/internal/interface/router/v1"
)

func Init(gin *gin.Engine, Controller *Controller) {
	root := gin.Group("api")

	v1.InitRoutes(root, Controller.V1)
}
