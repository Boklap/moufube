package gin

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/apperr"
	"moufube.com/m/internal/config"
)

func Init(cfg *config.Config) (*gin.Engine, error) {
	router := gin.Default()

	err := router.SetTrustedProxies(nil)
	if err != nil {
		return nil, apperr.FailToSetTrustedProxies(err)
	}

	router.MaxMultipartMemory = cfg.MinMultipartMemory << cfg.MaxMultipartMemory

	return router, nil
}
