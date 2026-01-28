package apperr

import (
	"errors"
	"fmt"
)

var ErrGRPCConnectionFailed = errors.New("fail to connect to GRPC server")

func NewGRPCError(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
