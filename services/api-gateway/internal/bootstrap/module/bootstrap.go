package module

import (
	"moufube.com/m/internal/bootstrap/grpc/stub"
	"moufube.com/m/internal/bootstrap/module/authentication"
	"moufube.com/m/internal/bootstrap/module/health"
	"moufube.com/m/internal/infrastructure/logger"
)

type Module struct {
	Authentication *authentication.Module
	Health         *health.Module
}

func BootstrapingModules(
	grpcStubs *stub.GRPCStub,
	logger *logger.AppLogger,
) *Module {
	authenticationModule := authentication.BootstrapingModules(grpcStubs.Authentication, logger)
	healthModule := health.BootstrapingModules(logger)

	return &Module{
		Authentication: authenticationModule,
		Health:         healthModule,
	}
}
