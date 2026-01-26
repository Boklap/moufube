package user

import repository "moufube.com/m/internal/domain/repository/user"

type UseCaseImpl struct {
	userReader repository.Reader
	userWriter repository.Writer
}
