package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/config"
)

func InitHTTP(gin *gin.Engine, cfg *config.Config) *http.Server {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler:        gin,
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(cfg.IdleTimeout) * time.Second,
		MaxHeaderBytes: 1 << cfg.MaxHeaderBytes,
	}

	return srv
}
