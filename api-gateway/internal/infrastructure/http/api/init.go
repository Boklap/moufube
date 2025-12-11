package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/shared/types"
)

func Init(gin *gin.Engine, cfg *types.Config) *http.Server {
	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.HttpPort),
		Handler:        gin,
		ReadTimeout:    time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:    time.Duration(cfg.IdleTimeout) * time.Second,
		MaxHeaderBytes: 1 << cfg.MaxHeaderBytes,
	}

	return srv
}
