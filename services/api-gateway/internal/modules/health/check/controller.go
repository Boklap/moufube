package check

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// Check godoc
//
//	@Summary		Instance Health Check
//	@Description	Check if the instance is healthy
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.Response
//	@Router			/health [get]
func (ctrl *Controller) Check(c *gin.Context) {
	response.Success(
		c,
		"Instance healthy",
		nil,
	)
}
