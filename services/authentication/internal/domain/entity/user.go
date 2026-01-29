package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"moufube.com/m/internal/domain/valueobject"
)

type User struct {
	ID           uuid.UUID         `gorm:"primaryKey"`
	Email        valueobject.Email `gorm:"index"`
	PasswordHash valueobject.PasswordHash
	IsVerified   bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
