package controller

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/modules/health/v1/constant"
)

type Check struct {
	response *response.Response
}

func NewCheck(
	response *response.Response,
) *Check {
	return &Check{
		response: response,
	}
}

// Execute godoc
//
//	@Summary		Instance Health Check
//	@Description	Check if the instance is healthy
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.SuccessResponse
//	@Router			/health [get]
func (ch *Check) Execute(c *gin.Context) {
	ch.response.Success(c, constant.InstanceHealthy, nil)
}
