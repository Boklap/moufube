package enverr

import (
	"errors"
	"fmt"
)

var ErrEnvNotFound = errors.New("env not found")

func NewEnvErr(err error, key string) error {
	return fmt.Errorf("%w: key = %s", err, key)
}
