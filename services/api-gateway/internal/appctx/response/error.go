package response

import (
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, httpStatus int, message string, err error) {
	response := &Response{
		Success: false,
		Message: message,
		Error:   err,
	}

	c.JSON(httpStatus, response)
}
