package server

import (
	"fiber/internal/auth"
	"fiber/internal/dtos"
	"fiber/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (s *FiberServer) getUsersHandler(c *fiber.Ctx) error {
	users, err := s.db.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func (s *FiberServer) getUser(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := s.db.GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func (s *FiberServer) createUser(c *fiber.Ctx) error {
	userToCreate := new(dtos.UserToCreateDto)
	if err := c.BodyParser(userToCreate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	passwordHash, err := auth.HashPassword(userToCreate.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := &models.User{
		ID:           uuid.New(),
		Email:        userToCreate.Email,
		FullName:     userToCreate.FullName,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.db.CreateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func (s *FiberServer) updateUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	userToUpdate := new(dtos.UserToUpdateDto)
	if err := c.BodyParser(userToUpdate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userFromDB, err := s.db.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userFromDB.FullName = userToUpdate.FullName

	if userToUpdate.Password != "" && !auth.CheckPasswordHash(userToUpdate.Password, userFromDB.PasswordHash) {
		passwordHash, err := auth.HashPassword(userToUpdate.Password)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		userFromDB.PasswordHash = passwordHash

	}

	err = s.db.UpdateUser(userFromDB)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(userFromDB)
}

func (s *FiberServer) deleteUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	userFromDB, err := s.db.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = s.db.DeleteUser(userFromDB.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted",
	})
}
