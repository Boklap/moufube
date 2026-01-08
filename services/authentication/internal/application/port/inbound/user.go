package inbound

import "moufube.com/m/internal/application/dto"

type UserUseCase interface {
	Register(registerDTO *dto.RegisterDTO) error
}
