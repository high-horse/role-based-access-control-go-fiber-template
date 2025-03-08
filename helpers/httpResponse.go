package helpers

import (
	"fmt"
	"rbac/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"strings"
)

func ResponseJson(c *fiber.Ctx, statusCode int, response interface{}) error {
	c.Status(statusCode)
	return c.JSON(response)
}

func ResponseString(c *fiber.Ctx, statusCode int, response string) error{
	c.Status(statusCode)
	return c.SendString(response)
}

func ResponseError(c *fiber.Ctx, statusCode int, errorMessage string) error {
	c.Status(statusCode)
	
	errResponse := map[string]string {
		"error" : errorMessage,
	}
	
	return c.JSON(errResponse)
}

func ResponseValidationError(c *fiber.Ctx, validationErrors  []utils.ErrorResponse) error {
	if len(validationErrors) > 0 {
		errMsgs := make([]string, 0)
		for _, err := range validationErrors {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | needs to be implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		return ResponseError(c, 422, strings.Join(errMsgs, " and "))
	}
	return nil
}