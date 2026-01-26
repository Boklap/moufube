package reader

import (
	"database/sql"

	"gorm.io/gorm"
)

type UserReaderImpl struct {
	db     *sql.DB
	gormDB *gorm.DB
}
