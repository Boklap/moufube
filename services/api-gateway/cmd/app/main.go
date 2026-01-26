package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "moufube.com/m/documentation/api/v1"
	"moufube.com/m/internal/bootstrap"
	"moufube.com/m/internal/infrastructure/http/server"
)

func main() {
	app := bootstrap.Init()
	server.StartHTTP(app.HTTPServer, app.AppLogger, app.Config)
}
