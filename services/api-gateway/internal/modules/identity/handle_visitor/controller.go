package handlevisitor

import (
	"errors"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/identity"
)

type Controller struct {
	service *Service
	logger  *logger.AppLogger
}

func NewController(service *Service, logger *logger.AppLogger) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
	}
}

func (ctrl *Controller) Execute(
	c *gin.Context,
	req *Request,
) (*Response, error) {
	ctx := c.Request.Context()
	if req.VisitorID == "" {
		idnty, err := ctrl.service.CreateNewVisitor(ctx)
		return &Response{
			Identity: idnty,
		}, err
	}

	identityData, err := ctrl.service.GetIdentityByID(ctx, req.VisitorID)

	if errors.Is(err, identity.ErrIdentityNotFound) {
		var newIdentity *identity.Identity

		newIdentity, err = ctrl.service.CreateNewVisitor(c)
		if err != nil {
			return &Response{
				Identity: nil,
			}, err
		}

		ctrl.logger.Warnf("Client with IP: %s, is creating their own VisitorID", c.ClientIP())
		identityData = newIdentity
	}

	return &Response{
		Identity: identityData,
	}, nil
}
