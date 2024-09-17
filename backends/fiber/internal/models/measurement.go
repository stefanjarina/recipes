package models

import "github.com/google/uuid"

type Measurement struct {
	ID     uuid.UUID       `gorm:"column:measurement_id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name   string          `gorm:"column:name;not null" json:"name"`
	TypeID uuid.UUID       `gorm:"column:type_id;not null" json:"-"`
	Type   MeasurementType `json:"type"`
	System string          `gorm:"column:system;not null" json:"system"`
}
