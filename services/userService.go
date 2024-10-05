package services

import (
	"merchant-bank-api/database"
	"merchant-bank-api/models"
	"merchant-bank-api/responses"

	"github.com/gofiber/fiber/v2"
)

func GetUserById(id uint) (*models.User, error) {
	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	if user.Id == 0 {
		return nil, responses.ErrResponse(fiber.StatusBadRequest, "Amount or Target User ID are required")
	}

	return &user, nil
}
