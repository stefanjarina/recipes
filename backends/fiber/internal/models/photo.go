package models

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID        uuid.UUID `gorm:"column:photo_id;primaryKey;default:gen_random_uuid()" json:"id"`
	RecipeID  uuid.UUID `gorm:"column:recipe_id" json:"-"`
	FilePath  string    `gorm:"column:file_path" json:"file_path"`
	URL       string    `gorm:"column:url" json:"url"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
