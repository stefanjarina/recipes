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
	s.App.Get("/api/protected", s.jwtMiddleware, s.ProtectedHandler)

	s.App.Get("/api/recipes", s.getRecipesHandler)
	s.App.Get("/api/recipes/:id", s.getRecipeHandler)
	s.App.Post("/api/recipes", s.jwtMiddleware, s.createRecipeHandler)
	s.App.Patch("/api/recipes/:id", s.jwtMiddleware, s.updateRecipeHandler)
	s.App.Delete("/api/recipes/:id", s.jwtMiddleware, s.deleteRecipeHandler)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}
