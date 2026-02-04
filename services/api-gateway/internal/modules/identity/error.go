package identity

import (
	"errors"
	"fmt"
)

var ErrIdentityNotFound = errors.New("identity not found")

func NewIdentityError(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
