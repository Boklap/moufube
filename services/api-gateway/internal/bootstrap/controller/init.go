package controller

import "moufube.com/m/internal/interface/router"

func Init() *router.Controller {
	v1Ctrl := initV1()

	return &router.Controller{
		V1: v1Ctrl,
	}
}
