package env

import (
	"os"

	"moufube.com/m/internal/apperr"
)

func Get(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return "", apperr.EnvNotFound(key)
	}

	return value, nil
}
