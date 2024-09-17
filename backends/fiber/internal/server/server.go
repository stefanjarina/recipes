package server

import (
	"github.com/gofiber/fiber/v2"

	"fiber/internal/database"
	"fiber/internal/middlewares"
)

type FiberServer struct {
	*fiber.App

	db            database.Service
	jwtMiddleware func(*fiber.Ctx) error
}

func New() *FiberServer {
	jwt := middlewares.NewAuthMiddleware()
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "fiber",
			AppName:      "Recipes API",
		}),

		db:            database.New(),
		jwtMiddleware: jwt,
	}

	return server
}
