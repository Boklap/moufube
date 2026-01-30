package user

import (
	"context"

	"moufube.com/m/internal/application/dto/command"
	authenticationpb "moufube.com/m/internal/generated/pb/authentication/v1/login"
)

func (c *Controller) Login(
	ctx context.Context,
	request *authenticationpb.Request,
) (*authenticationpb.Response, error) {
	command := &command.LoginUser{
		Identifier: request.GetIdentifier(),
		Password:   request.GetPassword(),
	}

	loginResult, err := c.userUseCase.Login(ctx, command)
	if err != nil {
		return nil, err
	}

	return &authenticationpb.Response{
		Message: loginResult.Message,
		User: &authenticationpb.User{
			Id:         loginResult.ID,
			IsVerified: loginResult.IsVerified,
			Email:      loginResult.Email,
		},
	}, nil
}
