package health

import (
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/health/check"
)

type CheckFeature struct {
	Controller *check.Controller
}

func bootstrapCheck(
	_ *logger.AppLogger,
) *CheckFeature {
	controller := check.NewController()

	return &CheckFeature{
		Controller: controller,
	}
}
