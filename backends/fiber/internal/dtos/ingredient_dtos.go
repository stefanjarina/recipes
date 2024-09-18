package dtos

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

type AddIngredientToRecipeDto struct {
	Name          string    `json:"name"`
	Amount        float64   `json:"amount"`
	AmountMetric  float64   `json:"amount_metric"`
	MeasurementID uuid.UUID `json:"measurement_id"`
}

type AddIngredientToStepDto struct {
	IngredientID  uuid.UUID `json:"ingredient_id"`
	Amount        float64   `json:"amount"`
	AmountMetric  float64   `json:"amount_metric"`
	MeasurementID uuid.UUID `json:"measurement_id"`
}

type StepIngredientDto struct {
	ID             uuid.UUID          `json:"id"`
	IngredientName string             `json:"name"`
	IngredientID   uuid.UUID          `json:"ingredient_id"`
	Amount         float64            `json:"amount"`
	AmountMetric   float64            `json:"amount_metric"`
	Measurement    models.Measurement `json:"measurement"`
}
