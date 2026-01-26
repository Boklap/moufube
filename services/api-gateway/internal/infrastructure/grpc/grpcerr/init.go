package grpcerr

import "fmt"

func NewGRPCError(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
