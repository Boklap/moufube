package bootstrap

import (
	"google.golang.org/grpc"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
)

type App struct {
	AppLogger  *logger.AppLogger
	GRPCServer *grpc.Server
	Cfg        *config.Config
}
