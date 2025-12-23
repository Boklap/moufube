package env

import (
	"os"
)

func Get(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		os.Exit(1)
	}

	return value
}
