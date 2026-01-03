package apperr

import (
	"errors"
	"fmt"
)

var ErrIdentityNotFound = errors.New("identity not found")

func IdentityNotFound(key string) error {
	return fmt.Errorf("%w: key=%s", ErrIdentityNotFound, key)
}
