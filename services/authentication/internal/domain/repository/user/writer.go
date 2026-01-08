package user

import (
	"context"

	"moufube.com/m/internal/domain/entity"
)

type UserWriter interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
}
