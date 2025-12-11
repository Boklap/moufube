package gin

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/shared/types"
)

func Init(cfg *types.Config) *gin.Engine {
	router := gin.Default()

	router.SetTrustedProxies(nil)
	router.MaxMultipartMemory = cfg.MinMultipartMemory << cfg.MaxMultipartMemory

	return router
}
