package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/authentication/controller"
	"moufube.com/m/internal/modules/authentication/grpc/stub"
	"moufube.com/m/internal/modules/authentication/router"
	"moufube.com/m/internal/modules/authentication/service"
)

func InitAuthenticationModule(
	authenticationRouterGroup *gin.RouterGroup,
	cfg *config.Config,
	logger *logger.AppLogger,
) {
	v1AuthenticationRouterGroup := authenticationRouterGroup.Group("v1")
	authenticationGRPCStub, _, err := stub.InitAuthenticationStub(cfg)
	if err != nil {
		logger.Error(err)
	}

	authenticationService := service.NewAuthenticationService(authenticationGRPCStub)
	authenticationController := controller.NewAuthenticationController(authenticationService, logger)
	router.InitAuthenticationRouter(v1AuthenticationRouterGroup, authenticationController)
}
