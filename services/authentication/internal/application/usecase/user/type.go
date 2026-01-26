package user

import (
	"moufube.com/m/internal/domain/repository/user"
)

type UseCaseImpl struct {
	userReader user.Reader
	userWriter user.Writer
}
