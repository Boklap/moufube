package register

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

// Register godoc
//
//	@Summary		User Registration
//	@Description	Register a new user account with email and password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	register.Request	true	"Registration data"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/api/v1/authentication/register [post]
func (ctrl *Controller) Register(c *gin.Context) {
	var req Request

	if err := c.ShouldBind(&req); err != nil {
		ctrl.logger.Error(err)
		response.Error(c, http.StatusBadRequest, "Please check your field again", err)
		return
	}

	res, err := ctrl.service.Register(c.Request.Context(), &req)
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

	response.Success(c, "Register Success", res)
}
