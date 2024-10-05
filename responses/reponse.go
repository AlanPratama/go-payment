package responses

import "github.com/gofiber/fiber/v2"

func Response(statusCode int, message string, data interface{}) fiber.Map {
	return fiber.Map{
		"statusCode": statusCode,
		"message":    message,
		"data":       data,
	}
}
