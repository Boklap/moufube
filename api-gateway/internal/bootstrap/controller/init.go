package controller

import "moufube.com/m/internal/modules/health/controller"

type Controller struct {
	Health *controller.Health
}

func Init() *Controller {
	healthController := InitHeatlhController()

	return &Controller{
		Health: healthController,
	}
}
