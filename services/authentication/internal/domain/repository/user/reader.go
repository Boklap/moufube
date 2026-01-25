package user

import (
	"context"

	"moufube.com/m/internal/domain/entity"
)

type Reader interface {
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
