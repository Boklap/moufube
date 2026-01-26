package writer

import (
	"context"

	"moufube.com/m/internal/domain/entity"
)

func (uw *UserWriterImpl) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	tx := uw.gormDB.
		WithContext(ctx).
		Create(user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}
