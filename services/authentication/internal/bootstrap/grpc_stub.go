package bootstrap

import (
	"google.golang.org/grpc"
	"moufube.com/m/internal/config"
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/contract"
	"moufube.com/m/internal/infrastructure/grpcserver"
)

func InitGRPCServer(cfg *config.Config, controller *Controller) *grpc.Server {
	grpcServer := grpcserver.NewGRPCServer(cfg)

	authenticationpb.RegisterAuthenticationServer(grpcServer, controller.UserController)

	return grpcServer
}
