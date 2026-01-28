package valueobject

import (
	"database/sql/driver"
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

func (e *Email) String() string {
	return e.value
}

// Scan implements the database/sql.Scanner interface.
func (e *Email) Scan(src any) error {
	if src == nil {
		return nil
	}

	switch v := src.(type) {
	case []byte:
		e.value = string(v)
	case string:
		e.value = v
	default:
		return voerr.NewEmailError("", voerr.ErrEmailInvalidFormat)
	}

	return nil
}

// Value implements the database/sql/driver.Valuer interface.
func (e *Email) Value() (driver.Value, error) {
	return e.value, nil
}
