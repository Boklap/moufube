package bootstrap

import (
	"moufube.com/m/internal/application/usecase/user"
)

type UseCase struct {
	UserUseCase *user.UseCaseImpl
}

func InitUseCase(writer *Writer, reader *Reader) *UseCase {
	userUseCase := user.NewUseCaseImpl(reader.UserReader, writer.UserWriter)

	return &UseCase{
		UserUseCase: userUseCase,
	}
}
