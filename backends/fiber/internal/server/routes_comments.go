package server

import (
	"fiber/internal/dtos"
	"fiber/internal/models"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (s *FiberServer) getCommentsForRecipeHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	comments, err := s.db.GetCommentsForRecipe(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(comments)
}

func (s *FiberServer) addCommentToRecipeHandler(c *fiber.Ctx) error {
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

	commentToAdd := new(dtos.AddCommentToRecipeDto)
	if err := c.BodyParser(commentToAdd); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	comment := &models.Comment{
		CommentText: commentToAdd.CommentText,
		UserID:      user.ID,
		RecipeID:    recipe.ID,
	}

	err = s.db.AddCommentToRecipe(comment)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(comment)
}

func (s *FiberServer) removeCommentFromRecipeHandler(c *fiber.Ctx) error {
	token := c.Locals("user").(*jtoken.Token)
	claims := token.Claims.(jtoken.MapClaims)
	email := claims["email"].(string)

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	commentId, err := uuid.Parse(c.Params("commentId"))
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

	comment, err := s.db.GetCommentByID(commentId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if recipe.UserID != user.ID && comment.UserID != user.ID {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You are not authorized to remove this comment",
		})
	}

	err = s.db.DeleteComment(comment)

	return c.JSON(fiber.Map{
		"message": "Comment removed successfully",
	})
}
