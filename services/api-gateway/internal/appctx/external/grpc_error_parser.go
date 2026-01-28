package external

import (
	"strings"
)

func ParseGRPCErrorMessage(err error) string {
	if err == nil {
		return ""
	}

	errMsg := err.Error()
	const prefix = "desc = "
	idx := strings.Index(errMsg, prefix)

	if idx == -1 {
		return errMsg
	}

	return errMsg[idx+len(prefix):]
}
