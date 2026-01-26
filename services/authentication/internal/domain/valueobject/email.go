package valueobject

import (
	"regexp"
	"strings"

	"moufube.com/m/internal/domain/valueobject/voerr"
)

type Email struct {
	value string
}

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func NewEmail(value string) (Email, error) {
	value = normalizeEmail(value)

	if err := validateEmail(value); err != nil {
		return Email{}, err
	}

	return Email{
		value: value,
	}, nil
}

func normalizeEmail(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ToLower(value)

	return value
}

func validateEmail(value string) error {
	if emailIsEmpty(value) {
		return voerr.NewEmailError(value, voerr.ErrEmailEmpty)
	}

	if strings.Contains(value, " ") {
		return voerr.NewEmailError(value, voerr.ErrEmailContainSpcae)
	}

	if !isValidEmailFormat(value) {
		return voerr.NewEmailError(value, voerr.ErrEmailInvalidFormat)
	}

	return nil
}

func emailIsEmpty(value string) bool {
	return value == ""
}

func isValidEmailFormat(value string) bool {
	return emailRegex.MatchString(value)
}

func (e Email) Value() string {
	return e.value
}
