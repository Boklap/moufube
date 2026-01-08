package bootstrap

import (
	"moufube.com/m/internal/application/port/inbound"
	"moufube.com/m/internal/application/usecase/user"
)

type UseCase struct {
	UserUseCase inbound.UserUseCase
}

func InitUseCase(writer *Writer, reader *Reader) *UseCase {
	userUseCase := user.NewUserUseCaseImpl(reader.UserReader, writer.UserWriter)

	return &UseCase{
		UserUseCase: userUseCase,
	}
}
