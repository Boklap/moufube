package main

import (
	_ "github.com/joho/godotenv/autoload"
	"moufube.com/m/internal/bootstrap"
)

func main() {
	bootstrap.InitApp()
}
