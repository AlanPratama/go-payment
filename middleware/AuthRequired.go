package middleware

import (
	"merchant-bank-api/responses"
	"merchant-bank-api/services"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired(c *fiber.Ctx) error {
	_, err := services.GetUser(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrResponse(fiber.StatusUnauthorized, "Unauthorized"))
	}

	return c.Next()
}
