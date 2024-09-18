package database

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

func (s *service) GetIngredientsForRecipe(recipeID uuid.UUID) ([]models.Ingredient, error) {
	ingredients := []models.Ingredient{}
	result := s.db.Where("recipe_id = ?", recipeID).Preload("Measurement.Type").Preload("Measurement").Find(&ingredients)
	if result.Error != nil {
		return nil, result.Error
	}
	return ingredients, nil
}

func (s *service) GetIngredientByID(id uuid.UUID) (*models.Ingredient, error) {
	ingredient := &models.Ingredient{}
	result := s.db.Where("ingredient_id = ?", id).First(&ingredient)
	if result.Error != nil {
		return nil, result.Error
	}
	return ingredient, nil
}

func (s *service) AddIngredientToRecipe(recipe *models.Recipe, ingredient *models.Ingredient) error {
	measurement := &models.Measurement{}
	result := s.db.Where("measurement_id = ?", ingredient.MeasurementID).First(measurement)
	if result.Error != nil {
		return result.Error
	}

	ingredient.Measurement = *measurement

	ingredient.RecipeID = recipe.ID

	result = s.db.Save(ingredient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteIngredient(ingredient *models.Ingredient) error {
	result := s.db.Delete(ingredient)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
