package response

import (
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, httpStatus int, message string, err error) {
	var errMsg string
	if err != nil {
		errMsg = err.Error()
	}

	response := &Response{
		Success: false,
		Message: message,
		Error:   errMsg,
	}

	c.JSON(httpStatus, response)
}
