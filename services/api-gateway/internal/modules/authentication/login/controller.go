package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/external"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/infrastructure/logger"
)

type Controller struct {
	service *Service
	logger  *logger.AppLogger
}

func NewController(
	service *Service,
	logger *logger.AppLogger,
) *Controller {
	return &Controller{
		service: service,
		logger:  logger,
	}
}

func (ctrl *Controller) Login(c *gin.Context) {
	var req Request

	if err := c.ShouldBind(&req); err != nil {
		ctrl.logger.Error(err)
		response.Error(c, http.StatusBadRequest, "Please check your field again", err)
		return
	}

	res, err := ctrl.service.Login(c.Request.Context(), &req)
	if err != nil {
		ctrl.logger.Error(err)
		parsedMsg := external.ParseGRPCErrorMessage(err)
		response.Error(
			c,
			http.StatusInternalServerError,
			parsedMsg,
			err,
		)
		return
	}

	response.Success(c, "Login successful", res)
}
