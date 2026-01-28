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
		return nil, apperr.NewInfrastructureError(err, "fail to set trusted proxies")
	}

	router.MaxMultipartMemory = cfg.MinMultipartMemory << cfg.MaxMultipartMemory

	return router, nil
}
