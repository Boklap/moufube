package repoerr

import "fmt"

func NewRepoErr(err error, message string) error {
	return fmt.Errorf("%w: %s", err, message)
}
