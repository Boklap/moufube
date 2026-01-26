package user

import (
	"moufube.com/m/internal/application/port/inbound"
	"moufube.com/m/internal/infrastructure/logger"
)

func NewController(
	logger *logger.AppLogger,
	userUseCase inbound.UserUseCase,
) *Controller {
	return &Controller{
		logger:      logger,
		userUseCase: userUseCase,
	}
}
