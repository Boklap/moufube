package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}

func (r *Response) Error(c *gin.Context, httpStatus int, message string, err error) {
	response := &ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	}

	c.JSON(httpStatus, response)
}
