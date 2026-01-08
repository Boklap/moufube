package user

import repository "moufube.com/m/internal/domain/repository/user"

func NewUserUseCaseImpl(
	userReader repository.UserReader,
	userWriter repository.UserWriter,
) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		userReader: userReader,
		userWriter: userWriter,
	}
}
