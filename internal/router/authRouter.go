package router

import (
	// "rpba-app/internal/auth"

	"rpba-app/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	r := app.Group("/auth")
	
	r.Post("login", auth.Login)
	r.Post("register", auth.Register)
	
	r.Get("test-user", getUsers)
	r.Get("hello", hello)
	
}

func getUsers(c *fiber.Ctx) error {
	// Get users logic
	return c.JSON(fiber.Map{"message": "Get users"})
}

func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, 1123!")
}