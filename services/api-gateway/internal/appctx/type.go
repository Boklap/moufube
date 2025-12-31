package appctx

import (
	"moufube.com/m/internal/appctx/env"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/appctx/strings"
)

type AppCtx struct {
	Strings  *strings.Strings
	Response *response.Response
	Env      *env.Env
}
