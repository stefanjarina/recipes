package models

import "github.com/google/uuid"

type StepIngredient struct {
	ID            uuid.UUID   `gorm:"column:step_ingredient_id;primaryKey;default:gen_random_uuid()" json:"id"`
	StepID        uuid.UUID   `gorm:"column:step_id" json:"-"`
	IngredientID  uuid.UUID   `gorm:"column:ingredient_id" json:"-"`
	Amount        float64     `gorm:"column:amount;not null" json:"amount"`
	AmountMetric  float64     `gorm:"column:amount_metric;not null" json:"amount_metric"`
	MeasurementID uuid.UUID   `gorm:"column:measurement_id" json:"-"`
	Measurement   Measurement `json:"measurement"`
}
