package models

import (
	"time"

	"github.com/google/uuid"
)

type Rating struct {
	ID        uuid.UUID `gorm:"column:rating_id;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id" json:"-"`
	RecipeID  uuid.UUID `gorm:"column:recipe_id" json:"-"`
	Rating    int32     `gorm:"column:rating" json:"rating"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
