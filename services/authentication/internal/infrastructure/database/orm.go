package database

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitOrm(db *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(
		postgres.New(postgres.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	return gormDB, nil
}
