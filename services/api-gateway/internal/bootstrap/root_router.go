package bootstrap

import "github.com/gin-gonic/gin"

func InitRootRouter(gin *gin.Engine) *gin.RouterGroup {
	root := gin.Group("api")

	return root
}
