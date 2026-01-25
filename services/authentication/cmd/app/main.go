package main

import (
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"moufube.com/m/internal/bootstrap"
)

func main() {
	bootstrap.InitApp()
}
