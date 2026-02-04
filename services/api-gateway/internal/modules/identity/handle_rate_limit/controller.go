package handleratelimit

import (
	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/bootstrap/middleware/constant"
	"moufube.com/m/internal/config"
	"moufube.com/m/internal/infrastructure/logger"
	"moufube.com/m/internal/modules/identity"
)

type Controller struct {
	service *Service
	logger  *logger.AppLogger
	cfg     *config.Config
}

func NewController(
	service *Service,
	logger *logger.AppLogger,
	cfg *config.Config,
) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
		cfg:     cfg,
	}
}

func (ctrl *Controller) Execute(
	c *gin.Context,
) (*Response, error) {
	identityData, exists := c.Get(constant.Identity)
	if !exists {
		return nil, identity.NewIdentityError(identity.ErrIdentityNotFound, "Identity not found in gin")
	}

	visitorIdentity, ok := identityData.(*identity.Identity)
	if !ok {
		return nil, identity.NewIdentityError(
			identity.ErrIdentityNotFound,
			"Identity data is not of type *identity.Identity",
		)
	}
	ctx := c.Request.Context()

	count, shouldBlock, err := ctrl.service.VisitorIDRateLimit(ctx, visitorIdentity.ID)
	if err != nil {
		ctrl.logger.Errorf("Rate limit error: %v", err)
		return &Response{
			Count:       count,
			ShouldBlock: false,
		}, err
	}

	if !shouldBlock && count > ctrl.cfg.RLVisitorMin {
		ctrl.logger.Warnf("Rate limit warning: visitor %s has made %d requests", visitorIdentity.ID, count)
	}

	return &Response{
		Count:       count,
		ShouldBlock: shouldBlock,
	}, nil
}
