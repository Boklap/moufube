package strings

import (
	"strconv"
)

func (s *Strings) ToInt64(value string) (int64, error) {
	valueInt64, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return 0, err
	}

	return valueInt64, nil
}
