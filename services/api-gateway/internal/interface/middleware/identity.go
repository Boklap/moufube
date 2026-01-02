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
	Id              string
	IsAuthenticated bool
}

func IdentityMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		visitorId, _ := c.Cookie(constant.VisitorInfo)

		if visitorId == "" {
			var err error

			visitorId, err = strings.GenerateBase64Token(cfg.SizeIdentityToken)
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
				Value:    visitorId,
				Path:     "/",
				Domain:   "",
				Expires:  time.Now().Add(time.Duration(cfg.VisitorTokenExpireDays) * 24 * time.Hour),
				Secure:   false,
				HttpOnly: true,
			})
		}

		identity := &Identity{
			Id:              visitorId,
			IsAuthenticated: false,
		}

		c.Set(constant.Identity, identity)
		c.Next()
	}
}
