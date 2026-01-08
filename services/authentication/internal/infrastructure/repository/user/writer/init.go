package writer

import (
	"database/sql"

	"gorm.io/gorm"
)

func NewUserWriterImpl(db *sql.DB, gormDB *gorm.DB) *UserWriterImpl {
	return &UserWriterImpl{
		db:     db,
		gormDB: gormDB,
	}
}
