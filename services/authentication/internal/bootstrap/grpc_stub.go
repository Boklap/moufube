package bootstrap

import (
	"google.golang.org/grpc"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/generated/contract"
	"moufube.com/m/internal/infrastructure/grpcserver"
)

func InitGRPCServer(cfg *config.Config, controller *Controller) *grpc.Server {
	grpcServer := grpcserver.NewGRPCServer(cfg)

	contract.RegisterUserServer(grpcServer, controller.UserController)

	return grpcServer
}
