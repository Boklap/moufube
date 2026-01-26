package bootstrap

import (
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/interface/controller/user"
)

type Controller struct {
	UserController *user.Controller
}

func InitController(logger *logger.AppLogger, useCase *UseCase) *Controller {
	userController := user.NewController(logger, useCase.UserUseCase)

	return &Controller{
		UserController: userController,
	}
}
