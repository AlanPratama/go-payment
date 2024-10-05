package responses

import "github.com/gofiber/fiber/v2"

func ErrResponse(statusCode int, message string) error {
	return fiber.NewError(statusCode, message)
}
