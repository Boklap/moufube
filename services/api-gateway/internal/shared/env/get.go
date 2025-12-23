package env

import (
	"os"
)

func (e *Env) Get(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		e.slog.Error("")
		os.Exit(1)
	}

	return value
}
