package response

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/shared/types"
)

func Error(c *gin.Context, httpStatus int, message string, err any) {
	response := &types.Response{
		Success: false,
		Message: message,
		Error:   err,
	}

	c.JSON(httpStatus, response)
}
