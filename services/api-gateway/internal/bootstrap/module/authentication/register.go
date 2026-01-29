package authentication

import (
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/authentication/register"
)

type RegisterFeature struct {
	Service    *register.Service
	Controller *register.Controller
}

func bootstrapingRegister(
	authenticationGRPCClient authenticationpb.AuthenticationClient,
	logger *logger.AppLogger,
) *RegisterFeature {
	grpcClient := register.NewGRPCClient(authenticationGRPCClient)
	service := register.NewService(grpcClient)
	controller := register.NewController(service, logger)

	return &RegisterFeature{
		Service:    service,
		Controller: controller,
	}
}
