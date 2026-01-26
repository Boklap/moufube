package inbound

import (
	"context"

	"moufube.com/m/internal/application/dto/command"
	"moufube.com/m/internal/application/dto/result"
)

type UserUseCase interface {
	Register(ctx context.Context, command *command.RegisterUser) (*result.RegisterUser, error)
}
