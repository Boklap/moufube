package router

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/authentication/controller"
)

func InitAuthenticationRouter(
	authenticationRouterGroup *gin.RouterGroup,
	authenticationController *controller.Authentication,
) {
	InitAuthenticationPostRoutes(authenticationRouterGroup, authenticationController)
}
