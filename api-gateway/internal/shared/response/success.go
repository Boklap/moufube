package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/shared/types"
)

func Success(c *gin.Context, message string, data any) {
	response := &types.Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
