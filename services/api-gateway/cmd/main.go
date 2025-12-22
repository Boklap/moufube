package main

import (
	_ "github.com/joho/godotenv/autoload"
	"moufube.com/m/internal/bootstrap/controller"
	"moufube.com/m/internal/infrastructure/config"
	"moufube.com/m/internal/infrastructure/http/gin"
	"moufube.com/m/internal/infrastructure/http/server"
	"moufube.com/m/internal/interface/middleware"
	"moufube.com/m/internal/interface/router"
)

func main() {
	cfg := config.Load()
	gin := gin.Init(cfg)
	controller := controller.Init()

	router.Init(gin, controller)
	middleware.Init(gin)

	httpServer := server.InitHTTP(gin, cfg)

	server.StartHTTP(httpServer)
}
