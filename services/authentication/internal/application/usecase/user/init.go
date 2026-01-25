package user

import "moufube.com/m/internal/domain/repository/user"

func NewUseCaseImpl(
	userReader user.Reader,
	userWriter user.Writer,
) *UseCaseImpl {
	return &UseCaseImpl{
		userReader: userReader,
		userWriter: userWriter,
	}
}
