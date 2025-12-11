package controller

import (
	"moufube.com/m/internal/modules/health/controller"
	"moufube.com/m/internal/modules/health/types"
)

func InitHeatlhController() *types.HealthController {
	check := controller.NewCheck()

	return &types.HealthController{
		Check: check,
	}
}
