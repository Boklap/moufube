package controller

import (
	ctrl "moufube.com/m/internal/interface/router/v1"
	"moufube.com/m/internal/modules/health/controller"
)

func initV1() *ctrl.Controller {
	healthController := controller.NewHealthController()

	return &ctrl.Controller{
		Health: healthController,
	}
}
