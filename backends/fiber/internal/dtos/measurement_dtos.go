package dtos

import "github.com/google/uuid"

type CreateMeasurementDto struct {
	Name   string    `json:"name"`
	TypeID uuid.UUID `json:"type_id"`
	System string    `json:"system"`
}
