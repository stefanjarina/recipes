package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID          uuid.UUID `gorm:"column:comment_id;primaryKey;default:gen_random_uuid()" json:"id"`
	RecipeID    uuid.UUID `gorm:"column:recipe_id" json:"-"`
	UserID      uuid.UUID `gorm:"column:user_id" json:"-"`
	User        User      `json:"user,omitempty"`
	CommentText string    `gorm:"column:comment_text;not null" json:"comment_text"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
