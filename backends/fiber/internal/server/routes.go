package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/api/health", s.healthHandler)

	s.App.Get("/api/users", s.getUsersHandler)
	s.App.Get("/api/users", s.getUsersHandler)
	s.App.Get("/api/users/:id", s.getUser)
	s.App.Post("/api/users", s.createUser)
	s.App.Put("/api/users/:id", s.jwtMiddleware, s.updateUser)
	s.App.Delete("/api/users/:id", s.jwtMiddleware, s.deleteUser)

	s.App.Post("/api/login", s.LoginHandler)

	s.App.Get("/api/recipes", s.getRecipesHandler)
	s.App.Get("/api/recipes/:id", s.getRecipeHandler)
	s.App.Post("/api/recipes", s.jwtMiddleware, s.createRecipeHandler)
	s.App.Put("/api/recipes/:id", s.jwtMiddleware, s.updateRecipeHandler)
	s.App.Patch("/api/recipes/:id", s.jwtMiddleware, s.patchRecipeHandler)
	s.App.Delete("/api/recipes/:id", s.jwtMiddleware, s.deleteRecipeHandler)

	s.App.Get("/api/recipes/:id/ingredients", s.getIngredientsForRecipeHandler)
	s.App.Post("/api/recipes/:id/ingredients", s.jwtMiddleware, s.addIngredientToRecipeHandler)
	s.App.Delete("/api/recipes/:id/ingredients/:ingredientID", s.jwtMiddleware, s.removeIngredientFromRecipeHandler)

	s.App.Get("/api/recipes/:id/comments", s.getCommentsForRecipeHandler)
	s.App.Post("/api/recipes/:id/comments", s.jwtMiddleware, s.addCommentToRecipeHandler)
	s.App.Delete("/api/recipes/:id/comments/:commentID", s.jwtMiddleware, s.removeCommentFromRecipeHandler)

	s.App.Get("/api/recipes/:id/steps", s.getStepsForRecipeHandler)
	s.App.Post("/api/recipes/:id/steps", s.jwtMiddleware, s.addStepToRecipeHandler)
	s.App.Delete("/api/recipes/:id/steps/:stepID", s.jwtMiddleware, s.removeStepFromRecipeHandler)

	s.App.Get("/api/recipes/:id/steps/:stepID/ingredients", s.getIngredientsForStepHandler)
	s.App.Post("/api/recipes/:id/steps/:stepID/ingredients", s.jwtMiddleware, s.AddIngredientToStepHandler)
	s.App.Delete("/api/recipes/:id/steps/:stepID/ingredients/:ingredientID", s.jwtMiddleware, s.removeIngredientFromStepHandler)

	s.App.Get("/api/recipes/:id/ratings", s.getRatingsForRecipeHandler)
	s.App.Post("/api/recipes/:id/ratings", s.jwtMiddleware, s.addRatingToRecipeHandler)
	s.App.Delete("/api/recipes/:id/ratings/:ratingID", s.jwtMiddleware, s.removeRatingFromRecipeHandler)

	s.App.Get("/api/measurements", s.getMeasurementsHandler)
	s.App.Post("/api/measurements", s.createMeasurementHandler)
	s.App.Delete("/api/measurements/:id", s.deleteMeasurementHandler)
	s.App.Get("/api/measurements/types", s.getMeasurementTypesHandler)
	s.App.Post("/api/measurements/types", s.createMeasurementTypeHandler)
	s.App.Delete("/api/measurements/types/:id", s.deleteMeasurementTypeHandler)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
