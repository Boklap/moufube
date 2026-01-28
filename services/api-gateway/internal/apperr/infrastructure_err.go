package apperr

import (
	"fmt"
)

func NewInfrastructureError(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
