package v1

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/modules/health/constant"
)

// Execute godoc
//
//	@Summary		Instance Health Check
//	@Description	Check if the instance is healthy
//	@Tags			Health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	response.SuccessResponse
//	@Router			/health [get]
func (hc *HealthController) Check(c *gin.Context) {
	response.Success(c, constant.InstanceHealthy, nil)
}
