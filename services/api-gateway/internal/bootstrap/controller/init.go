package controller

import (
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/modules/health/v1/controller"
)

type Controller struct {
	Health *controller.Health
}

func Init(
	response *response.Response,
) *Controller {
	healthController := InitHealthController(response)

	return &Controller{
		Health: healthController,
	}
}
