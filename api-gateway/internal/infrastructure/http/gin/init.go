package gin

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/config"
)

func Init(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		slog.Error("Fail to set trusted proxies.")
		os.Exit(1)
	}

	router.MaxMultipartMemory = cfg.MinMultipartMemory << cfg.MaxMultipartMemory

	return router
}
