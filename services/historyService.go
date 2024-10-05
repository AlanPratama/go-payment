package services

import (
	"merchant-bank-api/database"
	"merchant-bank-api/models"

	"github.com/gofiber/fiber/v2"
)

func GetLoginHistory(c *fiber.Ctx) ([]*models.History, error) {
	user, _ := GetUser(c)

	var loginHistory []*models.History
	database.DB.Where("type = ? AND user_id = ?", models.LoginType, user.Id).Order("created_at DESC").Find(&loginHistory)

	return loginHistory, nil
}

func GetLogoutHistory(c *fiber.Ctx) ([]*models.History, error) {
	user, _ := GetUser(c)

	var logoutHistory []*models.History
	database.DB.Where("type = ? AND user_id = ?", models.LogoutType, user.Id).Order("created_at DESC").Find(&logoutHistory)

	return logoutHistory, nil
}
