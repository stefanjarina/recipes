package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"column:user_id;primaryKey;default:gen_random_uuid()" json:"id"`
	Email        string    `gorm:"column:email;not null" json:"email"`
	FullName     string    `gorm:"column:full_name;not null" json:"full_name"`
	PasswordHash string    `gorm:"column:password_hash;not null" json:"-"`
	Recipes      []Recipe  `gorm:"foreignKey:UserID" json:"recipes,omitempty"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
