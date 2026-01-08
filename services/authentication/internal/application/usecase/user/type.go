package user

import repository "moufube.com/m/internal/domain/repository/user"

type UserUseCaseImpl struct {
	userReader repository.UserReader
	userWriter repository.UserWriter
}
