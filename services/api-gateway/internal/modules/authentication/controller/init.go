package controller

import (
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/authentication/service"
)

func NewAuthenticationController(
	authenticationService *service.Authentication,
	logger *logger.AppLogger,
) *Authentication {
	return &Authentication{
		service: authenticationService,
		logger:  logger,
	}
}
