package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/bootstrap/module/identity"
)

func RateLimit(
	identityMod *identity.Module,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := identityMod.HandleRateLimit.Controller.Execute(c)
		if err != nil {
			response.Abort(
				c,
				http.StatusInternalServerError,
				"Something went wrong",
				err,
			)
		}

		if res.ShouldBlock {
			response.Abort(
				c,
				http.StatusTooManyRequests,
				"Please try again later",
				nil,
			)
			return
		}

		c.Next()
	}
}
