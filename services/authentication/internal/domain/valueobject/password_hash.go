package valueobject

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"moufube.com/m/internal/domain/valueobject/voconstant"
	"moufube.com/m/internal/domain/valueobject/voerr"
)

type PasswordHash struct {
	value string
}

func NewPasswordHashFromPlain(plain string) (PasswordHash, error) {
	plain = normalizePlainPassword(plain)

	if err := validatePlainPassword(plain); err != nil {
		return PasswordHash{}, err
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return PasswordHash{}, err
	}

	return PasswordHash{
		value: string(hashBytes),
	}, nil
}

func normalizePlainPassword(plain string) string {
	return strings.TrimSpace(plain)
}

func validatePlainPassword(plain string) error {
	if passwordIsEmpty(plain) {
		return voerr.NewPasswordError(voerr.ErrPasswordEmpty)
	}

	if len(plain) < voconstant.PasswordMinLength {
		return voerr.NewPasswordError(voerr.ErrPasswordTooShort)
	}

	return nil
}

func NewPasswordHashFromHash(hash string) (*PasswordHash, error) {
	if hash == "" {
		return nil, errors.New("password hash must not be empty")
	}

	return &PasswordHash{
		value: hash,
	}, nil
}

func (p PasswordHash) Verify(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.value), []byte(plain))
	return err == nil
}

func (p PasswordHash) Value() string {
	return p.value
}
