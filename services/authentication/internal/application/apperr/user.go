package apperr

import (
	"errors"
	"fmt"
)

var ErrEmailUsed = errors.New("email is already used")

func NewEmailErr(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
