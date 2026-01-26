package main

import (
	"context"
	"fmt"
	"net"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"moufube.com/m/internal/bootstrap"
)

func main() {
	app := bootstrap.InitApp()
	listenConfig := &net.ListenConfig{}
	lis, err := listenConfig.Listen(
		context.TODO(),
		app.Cfg.GRPCTransportProtocol,
		fmt.Sprintf("%s:%d", app.Cfg.GRPCHost, app.Cfg.GRPCPort),
	)

	if err != nil {
		app.AppLogger.Fatal("failed to listen: " + err.Error())
	}

	if err = app.GRPCServer.Serve(lis); err != nil {
		app.AppLogger.Fatal("failed to serve: " + err.Error())
	}
}
