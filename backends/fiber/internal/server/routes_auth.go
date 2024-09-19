package server

import (
	"fiber/internal/auth"
	"fiber/internal/dtos"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) LoginHandler(c *fiber.Ctx) error {
	params := new(dtos.LoginRequestDto)
	if err := c.BodyParser(params); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := s.db.GetUserByEmail(params.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if !auth.CheckPasswordHash(params.Password, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	token, err := auth.CreateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	toReturn := dtos.LoginResponseDto{
		FullName: user.FullName,
		Email:    user.Email,
		Token:    token,
	}

	return c.JSON(toReturn)
}
