package env

import (
	"os"
	"strconv"
)

func (e *Env) StringToInt64(value string) int64 {
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		e.slog.Error("")
		os.Exit(1)
	}

	return valueInt64
}
