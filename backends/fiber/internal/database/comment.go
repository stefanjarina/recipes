package database

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

func (s *service) GetCommentsForRecipe(recipeID uuid.UUID) ([]models.Comment, error) {
	comments := []models.Comment{}
	result := s.db.Where("recipe_id = ?", recipeID).Preload("User").Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	return comments, nil
}

func (s *service) GetCommentByID(id uuid.UUID) (*models.Comment, error) {
	comment := &models.Comment{}
	result := s.db.Where("comment_id = ?", id).Preload("User").First(&comment)
	if result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

func (s *service) AddCommentToRecipe(comment *models.Comment) error {
	result := s.db.Save(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteComment(comment *models.Comment) error {
	result := s.db.Delete(comment)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
