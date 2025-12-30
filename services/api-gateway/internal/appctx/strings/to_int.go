package strings

import (
	"strconv"
)

func (s *Strings) ToInt(value string) (int, error) {
	valueNum, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return valueNum, nil
}
