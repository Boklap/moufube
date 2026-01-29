package authentication

import (
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"
	"moufube.com/m/internal/infrastructure/logger"
)

type Module struct {
	Login    *LoginFeature
	Register *RegisterFeature
}

func BootstrapingModules(
	authenticationGRPCClient authenticationpb.AuthenticationClient,
	logger *logger.AppLogger,
) *Module {
	loginFeature := bootstrapingLogin(authenticationGRPCClient, logger)
	registerFeature := bootstrapingRegister(authenticationGRPCClient, logger)

	return &Module{
		Login:    loginFeature,
		Register: registerFeature,
	}
}
