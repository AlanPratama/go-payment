package services

import (
	"merchant-bank-api/database"
	"merchant-bank-api/models"
	"merchant-bank-api/responses"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreatePayment(c *fiber.Ctx) (*models.Payment, error) {

	var data map[string]uint

	if err := c.BodyParser(&data); err != nil {
		return nil, err
	}

	if data["amount"] == 0 || data["target_user_id"] == 0 {
		return nil, responses.ErrResponse(fiber.StatusBadRequest, "Amount or Target User ID are required")
	}

	targetUser, err := GetUserById(data["target_user_id"])

	if err != nil {
		return nil, responses.ErrResponse(fiber.StatusNotFound, "Target User Not Found")
	}

	user, _ := GetUser(c)

	payment := models.Payment{
		Amount:        data["amount"],
		UserID:        user.Id,
		TargetUserID:  targetUser.Id,
		PaymentStatus: models.PaymentStatusPending,
	}

	database.DB.Create(&payment)

	return &payment, nil
}

func ChangePaymentStatus(c *fiber.Ctx) (*models.Payment, error) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return nil, responses.ErrResponse(fiber.StatusInternalServerError, err.Error())
	}

	payment, err := GetPaymentById(c)

	if err != nil {
		return nil, responses.ErrResponse(fiber.StatusNotFound, "Payment Not Found")
	}

	if !strings.EqualFold(string(payment.PaymentStatus), string(models.PaymentStatusPending)) {
		return nil, responses.ErrResponse(fiber.StatusBadRequest, "Cannot Change Status, Payment Status Already Failed or Success")
	}

	database.DB.Model(&payment).Update("payment_status", data["status"])

	return payment, nil
}

func GetMyPayments(c *fiber.Ctx) ([]*models.Payment, error) {
	user, _ := GetUser(c)

	var payments []*models.Payment

	database.DB.Where("user_id = ? OR target_user_id = ?", user.Id, user.Id).Find(&payments)

	return payments, nil
}

func GetPaymentById(c *fiber.Ctx) (*models.Payment, error) {
	var payment models.Payment

	user, _ := GetUser(c)

	database.DB.Where("id = ? AND (user_id = ? OR target_user_id = ?)", c.Params("paymentID"), user.Id, user.Id).First(&payment)

	if payment.Id == 0 {
		return nil, responses.ErrResponse(fiber.StatusNotFound, "Payment Not Found")
	}

	return &payment, nil
}
