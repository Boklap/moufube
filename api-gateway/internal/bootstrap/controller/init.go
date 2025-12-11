package controller

import (
	"moufube.com/m/internal/modules/health/types"
)

type Controller struct {
	Health *types.HealthController
}

func Init() *Controller {
	healthController := InitHeatlhController()

	return &Controller{
		Health: healthController,
	}
}
