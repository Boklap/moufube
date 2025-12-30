package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type successResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (r *Response) Success(c *gin.Context, message string, data any) {
	response := &successResponse{
		Success: true,
		Message: message,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}
