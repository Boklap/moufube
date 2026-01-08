package env

import (
	"os"

	"moufube.com/m/internal/appctx/env/enverr"
)

func Get(key string) (string, error) {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return "", enverr.NewEnvErr(enverr.ErrEnvNotFound, key)
	}

	return value, nil
}
