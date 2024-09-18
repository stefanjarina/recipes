package database

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

func (s *service) GetStepsForRecipe(recipeID uuid.UUID) ([]models.Step, error) {
	steps := []models.Step{}
	result := s.db.Where("recipe_id = ?", recipeID).Preload("StepIngredients.Measurement.Type").Preload("StepIngredients.Measurement").Preload("StepIngredients").Find(&steps)
	if result.Error != nil {
		return nil, result.Error
	}
	return steps, nil
}

func (s *service) GetStepByID(id uuid.UUID) (*models.Step, error) {
	step := &models.Step{}
	result := s.db.Where("step_id = ?", id).Preload("StepIngredients.Measurement.Type").Preload("StepIngredients.Measurement").Preload("StepIngredients").First(&step)
	if result.Error != nil {
		return nil, result.Error
	}
	return step, nil
}

func (s *service) AddStepToRecipe(recipe *models.Recipe, step *models.Step) error {
	step.RecipeID = recipe.ID

	result := s.db.Save(step)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteStep(step *models.Step) error {
	result := s.db.Delete(step)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) GetIngredientsForStep(stepID uuid.UUID) ([]models.StepIngredient, error) {
	ingredients := []models.StepIngredient{}
	result := s.db.Where("step_id = ?", stepID).Preload("Measurement.Type").Preload("Measurement").Preload("Ingredient").Find(&ingredients)
	if result.Error != nil {
		return nil, result.Error
	}
	return ingredients, nil
}

func (s *service) GetStepIngredientByID(id uuid.UUID) (*models.StepIngredient, error) {
	ingredient := &models.StepIngredient{}
	result := s.db.Where("step_ingredient_id = ?", id).Preload("Ingredient").First(&ingredient)
	if result.Error != nil {
		return nil, result.Error
	}
	return ingredient, nil
}

func (s *service) AddIngredientToStep(step *models.Step, stepIngredient *models.StepIngredient) error {
	measurement := &models.Measurement{}
	result := s.db.Where("measurement_id = ?", stepIngredient.MeasurementID).First(measurement)
	if result.Error != nil {
		return result.Error
	}

	stepIngredient.Measurement = *measurement

	stepIngredient.StepID = step.ID

	result = s.db.Save(stepIngredient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteStepIngredient(stepIngredient *models.StepIngredient) error {
	result := s.db.Delete(stepIngredient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
