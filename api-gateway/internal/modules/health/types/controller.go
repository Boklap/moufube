package types

import "moufube.com/m/internal/modules/health/controller"

type HealthController struct {
	Check *controller.Check
}
