package controllers

import (
	"merchant-bank-api/responses"
	"merchant-bank-api/services"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user, err := services.Register(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(responses.Response(200, "Register Successfully!", user))
}

func Login(c *fiber.Ctx) error {
	token, err := services.Login(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(responses.Response(200, "Login Successfully!", token))
}

func User(c *fiber.Ctx) error {
	user, _ := services.GetUser(c)

	return c.JSON(responses.Response(200, "Get User Successfully!", user))
}

func Logout(c *fiber.Ctx) error {
	services.Logout(c)

	return c.JSON(responses.Response(200, "Logout Successfully!", nil))
}
