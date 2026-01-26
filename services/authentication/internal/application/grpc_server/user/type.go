package user

import (
	"moufube.com/m/internal/application/port/inbound"
	"moufube.com/m/internal/domain/repository/user"
)

type GRPCServerImpl struct {
	inbound.UnimplementedUserServer

	userReader user.Reader
	userWriter user.Writer
}
