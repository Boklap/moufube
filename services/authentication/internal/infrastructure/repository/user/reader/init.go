package reader

import (
	"database/sql"

	"gorm.io/gorm"
)

func NewUserReaderImpl(db *sql.DB, gormDB *gorm.DB) *UserReaderImpl {
	return &UserReaderImpl{
		db:     db,
		gormDB: gormDB,
	}
}
