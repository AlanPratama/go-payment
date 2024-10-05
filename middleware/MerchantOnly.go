package middleware

import (
	"merchant-bank-api/models"
	"merchant-bank-api/responses"
	"merchant-bank-api/services"

	"github.com/gofiber/fiber/v2"
)

func MerchantOnly(c *fiber.Ctx) error {
	user, _ := services.GetUser(c)

	if user.Role != models.MerchantRole {
		return c.Status(fiber.StatusForbidden).JSON(responses.ErrResponse(fiber.StatusForbidden, "You don't have permission to access this resource"))
	}

	return c.Next()
}
