package dtos

import "github.com/google/uuid"

type CreateStepDto struct {
	Title        string    `gorm:"column:title;not null" json:"title"`
	Instructions string    `gorm:"column:instructions;not null" json:"instructions"`
	StepNumber   int32     `gorm:"column:step_number;not null" json:"step_number"`
	RecipeID     uuid.UUID `gorm:"column:recipe_id" json:"-"`
}
