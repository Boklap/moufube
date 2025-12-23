package controller

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/modules/health/constant"
	"moufube.com/m/internal/shared/response"
)

type Check struct {
}

func NewCheck() *Check {
	return &Check{}
}

func (ch *Check) Execute(c *gin.Context) {
	response.Success(c, constant.InstanceHealthy, nil)
}
