package apperr

import (
	"errors"
	"fmt"
)

var ErrEnvNotFound = errors.New("env key not found")

func NewEnvError(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
