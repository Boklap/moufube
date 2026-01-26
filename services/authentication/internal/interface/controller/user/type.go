package user

import (
	"moufube.com/m/internal/application/port/inbound"
	"moufube.com/m/internal/generated/contract"
	"moufube.com/m/internal/infrastructure/logger"
)

type Controller struct {
	contract.UnimplementedUserServer

	logger      *logger.AppLogger
	userUseCase inbound.UserUseCase
}
