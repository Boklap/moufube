package mockup

import (
	"context"
	"fmt"

	"moufube.com/m/internal/domain/entity"
	"moufube.com/m/internal/domain/factory"
	"moufube.com/m/internal/domain/repository/repoerr"
)

type ReaderMockup struct {
}

func NewReaderMockup() *ReaderMockup {
	return &ReaderMockup{}
}

func (r *ReaderMockup) GetByEmail(_ context.Context, email string) (*entity.User, error) {
	mockupUser, _ := factory.NewUser("marvino.test@email.com", "Marvin0.")

	if email == mockupUser.Email.String() {
		return mockupUser, nil
	}

	return nil, repoerr.NewRepoErr(
		repoerr.ErrUserNotFound,
		fmt.Sprintf("User with email: %s, not found", email),
	)
}
