package user

import (
	"context"
	"errors"
	"fmt"

	"moufube.com/m/internal/application/apperr"
	"moufube.com/m/internal/application/dto/command"
	"moufube.com/m/internal/application/dto/result"
	"moufube.com/m/internal/domain/factory"
	"moufube.com/m/internal/domain/repository/repoerr"
)

func (u *UseCaseImpl) Register(
	ctx context.Context,
	command *command.RegisterUser,
) (*result.RegisterUser, error) {
	fetchedUser, err := u.userReader.GetByEmail(ctx, command.Email)
	if err != nil && errors.Is(err, repoerr.ErrUserNotFound) {
		return nil, err
	}

	if fetchedUser != nil {
		return nil, apperr.NewEmailErr(apperr.ErrEmailUsed, fmt.Sprintf("email=%s", command.Email))
	}

	newUser, err := factory.NewUser(command.Email, command.Password)
	if err != nil {
		return nil, err
	}

	_, err = u.userWriter.Create(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &result.RegisterUser{
		ID:    newUser.ID,
		Email: newUser.Email.Value(),
	}, nil
}
