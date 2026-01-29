package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	_ "moufube.com/m/documentation/api"
	"moufube.com/m/internal/bootstrap"
	"moufube.com/m/internal/infrastructure/http/server"
)

// @title           API Gateway Service
// @version         1.0
// @description     API Gateway for microservices architecture
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

func main() {
	app := bootstrap.Init()
	server.StartHTTP(app.HTTPServer, app.AppLogger, app.Config)
}
