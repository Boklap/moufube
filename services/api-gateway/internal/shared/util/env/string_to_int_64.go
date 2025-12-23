package env

import (
	"os"
	"strconv"
)

func StringToInt64(value string) int64 {
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		os.Exit(1)
	}

	return valueInt64
}
