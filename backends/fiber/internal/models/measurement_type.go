package models

import "github.com/google/uuid"

type MeasurementType struct {
	ID   uuid.UUID `gorm:"column:type_id;primaryKey;default:gen_random_uuid()" json:"id"`
	Name string    `gorm:"column:name;not null" json:"name"`
}
