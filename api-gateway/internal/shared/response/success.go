package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, message string, data any) {
	response := &Response{
		Success: true,
		Message: message,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
