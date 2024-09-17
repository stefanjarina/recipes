package database

import (
	"fiber/internal/models"

	"github.com/google/uuid"
)

func (s *service) GetUsers() ([]models.User, error) {
	users := []models.User{}
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *service) GetUserByID(id uuid.UUID) (*models.User, error) {
	user := &models.User{}
	result := s.db.Where("user_id = ?", id).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (s *service) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	result := s.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (s *service) CreateUser(user *models.User) error {
	result := s.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) UpdateUser(user *models.User) error {
	result := s.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *service) DeleteUser(id uuid.UUID) error {
	result := s.db.Where("user_id = ?", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
