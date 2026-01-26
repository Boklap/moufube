package bootstrap

import (
	"google.golang.org/grpc"
	"moufube.com/m/internal/application/port/inbound"
)

func InitGRPCServer(applicationGRPCServer *ApplicationGRPCServer) *grpc.Server {
	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	inbound.RegisterUserServer(grpcServer, applicationGRPCServer.UserGRPCServer)

	return grpcServer
}
