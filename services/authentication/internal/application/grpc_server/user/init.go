package user

import "moufube.com/m/internal/domain/repository/user"

func NewGRPCServerImpl(
	userReader user.Reader,
	userWriter user.Writer,
) *GRPCServerImpl {
	return &GRPCServerImpl{
		userReader: userReader,
		userWriter: userWriter,
	}
}
