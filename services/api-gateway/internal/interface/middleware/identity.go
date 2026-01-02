package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/appctx/strings"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/interface/middleware/constant"
)

type Identity struct {
	ID              string
	IsAuthenticated bool
}

func IdentityMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		visitorID, _ := c.Cookie(constant.VisitorInfo)

		if visitorID == "" {
			var err error

			visitorID, err = strings.GenerateBase64Token(cfg.SizeIdentityToken)
			if err != nil {
				response.Abort(
					c,
					http.StatusInternalServerError,
					"Something went wrong, please try again later...",
					nil,
				)
				return
			}

			c.SetCookieData(&http.Cookie{
				Name:     constant.VisitorInfo,
				Value:    visitorID,
				Path:     "/",
				Domain:   "",
				Expires:  time.Now().Add(time.Duration(cfg.VisitorTokenExpireDays) * 24 * time.Hour),
				Secure:   false,
				HttpOnly: true,
			})
		}

		identity := &Identity{
			ID:              visitorID,
			IsAuthenticated: false,
		}

		c.Set(constant.Identity, identity)
		c.Next()
	}
}
