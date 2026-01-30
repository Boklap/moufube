package health

import (
	"moufube.com/m/internal/infrastructure/logger"
)

type Module struct {
	Check *CheckFeature
}

func BootstrapingModules(logger *logger.AppLogger) *Module {
	checkFeature := bootstrapCheck(logger)

	return &Module{
		Check: checkFeature,
	}
}
