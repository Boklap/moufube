package bootstrap

import (
	"database/sql"

	"gorm.io/gorm"
	"moufube.com/m/internal/domain/repository/user"
	"moufube.com/m/internal/infrastructure/repository/user/reader"
)

type Reader struct {
	UserReader user.UserReader
}

func InitReader(db *sql.DB, gormDB *gorm.DB) *Reader {
	userReader := reader.NewUserReaderImpl(db, gormDB)

	return &Reader{
		UserReader: userReader,
	}
}
