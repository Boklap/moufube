package module

import (
	"moufube.com/m/internal/bootstrap/grpc/stub"
	"moufube.com/m/internal/bootstrap/module/authentication"
	"moufube.com/m/internal/infrastructure/logger"
)

type Module struct {
	AuthenticationModule *authentication.Module
}

func BootstrapingModules(
	grpcStubs *stub.GRPCStub,
	logger *logger.AppLogger,
) *Module {
	authenticationModule := authentication.BootstrapingModules(grpcStubs.Authentication, logger)

	return &Module{
		AuthenticationModule: authenticationModule,
	}
}
