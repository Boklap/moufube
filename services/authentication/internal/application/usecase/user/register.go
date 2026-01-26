package user

import (
	"context"
	"errors"
	"fmt"

	"moufube.com/m/internal/application/apperr"
	"moufube.com/m/internal/application/dto/request"
	"moufube.com/m/internal/domain/factory"
	"moufube.com/m/internal/domain/repository/repoerr"
)

func (u *UseCaseImpl) Register(registerRequest *request.RegisterRequest) error {
	ctx := context.Background()

	fetchedUser, err := u.userReader.GetByEmail(ctx, registerRequest.GetEmail())
	if err != nil && errors.Is(err, repoerr.ErrUserNotFound) {
		return err
	}

	if fetchedUser != nil {
		return apperr.NewEmailErr(apperr.ErrEmailUsed, fmt.Sprintf("email=%s", registerRequest.GetEmail()))
	}

	newUser, err := factory.NewUser(registerRequest.GetEmail(), registerRequest.GetPassword())
	if err != nil {
		return err
	}

	_, err = u.userWriter.Create(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
