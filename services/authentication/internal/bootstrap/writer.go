package bootstrap

import (
	"database/sql"

	"gorm.io/gorm"
	"moufube.com/m/internal/domain/repository/user"
	"moufube.com/m/internal/infrastructure/repository/user/writer"
)

type Writer struct {
	UserWriter user.UserWriter
}

func InitWriter(db *sql.DB, gormDB *gorm.DB) *Writer {
	userWriter := writer.NewUserWriterImpl(db, gormDB)

	return &Writer{
		UserWriter: userWriter,
	}
}
