package identity

import (
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/identity"
	handlevisitor "moufube.com/m/internal/modules/identity/handle_visitor"
)

type HandleVisitorFeature struct {
	Controller *handlevisitor.Controller
	Service    *handlevisitor.Service
}

func bootstrapingHandleFisitorFeature(
	reader identity.Reader,
	writer identity.Writer,
	cfg *config.Config,
	logger *logger.AppLogger,
) *HandleVisitorFeature {
	service := handlevisitor.NewService(reader, writer, cfg)
	controller := handlevisitor.NewController(service, logger)

	return &HandleVisitorFeature{
		Service:    service,
		Controller: controller,
	}
}
