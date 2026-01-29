package stub

import (
	"moufube.com/m/internal/config"
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1"
	"moufube.com/m/internal/infrastructure/grpc/stub"
	"moufube.com/m/internal/infrastructure/logger"
)

type GRPCStub struct {
	Authentication authenticationpb.AuthenticationClient
}

func CreateGRPCStubs(
	cfg *config.Config,
	logger *logger.AppLogger,
) *GRPCStub {
	authenticationGRPCStub, _, err := stub.InitAuthenticationStub(cfg)
	if err != nil {
		logger.Error(err)
	}

	return &GRPCStub{
		Authentication: authenticationGRPCStub,
	}
}
