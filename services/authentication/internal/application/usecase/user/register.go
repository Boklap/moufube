package user

import (
	"context"
	"fmt"

	"moufube.com/m/internal/application/apperr"
	"moufube.com/m/internal/application/dto"
	"moufube.com/m/internal/domain/factory"
)

func (u *UserUseCaseImpl) Register(registerDTO *dto.RegisterDTO) error {
	ctx := context.Background()

	fetchedUser, err := u.userReader.GetByEmail(ctx, registerDTO.Email)
	if err != nil {
		return err
	}

	if fetchedUser != nil {
		return apperr.NewEmailErr(apperr.ErrEmailUsed, fmt.Sprintf("email=%s", registerDTO.Email))
	}

	newUser, err := factory.NewUser(registerDTO.Email, registerDTO.Password)
	if err != nil {
		return err
	}

	_, err = u.userWriter.Create(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}
