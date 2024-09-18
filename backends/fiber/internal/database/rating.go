package database

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

func (s *service) GetRatingsForRecipe(recipeID uuid.UUID) ([]models.Rating, error) {
	ratings := []models.Rating{}
	result := s.db.Where("recipe_id = ?", recipeID).Find(&ratings)
	if result.Error != nil {
		return nil, result.Error
	}
	return ratings, nil
}

func (s *service) GetRatingByID(id uuid.UUID) (*models.Rating, error) {
	rating := &models.Rating{}
	result := s.db.Where("rating_id = ?", id).First(&rating)
	if result.Error != nil {
		return nil, result.Error
	}
	return rating, nil
}

func (s *service) AddRatingToRecipe(rating *models.Rating) error {
	result := s.db.Save(rating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteRating(rating *models.Rating) error {
	result := s.db.Delete(rating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
