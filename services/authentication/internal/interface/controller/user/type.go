package user

import (
	"moufube.com/m/internal/application/port/inbound"
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"
	"moufube.com/m/internal/infrastructure/logger"
)

type Controller struct {
	authenticationpb.UnimplementedAuthenticationServer

	logger      *logger.AppLogger
	userUseCase inbound.UserUseCase
}
