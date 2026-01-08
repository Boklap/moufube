package valueobject

import (
	"strings"

	"moufube.com/m/internal/domain/valueobject/voconstant"
	"moufube.com/m/internal/domain/valueobject/voerr"
)

type Password struct {
	value string
}

func NewPassword(value string) (Password, error) {
	value = normalizePassword(value)

	if err := validatePassword(value); err != nil {
		return Password{}, err
	}

	return Password{
		value: value,
	}, nil
}

func normalizePassword(value string) string {
	return strings.TrimSpace(value)
}

func validatePassword(value string) error {
	if passwordIsEmpty(value) {
		return voerr.NewPasswordError(voerr.ErrPasswordEmpty)
	}

	if len(value) < voconstant.PasswordMinLength {
		return voerr.NewPasswordError(voerr.ErrPasswordTooShort)
	}

	return nil
}

func passwordIsEmpty(value string) bool {
	return value == ""
}

func (p Password) Value() string {
	return p.value
}
