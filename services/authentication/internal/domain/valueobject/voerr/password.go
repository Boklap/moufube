package voerr

import (
	"errors"
	"fmt"
)

var ErrPasswordEmpty = errors.New("password is empty.")
var ErrPasswordTooShort = errors.New("password is too short.")

func NewPasswordError(err error) error {
	return fmt.Errorf("%w", err)
}
