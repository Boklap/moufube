package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/controller"
	v1 "moufube.com/m/internal/interface/router/v1"
)

func Init(gin *gin.Engine, controller *controller.Controller) {
	root := gin.Group("api")

	v1.InitV1Routes(root, controller)
}
