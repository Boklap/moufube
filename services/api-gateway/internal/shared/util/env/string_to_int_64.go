package env

import (
	"log/slog"
	"os"
	"strconv"
)

func StringToInt64(value string) int64 {
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		slog.Error("Cannot convert env string value to int64", slog.String("value", value))
		os.Exit(1)
	}

	return valueInt64
}
