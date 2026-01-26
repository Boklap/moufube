package inbound

import "moufube.com/m/internal/application/dto/request"

type UserUseCase interface {
	Register(registerRequest *request.RegisterRequest) error
}
