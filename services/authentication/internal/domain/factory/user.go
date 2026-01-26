package factory

import (
	"moufube.com/m/internal/domain/entity"
	"moufube.com/m/internal/domain/valueobject"
)

func NewUser(email string, password string) (*entity.User, error) {
	newEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	newPassword, err := valueobject.NewPassword(password)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := valueobject.NewPasswordHashFromPlain(newPassword.Value())
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Email:    newEmail,
		Password: hashedPassword,
	}, nil
}
