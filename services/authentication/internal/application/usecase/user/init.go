package user

import "moufube.com/m/internal/domain/repository/user"

func NewUserUseCaseImpl(
	userReader user.UserReader,
	userWriter user.UserWriter,
) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		userReader: userReader,
		userWriter: userWriter,
	}
}
