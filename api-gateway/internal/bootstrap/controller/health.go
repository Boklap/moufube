package controller

import (
	"moufube.com/m/internal/modules/health/controller"
)

func InitHeatlhController() *controller.Health {
	check := controller.NewCheck()

	return &controller.Health{
		Check: check,
	}
}
