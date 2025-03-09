package router

import (
	// "rbac/internal/auth"

	"rbac/internal/auth"
	"rbac/middleware"

	"github.com/gofiber/fiber/v2"
)

func NonAuthRoutes(app *fiber.App) {
	r := app.Group("/auth")
	
	r.Post("login", auth.Login)
	r.Post("register", auth.Register)
	
	// r.Get("profile", auth.Profile)
	r.Get("hello", hello)
	
}

func AuthRoutes(app *fiber.App) {
	r := app.Group("/auth").Use(middleware.JWTAuthMiddleware())
	
	r.Get("profile", auth.Profile)
}

func getUsers(c *fiber.Ctx) error {
	// Get users logic
	return c.JSON(fiber.Map{"message": "Get users"})
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, 1123!")
}