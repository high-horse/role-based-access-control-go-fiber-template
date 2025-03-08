package middleware

import (
	"net/http"
	"rbac/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// If the token is missing, return unauthorized
func JWTAuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Missing token"})
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		// Store claims in the context for later use
		c.Locals("claims", claims)

		// Proceed to the next handler
		return c.Next()
	}
}
