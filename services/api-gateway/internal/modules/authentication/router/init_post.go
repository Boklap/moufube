package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/authentication/controller"
)

func InitAuthenticationPostRoutes(
	rootAuthentication *gin.RouterGroup,
	controller *controller.Authentication,
) {
	rootAuthentication.POST("register", controller.Register)
}
