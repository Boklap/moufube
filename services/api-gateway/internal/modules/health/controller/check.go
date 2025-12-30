package controller

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/modules/health/constant"
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

func (ch *Check) Execute(c *gin.Context) {
	ch.response.Success(c, constant.InstanceHealthy, nil)
}
