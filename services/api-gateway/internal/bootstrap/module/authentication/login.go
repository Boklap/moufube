package authentication

import (
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/authentication/login"
)

type LoginFeature struct {
	Service    *login.Service
	Controller *login.Controller
}

func bootstrapingLogin(
	authenticationGRPCClient authenticationpb.AuthenticationClient,
	logger *logger.AppLogger,
) *LoginFeature {
	grpcClient := login.NewGRPCClient(authenticationGRPCClient)
	service := login.NewService(grpcClient)
	controller := login.NewController(service, logger)

	return &LoginFeature{
		Service:    service,
		Controller: controller,
	}
}
