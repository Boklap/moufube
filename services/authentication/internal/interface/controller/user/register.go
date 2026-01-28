package user

import (
	"context"

	"moufube.com/m/internal/application/dto/command"
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/dto/register"
)

func (c *Controller) Register(
	ctx context.Context,
	request *authenticationpb.RegisterRequest,
) (*authenticationpb.RegisterResponse, error) {
	command := &command.RegisterUser{
		Email:    request.GetEmail(),
		Password: request.GetPassword(),
	}

	registerUserResult, err := c.userUseCase.Register(ctx, command)
	if err != nil {
		return nil, err
	}

	return &authenticationpb.RegisterResponse{
		Id:    registerUserResult.ID,
		Email: registerUserResult.Email,
	}, nil
}
