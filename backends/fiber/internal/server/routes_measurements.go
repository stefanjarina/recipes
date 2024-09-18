package server

import (
	"fiber/internal/dtos"
	"fiber/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (s *FiberServer) getMeasurementsHandler(c *fiber.Ctx) error {
	measurements, err := s.db.GetMeasurements()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(measurements)
}

func (s *FiberServer) createMeasurementHandler(c *fiber.Ctx) error {
	createMeasurementDto := new(dtos.CreateMeasurementDto)
	if err := c.BodyParser(createMeasurementDto); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	measurement := &models.Measurement{
		Name:   createMeasurementDto.Name,
		TypeID: createMeasurementDto.TypeID,
		System: createMeasurementDto.System,
	}

	err := s.db.CreateMeasurement(measurement)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(measurement)
}

func (s *FiberServer) deleteMeasurementHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = s.db.DeleteMeasurement(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (s *FiberServer) getMeasurementTypesHandler(c *fiber.Ctx) error {
	measurementTypes, err := s.db.GetMeasurementTypes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(measurementTypes)
}

func (s *FiberServer) createMeasurementTypeHandler(c *fiber.Ctx) error {
	measurementType := new(models.MeasurementType)
	if err := c.BodyParser(measurementType); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := s.db.CreateMeasurementType(measurementType)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(measurementType)
}

func (s *FiberServer) deleteMeasurementTypeHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = s.db.DeleteMeasurementType(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
