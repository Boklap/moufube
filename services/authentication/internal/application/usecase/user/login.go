package user

import (
	"context"
	"errors"

	"moufube.com/m/internal/application/apperr"
	"moufube.com/m/internal/application/dto/command"
	"moufube.com/m/internal/application/dto/result"
	"moufube.com/m/internal/domain/repository/repoerr"
)

func (u *UseCaseImpl) Login(
	ctx context.Context,
	command *command.LoginUser,
) (*result.LoginUser, error) {
	// Get user by email
	fetchedUser, err := u.userReader.GetByEmail(ctx, command.Identifier)
	if err != nil {
		if errors.Is(err, repoerr.ErrUserNotFound) {
			return nil, apperr.ErrInvalidCredentials
		}
		return nil, err
	}

	// Verify password
	if !fetchedUser.PasswordHash.Verify(command.Password) {
		return nil, apperr.ErrInvalidCredentials
	}

	// Return success result
	return &result.LoginUser{
		Message:    "Login successful",
		ID:         fetchedUser.ID.String(),
		IsVerified: fetchedUser.IsVerified,
		Email:      fetchedUser.Email.String(),
	}, nil
}
