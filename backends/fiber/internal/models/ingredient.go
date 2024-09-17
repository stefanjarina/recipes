package models

import "github.com/google/uuid"

type Ingredient struct {
	ID            uuid.UUID   `gorm:"column:ingredient_id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name          string      `gorm:"column:name;not null" json:"name"`
	Amount        float64     `gorm:"column:amount;not null" json:"amount"`
	AmountMetric  float64     `gorm:"column:amount_metric;not null" json:"amount_metric"`
	MeasurementID uuid.UUID   `gorm:"column:measurement_id" json:"-"`
	Measurement   Measurement `json:"measurement"`
	RecipeID      uuid.UUID   `gorm:"column:recipe_id" json:"-"`
}
