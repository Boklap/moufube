package user

import (
	"context"
	"errors"
	"fmt"

	"moufube.com/m/internal/application/apperr"
	"moufube.com/m/internal/application/dto/request"
	"moufube.com/m/internal/application/dto/response"
	"moufube.com/m/internal/domain/factory"
	"moufube.com/m/internal/domain/repository/repoerr"
)

func (u *GRPCServerImpl) Register(
	ctx context.Context,
	registerRequest *request.RegisterRequest,
) (*response.RegisterResponse, error) {
	fetchedUser, err := u.userReader.GetByEmail(ctx, registerRequest.GetEmail())
	if err != nil && errors.Is(err, repoerr.ErrUserNotFound) {
		return nil, err
	}

	if fetchedUser != nil {
		return nil, apperr.NewEmailErr(apperr.ErrEmailUsed, fmt.Sprintf("email=%s", registerRequest.GetEmail()))
	}

	newUser, err := factory.NewUser(registerRequest.GetEmail(), registerRequest.GetPassword())
	if err != nil {
		return nil, err
	}

	_, err = u.userWriter.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &response.RegisterResponse{}, nil
}
