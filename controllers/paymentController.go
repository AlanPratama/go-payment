package controllers

import (
	"merchant-bank-api/responses"
	"merchant-bank-api/services"

	"github.com/gofiber/fiber/v2"
)

func CreatePayment(c *fiber.Ctx) error {

	payment, err := services.CreatePayment(c)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(err)
	}

	return c.JSON(responses.Response(200, "Transaction Successfully!", payment))
}

func ChangePaymentStatus(c *fiber.Ctx) error {
	payment, err := services.ChangePaymentStatus(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(responses.Response(200, "Transaction Status Changed Successfully!", payment))
}

func GetMyPayments(c *fiber.Ctx) error {
	payments, _ := services.GetMyPayments(c)

	return c.JSON(responses.Response(200, "Get Payments Successfully!", payments))
}

func GetPaymentById(c *fiber.Ctx) error {
	payment, err := services.GetPaymentById(c)

	if err != nil {
		return c.JSON(err)
	}

	return c.JSON(responses.Response(200, "Get Payment Successfully!", payment))
}
