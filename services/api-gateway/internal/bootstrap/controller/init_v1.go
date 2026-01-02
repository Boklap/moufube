package controller

import (
	ctrl "moufube.com/m/internal/interface/router/v1"
	v1 "moufube.com/m/internal/modules/health/controller/v1"
)

func initV1() *ctrl.Controller {
	healthController := v1.NewHealthController()

	return &ctrl.Controller{
		Health: healthController,
	}
}
