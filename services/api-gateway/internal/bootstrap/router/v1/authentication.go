package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/module/authentication"
)

func registerAuthenticationRoutes(
	rootRoute *gin.RouterGroup,
	authenticationModule *authentication.Module,
) {
	authenticationRoute := rootRoute.Group("authentication")

	authenticationRoute.POST("register", authenticationModule.Register.Controller.Register)
	authenticationRoute.POST("login", authenticationModule.Login.Controller.Login)
}
