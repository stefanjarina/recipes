package server

import (
	"fiber/internal/dtos"
	"fiber/internal/models"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (s *FiberServer) getIngredientsForRecipeHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredients, err := s.db.GetIngredientsForRecipe(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(ingredients)
}

func (s *FiberServer) addIngredientToRecipeHandler(c *fiber.Ctx) error {
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

	if recipe.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You are not authorized to add ingredients to this recipe",
		})
	}

	ingredientToAdd := new(dtos.AddIngredientToRecipeDto)
	if err := c.BodyParser(ingredientToAdd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredient := &models.Ingredient{
		Name:          ingredientToAdd.Name,
		Amount:        ingredientToAdd.Amount,
		AmountMetric:  ingredientToAdd.AmountMetric,
		MeasurementID: ingredientToAdd.MeasurementID,
	}

	s.db.AddIngredientToRecipe(recipe, ingredient)

	return c.JSON(ingredient)
}

func (s *FiberServer) removeIngredientFromRecipeHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredientId, err := uuid.Parse(c.Params("ingredientId"))
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

	if recipe.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You are not authorized to remove ingredients from this recipe",
		})
	}

	ingredient, err := s.db.GetIngredientByID(ingredientId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = s.db.DeleteIngredient(ingredient)

	return c.JSON(fiber.Map{
		"message": "Ingredient removed successfully",
	})
}
