package mockup

import (
	"context"
	"errors"

	"moufube.com/m/internal/domain/entity"
)

// ErrorReaderMockup is a mock reader that can be configured to return custom errors.
type ErrorReaderMockup struct {
	ReturnError error
}

func NewErrorReaderMockup(returnError error) *ErrorReaderMockup {
	return &ErrorReaderMockup{
		ReturnError: returnError,
	}
}

func (r *ErrorReaderMockup) GetByEmail(_ context.Context, _ string) (*entity.User, error) {
	if r.ReturnError != nil {
		return nil, r.ReturnError
	}
	return nil, errors.New("no error configured")
}
