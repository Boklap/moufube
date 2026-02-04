package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/bootstrap/middleware/constant"
	"moufube.com/m/internal/bootstrap/module/identity"
	"moufube.com/m/internal/config"
	handlevisitor "moufube.com/m/internal/modules/identity/handle_visitor"
)

func Identity(
	cfg *config.Config,
	identityModule *identity.Module,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		visitorID, _ := c.Cookie(constant.VisitorInfo)
		req := &handlevisitor.Request{
			VisitorID: visitorID,
		}
		res, err := identityModule.HandleVisitor.Controller.Execute(c, req)
		if err != nil {
			response.Abort(
				c,
				http.StatusUnauthorized,
				constant.IdentityUnknownMessage,
				nil,
			)
		}

		c.SetCookieData(&http.Cookie{
			Name:     constant.VisitorInfo,
			Value:    res.Identity.ID,
			Path:     "/",
			Domain:   "",
			Expires:  time.Now().Add(time.Duration(cfg.VisitorTokenExpireDays) * 24 * time.Hour),
			Secure:   false,
			HttpOnly: true,
		})

		c.Set(constant.Identity, res.Identity)
		c.Next()
	}
}
