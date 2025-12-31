package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"moufube.com/m/internal/bootstrap/controller"
	"moufube.com/m/internal/modules/health/v1/router"
)

//	@title			Moufube API
//	@version		1.0
//	@description	This is the API documentation for the Moufube application.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Moufube Support Center (Marvino Fransisco)
//	@contact.url	http://www.moufube.com/support
//	@contact.email	email.example@email.xoxo

//	@license.name	MIT License
//	@license.url	https://opensource.org/licenses/MIT

// @host	localhost:1000
// @BasePath	/api/v1

func InitV1Routes(root *gin.RouterGroup, controller *controller.Controller) {
	rootV1 := root.Group("v1")

	rootV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.InitHealthRoute(rootV1, controller.Health)
}
