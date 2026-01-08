package entity

import (
	"time"

	"gorm.io/gorm"
	"moufube.com/m/internal/domain/valueobject"
)

type User struct {
	ID        string            `gorm:"primaryKey"`
	Email     valueobject.Email `gorm:"index"`
	Password  valueobject.PasswordHash
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
