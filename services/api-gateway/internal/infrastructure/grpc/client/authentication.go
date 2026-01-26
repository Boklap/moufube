package client

import (
	"fmt"

	"google.golang.org/grpc"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

func InitAuthenticationStub(appLogger *logger.AppLogger, cfg *config.Config) {
	var opts []grpc.DialOption
	serverAddress := fmt.Sprintf("%s:%d", cfg.GRPCAuthenticationHost, cfg.GRPCAuthenticationPort)

	conn, err := grpc.NewClient(serverAddress, opts...)
	if err != nil {
		appLogger.Debug()
	}
	defer conn.Close()
	// client :=
}
