package router

import v1 "moufube.com/m/internal/interface/router/v1"

type Controller struct {
	V1 *v1.Controller
}
