package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/module"
)

func RegisterRoutes(rootRoute *gin.RouterGroup, mod *module.Module) {
	v1Route := rootRoute.Group("v1")

	RegisterHealthRoutes(v1Route, mod.Health)
	registerAuthenticationRoutes(v1Route, mod.Authentication)
}
