package module

import (
	"github.com/redis/go-redis/v9"
	"moufube.com/m/internal/bootstrap/grpc/stub"
	"moufube.com/m/internal/bootstrap/module/authentication"
	"moufube.com/m/internal/bootstrap/module/health"
	"moufube.com/m/internal/bootstrap/module/identity"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

type Module struct {
	Authentication *authentication.Module
	Health         *health.Module
	Identity       *identity.Module
}

func BootstrapingModules(
	grpcStubs *stub.GRPCStub,
	logger *logger.AppLogger,
	rdb *redis.Client,
	cfg *config.Config,
) *Module {
	authenticationModule := authentication.BootstrapingModules(grpcStubs.Authentication, logger)
	healthModule := health.BootstrapingModules(logger)
	identityModule := identity.BootstrapingIdentityModule(rdb, cfg, logger)

	return &Module{
		Authentication: authenticationModule,
		Health:         healthModule,
		Identity:       identityModule,
	}
}
