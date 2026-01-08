package writer

import (
	"database/sql"

	"gorm.io/gorm"
)

type UserWriterImpl struct {
	db     *sql.DB
	gormDB *gorm.DB
}
