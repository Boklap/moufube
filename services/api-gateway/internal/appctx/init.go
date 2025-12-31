package appctx

import (
	"moufube.com/m/internal/appctx/env"
	"moufube.com/m/internal/appctx/response"
	"moufube.com/m/internal/appctx/strings"
)

func Init() *AppCtx {
	return &AppCtx{
		Env:      env.Init(),
		Response: response.Init(),
		Strings:  strings.Init(),
	}
}
