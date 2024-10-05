package controllers

import (
	"merchant-bank-api/responses"
	"merchant-bank-api/services"

	"github.com/gofiber/fiber/v2"
)

func GetLoginHistory(c *fiber.Ctx) error {
	history, _ := services.GetLoginHistory(c)

	return c.JSON(responses.Response(200, "Get Login History Successfully!", history))
}

func GetLogoutHistory(c *fiber.Ctx) error {
	history, _ := services.GetLogoutHistory(c)

	return c.JSON(responses.Response(200, "Get Logout History Successfully!", history))
}
