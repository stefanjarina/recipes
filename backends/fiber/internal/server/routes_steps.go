package server

import (
	"fiber/internal/dtos"
	"fiber/internal/models"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (s *FiberServer) getStepsForRecipeHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	steps, err := s.db.GetStepsForRecipe(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(steps)
}

func (s *FiberServer) addStepToRecipeHandler(c *fiber.Ctx) error {
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

	stepToAdd := new(dtos.CreateStepDto)
	if err := c.BodyParser(stepToAdd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	step := &models.Step{
		Title:        stepToAdd.Title,
		Instructions: stepToAdd.Instructions,
		StepNumber:   stepToAdd.StepNumber,
	}

	s.db.AddStepToRecipe(recipe, step)

	return c.JSON(step)
}

func (s *FiberServer) removeStepFromRecipeHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	stepID, err := uuid.Parse(c.Params("stepID"))
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
			"error": "You are not authorized to remove steps from this recipe",
		})
	}

	step, err := s.db.GetStepByID(stepID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = s.db.DeleteStep(step)

	return c.JSON(fiber.Map{
		"message": "Step removed successfully",
	})
}

func (s *FiberServer) getIngredientsForStepHandler(c *fiber.Ctx) error {
	stepID, err := uuid.Parse(c.Params("stepID"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredients, err := s.db.GetIngredientsForStep(stepID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredientsToReturn := make([]dtos.StepIngredientDto, len(ingredients))
	for i, ingredient := range ingredients {
		ingredientsToReturn[i] = dtos.StepIngredientDto{
			ID:             ingredient.ID,
			IngredientName: ingredient.Ingredient.Name,
			IngredientID:   ingredient.IngredientID,
			Amount:         ingredient.Amount,
			AmountMetric:   ingredient.AmountMetric,
			Measurement:    ingredient.Measurement,
		}
	}

	return c.JSON(ingredientsToReturn)
}

func (s *FiberServer) AddIngredientToStepHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	stepID, err := uuid.Parse(c.Params("stepID"))
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

	step, err := s.db.GetStepByID(stepID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if recipe.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You are not authorized to add ingredients to this step",
		})
	}

	ingredientToAdd := new(dtos.AddIngredientToStepDto)
	if err := c.BodyParser(ingredientToAdd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredient := &models.StepIngredient{
		IngredientID:  ingredientToAdd.IngredientID,
		Amount:        ingredientToAdd.Amount,
		AmountMetric:  ingredientToAdd.AmountMetric,
		MeasurementID: ingredientToAdd.MeasurementID,
	}

	s.db.AddIngredientToStep(step, ingredient)

	return c.JSON(ingredient)
}

func (s *FiberServer) removeIngredientFromStepHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	stepID, err := uuid.Parse(c.Params("stepID"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredientID, err := uuid.Parse(c.Params("ingredientID"))
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
			"error": "You are not authorized to remove ingredients from this recipe step",
		})
	}

	_, err = s.db.GetStepByID(stepID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ingredient, err := s.db.GetStepIngredientByID(ingredientID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err = s.db.DeleteStepIngredient(ingredient)

	return c.JSON(fiber.Map{
		"message": "Step removed successfully",
	})
}
