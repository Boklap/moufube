package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"moufube.com/m/internal/config"
)

func InitConnection(cfg *config.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := sql.Open(cfg.DBMS, psqlInfo)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.DBTimeout)*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
