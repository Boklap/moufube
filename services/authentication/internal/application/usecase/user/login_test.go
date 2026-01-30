package user_test

import (
	"context"
	"errors"
	"testing"

	"moufube.com/m/internal/application/apperr"
	"moufube.com/m/internal/application/dto/command"
	"moufube.com/m/internal/application/usecase/user"
	"moufube.com/m/internal/application/usecase/user/mockup"
	"moufube.com/m/internal/domain/repository/repoerr"
)

func TestLoginEmpty(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "",
		Password:   "",
	}
	reader := mockup.NewReaderMockup()
	userUseCase := user.NewUseCaseImpl(reader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	// Expect error since user with empty email won't be found
	if err == nil {
		t.Error("TestLoginEmpty: expected error, got nil")
	}

	if !errors.Is(err, apperr.ErrInvalidCredentials) && !errors.Is(err, repoerr.ErrUserNotFound) {
		t.Errorf("TestLoginEmpty: expected ErrInvalidCredentials or ErrUserNotFound, got %v", err)
	}

	if result != nil {
		t.Errorf("TestLoginEmpty: expected nil result, got %v", result)
	}
}

func TestLoginEmailEmpty(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "",
		Password:   "somePassword123",
	}
	reader := mockup.NewReaderMockup()
	userUseCase := user.NewUseCaseImpl(reader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	// Expect error since user with empty email won't be found
	if err == nil {
		t.Error("TestLoginEmailEmpty: expected error, got nil")
	}

	if !errors.Is(err, apperr.ErrInvalidCredentials) && !errors.Is(err, repoerr.ErrUserNotFound) {
		t.Errorf("TestLoginEmailEmpty: expected ErrInvalidCredentials or ErrUserNotFound, got %v", err)
	}

	if result != nil {
		t.Errorf("TestLoginEmailEmpty: expected nil result, got %v", result)
	}
}

func TestLoginPasswordEmpty(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "test@email.com",
		Password:   "",
	}
	reader := mockup.NewReaderMockup()
	userUseCase := user.NewUseCaseImpl(reader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	// Expect error since user with empty password won't match
	if err == nil {
		t.Error("TestLoginPasswordEmpty: expected error, got nil")
	}

	if !errors.Is(err, apperr.ErrInvalidCredentials) && !errors.Is(err, repoerr.ErrUserNotFound) {
		t.Errorf("TestLoginPasswordEmpty: expected ErrInvalidCredentials or ErrUserNotFound, got %v", err)
	}

	if result != nil {
		t.Errorf("TestLoginPasswordEmpty: expected nil result, got %v", result)
	}
}

func TestLoginEmailNotExist(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "nonexistent@email.com",
		Password:   "password123",
	}
	reader := mockup.NewReaderMockup()
	userUseCase := user.NewUseCaseImpl(reader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	if err == nil {
		t.Error("TestLoginEmailNotExist: expected error, got nil")
	}

	if !errors.Is(err, apperr.ErrInvalidCredentials) {
		t.Errorf("TestLoginEmailNotExist: expected ErrInvalidCredentials, got %v", err)
	}

	if result != nil {
		t.Errorf("TestLoginEmailNotExist: expected nil result, got %v", result)
	}
}

func TestLoginPasswordInvalid(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "marvino.test@email.com", // Mock user's email
		Password:   "wrongPassword123",       // Wrong password
	}
	reader := mockup.NewReaderMockup()
	userUseCase := user.NewUseCaseImpl(reader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	if err == nil {
		t.Error("TestLoginPasswordInvalid: expected error, got nil")
	}

	if !errors.Is(err, apperr.ErrInvalidCredentials) {
		t.Errorf("TestLoginPasswordInvalid: expected ErrInvalidCredentials, got %v", err)
	}

	if result != nil {
		t.Errorf("TestLoginPasswordInvalid: expected nil result, got %v", result)
	}
}

func TestLoginSuccess(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "marvino.test@email.com", // Mock user's email
		Password:   "Marvin0.",               // Mock user's password
	}
	reader := mockup.NewReaderMockup()
	userUseCase := user.NewUseCaseImpl(reader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	if err != nil {
		t.Errorf("TestLoginSuccess: unexpected error: %v", err)
	}

	if result == nil {
		t.Error("TestLoginSuccess: expected result, got nil")
		return
	}

	// Verify expected fields
	if result.Message != "Login successful" {
		t.Errorf("TestLoginSuccess: expected message 'Login successful', got '%s'", result.Message)
	}

	if result.Email != "marvino.test@email.com" {
		t.Errorf("TestLoginSuccess: expected email 'marvino.test@email.com', got '%s'", result.Email)
	}

	if result.ID == "" {
		t.Error("TestLoginSuccess: expected non-empty ID")
	}
}

// TestLoginRepositoryError tests the scenario where the repository returns an unexpected error.
func TestLoginRepositoryError(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "test@email.com",
		Password:   "password123",
	}

	// Create a custom mock reader that returns a repository error
	errorReader := &mockup.ErrorReaderMockup{
		ReturnError: errors.New("database connection failed"),
	}

	userUseCase := user.NewUseCaseImpl(errorReader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	if err == nil {
		t.Error("TestLoginRepositoryError: expected error, got nil")
	}

	// The error should not be wrapped as ErrInvalidCredentials since it's a repository error
	if errors.Is(err, apperr.ErrInvalidCredentials) {
		t.Errorf("TestLoginRepositoryError: unexpected ErrInvalidCredentials, got %v", err)
	}

	if result != nil {
		t.Errorf("TestLoginRepositoryError: expected nil result, got %v", result)
	}
}

// TestLoginWithWhitespaceEmail tests login with email that contains leading/trailing whitespace.
func TestLoginWithWhitespaceEmail(t *testing.T) {
	ctx := context.Background()
	loginCommand := &command.LoginUser{
		Identifier: "  marvino.test@email.com  ", // Email with whitespace
		Password:   "Marvin0.",
	}
	reader := mockup.NewReaderMockup()
	userUseCase := user.NewUseCaseImpl(reader, nil)

	result, err := userUseCase.Login(ctx, loginCommand)

	// Should fail since the email won't match after any potential normalization
	if err == nil {
		t.Error("TestLoginWithWhitespaceEmail: expected error, got nil")
	}

	if result != nil {
		t.Errorf("TestLoginWithWhitespaceEmail: expected nil result, got %v", result)
	}
}
