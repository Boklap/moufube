package apperr

import (
	"errors"
	"fmt"
)

var ErrFailToLoadConfig = errors.New("fail to load config")

func NewConfigErr(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
