package middleware

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/module/identity"
	"moufube.com/m/internal/config"
)

func InitGlobalMiddleware(
	_ *gin.Engine,
	_ *config.Config,
	_ *identity.Module,
) {
	// _ginServer.Use(identityMiddleware(_cfg, _identityModule))
}
