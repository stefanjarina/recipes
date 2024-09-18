package database

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

func (s *service) GetMeasurementTypes() ([]models.MeasurementType, error) {
	measurementTypes := []models.MeasurementType{}
	result := s.db.Find(&measurementTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return measurementTypes, nil
}

func (s *service) CreateMeasurementType(measurementType *models.MeasurementType) error {
	result := s.db.Create(measurementType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteMeasurementType(id uuid.UUID) error {
	result := s.db.Where("type_id = ?", id).Delete(&models.MeasurementType{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) GetMeasurements() ([]models.Measurement, error) {
	measurements := []models.Measurement{}
	result := s.db.Preload("Type").Find(&measurements)
	if result.Error != nil {
		return nil, result.Error
	}
	return measurements, nil
}

func (s *service) CreateMeasurement(measurement *models.Measurement) error {
	measurementType := &models.MeasurementType{}
	result := s.db.Where("type_id = ?", measurement.TypeID).First(measurementType)
	if result.Error != nil {
		return result.Error
	}

	measurement.Type = *measurementType

	result = s.db.Create(measurement)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteMeasurement(id uuid.UUID) error {
	result := s.db.Where("measurement_id = ?", id).Delete(&models.Measurement{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
