package server

import (
	"fiber/internal/dtos"
	"fiber/internal/models"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (s *FiberServer) getRatingsForRecipeHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ratings, err := s.db.GetRatingsForRecipe(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ratings)
}

func (s *FiberServer) addRatingToRecipeHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := s.db.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	recipe, err := s.db.GetRecipeByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ratingToAdd := new(dtos.AddRatingToRecipeDto)
	if err := c.BodyParser(ratingToAdd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rating := &models.Rating{
		Rating:   ratingToAdd.Rating,
		UserID:   user.ID,
		RecipeID: recipe.ID,
	}

	err = s.db.AddRatingToRecipe(rating)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rating)
}

func (s *FiberServer) removeRatingFromRecipeHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ratingID, err := uuid.Parse(c.Params("ratingID"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := s.db.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	recipe, err := s.db.GetRecipeByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rating, err := s.db.GetRatingByID(ratingID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if recipe.UserID != user.ID && rating.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You are not authorized to remove this rating",
		})
	}

	err = s.db.DeleteRating(rating)

	return c.JSON(fiber.Map{
		"message": "Rating removed successfully",
	})
}
