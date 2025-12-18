package response

import (
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, httpStatus int, message string, err any) {
	response := &Response{
		Success: false,
		Message: message,
		Error:   err,
	}

	c.JSON(httpStatus, response)
}
