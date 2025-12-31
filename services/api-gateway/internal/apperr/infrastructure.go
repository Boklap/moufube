package apperr

import "fmt"

func newInfrastructureError(message string, err error) *Error {
	return &Error{
		Kind:    Infrastructure,
		Message: message,
		Err:     err,
	}
}

func EnvNotFound(key string) error {
	return newInfrastructureError(
		fmt.Sprintf("Env with key %s not found: ", key),
		nil,
	)
}

func FailToLoadConfig(err error) error {
	return newInfrastructureError(
		"Failed to load config: ",
		err,
	)
}

func FailToSetTrustedProxies(err error) error {
	return newInfrastructureError(
		"Fail to set trusted proxies for gin engine: ",
		err,
	)
}
