package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"moufube.com/m/internal/appctx/external"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/modules/authentication/dto/request"
)

// Login godoc
//
//	@Summary		User Login
//	@Description	Login a user account with identifier and password
//	@Tags			Authentication
//	@Accept			json
//	@Produce		json
//	@Param			request	body	request.Login	true	"Login data"
//	@Success		200	{object}	response.Response
//	@Failure		400	{object}	response.Response
//	@Failure		500	{object}	response.Response
//	@Router			/api/authentication/v1/login [post]
func (a *Authentication) Login(c *gin.Context) {
	var req request.Login

	if err := c.ShouldBind(&req); err != nil {
		a.logger.Error(err)
		response.Error(c, http.StatusBadRequest, "Please check your field again", err)
		return
	}

	res, err := a.service.Login(c.Request.Context(), req)
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

	response.Success(c, "Login successful", res)
}
