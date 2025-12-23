package env

import (
	"os"
	"strconv"
)

func StringToInt(value string) int {
	valueNum, err := strconv.Atoi(value)
	if err != nil {
		os.Exit(1)
	}

	return valueNum
}
