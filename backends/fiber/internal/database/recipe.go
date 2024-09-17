package database

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

func (s *service) GetRecipes() ([]models.Recipe, error) {
	recipes := []models.Recipe{}
	result := s.db.Preload("User").Find(&recipes)
	if result.Error != nil {
		return nil, result.Error
	}
	return recipes, nil
}

func (s *service) GetRecipeByID(id uuid.UUID) (*models.Recipe, error) {
	recipe := &models.Recipe{}
	result := s.db.Where("recipe_id = ?", id).Preload("Ingredients.Measurement.Type").Preload("Ingredients.Measurement").Preload("Ingredients").Preload("User").Preload("Steps").Preload("Comments").First(recipe)
	if result.Error != nil {
		return nil, result.Error
	}
	return recipe, nil
}

func (s *service) CreateRecipe(recipe *models.Recipe) error {
	result := s.db.Create(recipe)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) UpdateRecipe(recipe *models.Recipe) error {
	result := s.db.Save(recipe)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteRecipe(id uuid.UUID) error {
	result := s.db.Where("recipe_id = ?", id).Delete(&models.Recipe{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
