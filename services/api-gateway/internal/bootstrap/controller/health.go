package controller

import (
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/modules/health/controller"
)

func InitHealthController(response *response.Response) *controller.Health {
	check := controller.NewCheck(response)

	return &controller.Health{
		Check: check,
	}
}
