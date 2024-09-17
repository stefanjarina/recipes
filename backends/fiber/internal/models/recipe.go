package models

import (
	"time"

	"github.com/google/uuid"
)

type Recipe struct {
	ID          uuid.UUID    `gorm:"column:recipe_id;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string       `gorm:"column:title;not null" json:"title"`
	Description string       `gorm:"column:description" json:"description"`
	Visibility  string       `gorm:"column:visibility;default:public" json:"visibility"`
	UserID      uuid.UUID    `gorm:"column:user_id" json:"-"`
	User        User         `json:"user,omitempty"`
	Ingredients []Ingredient `gorm:"foreignKey:RecipeID" json:"ingredients,omitempty"`
	Steps       []Step       `gorm:"foreignKey:RecipeID" json:"steps,omitempty"`
	Photos      []Photo      `gorm:"foreignKey:RecipeID" json:"photos,omitempty"`
	Comments    []Comment    `gorm:"foreignKey:RecipeID" json:"comments,omitempty"`
	CreatedAt   time.Time    `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
