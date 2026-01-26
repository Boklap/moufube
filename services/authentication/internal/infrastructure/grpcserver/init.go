package grpcserver

import (
	"google.golang.org/grpc"
	"moufube.com/m/internal/config"
)

func NewGRPCServer(_ *config.Config) *grpc.Server {
	var opts []grpc.ServerOption

	// Create TLS here

	server := grpc.NewServer(opts...)

	return server
}
