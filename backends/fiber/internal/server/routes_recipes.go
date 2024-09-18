package server

import (
	"fiber/internal/dtos"
	"fiber/internal/models"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (s *FiberServer) getRecipesHandler(c *fiber.Ctx) error {
	recipes, err := s.db.GetRecipes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recipes)
}

func (s *FiberServer) getRecipeHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
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

	return c.JSON(recipe)
}

func (s *FiberServer) createRecipeHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	user, err := s.db.GetUserByEmail(email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	recipeToCreate := new(dtos.CreateRecipeDto)
	if err := c.BodyParser(recipeToCreate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	recipe := &models.Recipe{
		ID:          uuid.New(),
		Title:       recipeToCreate.Title,
		Description: recipeToCreate.Description,
		UserID:      user.ID,
	}

	err = s.db.CreateRecipe(recipe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recipe)
}

func (s *FiberServer) updateRecipeHandler(c *fiber.Ctx) error {
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

	recipeToUpdate := new(dtos.UpdateRecipeDto)
	if err := c.BodyParser(recipeToUpdate); err != nil {
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
			"error": "You are not authorized to update this recipe",
		})
	}

	recipe.Title = recipeToUpdate.Title
	recipe.Description = recipeToUpdate.Description

	err = s.db.UpdateRecipe(recipe)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recipe)
}

func (s *FiberServer) patchRecipeHandler(c *fiber.Ctx) error {
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

	recipeToUpdate := new(dtos.UpdateRecipeDto)
	if err := c.BodyParser(recipeToUpdate); err != nil {
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
			"error": "You are not authorized to update this recipe",
		})
	}

	updatedFields := map[string]interface{}{}
	if recipeToUpdate.Title != "" {
		updatedFields["title"] = recipeToUpdate.Title
	}
	if recipeToUpdate.Description != "" {
		updatedFields["description"] = recipeToUpdate.Description
	}
	if recipeToUpdate.Visibility != "" {
		updatedFields["visibility"] = recipeToUpdate.Visibility
	}

	err = s.db.PatchRecipe(recipe, updatedFields)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(recipe)
}

func (s *FiberServer) deleteRecipeHandler(c *fiber.Ctx) error {
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
			"error": "You are not authorized to delete this recipe",
		})
	}

	err = s.db.DeleteRecipe(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
