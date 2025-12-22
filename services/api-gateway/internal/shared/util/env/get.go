package env

import (
	"log/slog"
	"os"
)

func Get(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		slog.Error("Required environment variable is missing", slog.String("key", key))
		os.Exit(1)
	}

	return value
}
