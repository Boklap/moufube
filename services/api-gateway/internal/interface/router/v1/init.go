package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/controller"
	v1 "moufube.com/m/internal/modules/health/router/v1"
)

func InitV1Routes(root *gin.RouterGroup, controller *controller.Controller) {
	rootV1 := root.Group("v1")

	v1.InitHealthRoute(rootV1, controller.Health)
}
