package response

import (
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}

func (r *Response) Error(c *gin.Context, httpStatus int, message string, err error) {
	response := &errorResponse{
		Success: false,
		Message: message,
		Error:   err,
	}

	c.JSON(httpStatus, response)
}
