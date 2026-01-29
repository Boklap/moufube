package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/external"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/modules/authentication/dto/request"
)

// Register godoc
//
//	@Summary		User Registration
//	@Description	Register a new user account with email and password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	request.Register	true	"Registration data"
//	@Success		201	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/api/authentication/v1/register [post]
func (a *Authentication) Register(c *gin.Context) {
	var req request.Register

	if err := c.ShouldBind(&req); err != nil {
		a.logger.Error(err)
		response.Error(c, http.StatusBadRequest, "Please check your field again", err)
		return
	}

	res, err := a.service.Register(c.Request.Context(), req)
	if err != nil {
		a.logger.Error(err)
		parsedMsg := external.ParseGRPCErrorMessage(err)
		response.Error(
			c,
			http.StatusInternalServerError,
			parsedMsg,
			err,
		)
		return
	}

	response.Success(c, "Registration successful", res)
}
