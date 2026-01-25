package user

import (
	"context"

	"moufube.com/m/internal/domain/entity"
)

type Writer interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
