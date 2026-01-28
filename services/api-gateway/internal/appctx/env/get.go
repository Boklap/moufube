package env

import (
	"fmt"
	"os"

	"moufube.com/m/internal/apperr"
)

func Get(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return "", apperr.NewEnvError(
			apperr.ErrEnvNotFound,
			fmt.Sprintf("key: %s", key),
		)
	}

	return value, nil
}
