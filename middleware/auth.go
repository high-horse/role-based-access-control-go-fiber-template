package middleware

import (
	"errors"
	"net/http"

	// "rbac/db/pool"
	// "rbac/pkg/models"
	"rbac/db/pool"
	"rbac/internal/auth"
	"rbac/pkg/models"
	"rbac/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
		
		user, err := getUser(c)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
		}
		auth.SetAuthUser(user)
		
		// Proceed to the next handler
		return c.Next()
	}
}

func getUser(c *fiber.Ctx) (*models.User, error) {
	claims, ok := c.Locals("claims").(*utils.Claims)
	if !ok {
		return nil, errors.New("Auth token error ")
	}
	
	var user models.User
	if err := pool.DB.Where("username = ?", claims.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &user, nil
}
