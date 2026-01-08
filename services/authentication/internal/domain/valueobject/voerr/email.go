package voerr

import (
	"errors"
	"fmt"
)

var ErrEmailEmpty = errors.New("email must not be empty.")
var ErrEmailContainSpcae = errors.New("email must not contain space.")
var ErrEmailInvalidFormat = errors.New("invalid email format.")

func NewEmailError(email string, err error) error {
	return fmt.Errorf("%w: email = %s", err, email)
}
