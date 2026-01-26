package user

import (
	"context"

	"moufube.com/m/internal/application/dto/command"
	"moufube.com/m/internal/generated/dto/request"
	"moufube.com/m/internal/generated/dto/response"
)

func (c *Controller) Register(
	ctx context.Context,
	request *request.RegisterRequest,
) (*response.RegisterResponse, error) {
	command := &command.RegisterUser{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}

	registerUserResult, err := c.userUseCase.Register(ctx, command)
	if err != nil {
		return nil, err
	}

	return &response.RegisterResponse{
		Id:    registerUserResult.ID,
		Email: registerUserResult.Email,
	}, nil
}
