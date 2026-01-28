package controller

import (
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/authentication/port/outbound"
)

type Authentication struct {
	service outbound.AuthenticationService
	logger  *logger.AppLogger
}
