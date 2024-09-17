package models

import (
	"time"

	"github.com/google/uuid"
)

type Step struct {
	ID              uuid.UUID        `gorm:"column:step_id;primaryKey;default:gen_random_uuid()" json:"id"`
	Title           string           `gorm:"column:title;not null" json:"title"`
	Instructions    string           `gorm:"column:instructions;not null" json:"instructions"`
	StepNumber      int32            `gorm:"column:step_number;not null" json:"step_number"`
	RecipeID        uuid.UUID        `gorm:"column:recipe_id" json:"-"`
	StepIngredients []StepIngredient `gorm:"foreignKey:StepID" json:"step_ingredients,omitempty"`
	CreatedAt       time.Time        `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time        `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
