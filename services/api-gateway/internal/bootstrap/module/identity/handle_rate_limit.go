package identity

import (
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/identity"
	handleratelimit "moufube.com/m/internal/modules/identity/handle_rate_limit"
)

func bootstrapingHandleRateLimitFeature(
	reader identity.Reader,
	writer identity.Writer,
	cfg *config.Config,
	logger *logger.AppLogger,
) *HandleRateLimitFeature {
	service := handleratelimit.NewService(writer, reader, cfg)
	controller := handleratelimit.NewController(service, logger, cfg)
	return &HandleRateLimitFeature{
		Service:    service,
		Controller: controller,
	}
}
