package reader

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"moufube.com/m/internal/domain/entity"
)

func (ur *UserReaderImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	tx := ur.gormDB.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, tx.Error
	}

	return &user, nil
}
