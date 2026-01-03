package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/appctx/strings"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/interface/middleware/constant"
	"moufube.com/m/internal/modules/identity"
	"moufube.com/m/internal/modules/identity/apperr"
	"moufube.com/m/internal/modules/identity/repository"
)

func IdentityMiddleware(
	cfg *config.Config,
	identityReader repository.IdentityReader,
	identityWriter repository.IdentityWriter,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var idnty *identity.Identity
		var err error

		visitorID, _ := c.Cookie(constant.VisitorInfo)
		if visitorID == "" {
			idnty, err = createNewVisitor(c, cfg, identityWriter)
		} else {
			idnty, err = handleIdentifiedVisitor(visitorID, c, cfg, identityReader, identityWriter)
		}

		if err != nil {
			response.Abort(
				c,
				http.StatusUnauthorized,
				constant.IdentityUnknownMessage,
				nil,
			)
		}

		c.Set(constant.Identity, idnty)
		c.Next()
	}
}

func handleIdentifiedVisitor(
	visitorID string,
	c *gin.Context,
	cfg *config.Config,
	identityReader repository.IdentityReader,
	identityWriter repository.IdentityWriter,
) (*identity.Identity, error) {
	identityData, err := identityReader.GetIdentityByID(c.Request.Context(), visitorID)
	if errors.Is(err, apperr.ErrIdentityNotFound) {
		var newIdentity *identity.Identity

		newIdentity, err = createNewVisitor(c, cfg, identityWriter)
		if err != nil {
			return nil, err
		}

		// Logger warn that this visitor is creating their ID manually

		// Logic to check the client's IP and increment their failure counter
		identityData = *newIdentity
	}

	return &identityData, nil
}

func createNewVisitor(
	c *gin.Context,
	cfg *config.Config,
	identityWriter repository.IdentityWriter,
) (*identity.Identity, error) {
	visitorID, err := strings.GenerateBase64Token(cfg.SizeIdentityToken)
	if err != nil {
		return nil, err
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

	idnty := &identity.Identity{
		ID:              visitorID,
		IsAuthenticated: false,
	}

	err = identityWriter.SetIdentity(c.Request.Context(), visitorID, *idnty)
	if err != nil {
		return nil, err
	}

	return idnty, nil
}
