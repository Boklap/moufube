package reader

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"moufube.com/m/internal/domain/entity"
	"moufube.com/m/internal/domain/repository/repoerr"
)

func (ur *UserReaderImpl) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	tx := ur.gormDB.
		WithContext(ctx).
		Where("email = ?", email).
		First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, repoerr.NewRepoErr(
				repoerr.ErrUserNotFound,
				fmt.Sprintf("User with email %s not found.", email),
			)
		}

		return nil, tx.Error
	}

	return &user, nil
}
