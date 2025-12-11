package env

import (
	"log/slog"
	"os"
	"strconv"
)

func StringToInt(value string) int {
	valueNum, err := strconv.Atoi(value)
	if err != nil {
		slog.Error("Cannot convert env string value to int", slog.String("value", value))
		os.Exit(1)
	}

	return valueNum
}
