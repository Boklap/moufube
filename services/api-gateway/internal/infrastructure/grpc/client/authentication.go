package client

import (
	"fmt"

	"google.golang.org/grpc"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/grpc/grpcerr"
	"moufube.com/m/internal/infrastructure/logger"
)

func InitAuthenticationStub(appLogger *logger.AppLogger, cfg *config.Config) {
	var opts []grpc.DialOption
	serverAddress := fmt.Sprintf("%s:%d", cfg.GRPCAuthenticationHost, cfg.GRPCAuthenticationPort)

	conn, err := grpc.NewClient(serverAddress, opts...)
	if err != nil {
		grpcErr := grpcerr.NewGRPCError(grpcerr.ErrGRPCConnectionFailed, "grpc server: authentication")
		appLogger.Error(grpcErr.Error())
		return
	}

	defer conn.Close()
	// client :=
}
