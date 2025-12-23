package env

import (
	"os"
	"strconv"
)

func (e *Env) StringToInt(value string) int {
	valueNum, err := strconv.Atoi(value)
	if err != nil {
		e.slog.Error("")
		os.Exit(1)
	}

	return valueNum
}
