package apperr

import (
	"errors"
	"fmt"
)

var ErrEmailUsed = errors.New("email is already used")
var ErrInvalidCredentials = errors.New("invalid credentials")

func NewEmailErr(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
