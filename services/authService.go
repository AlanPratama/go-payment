package services

import (
	"merchant-bank-api/database"
	"merchant-bank-api/models"
	"merchant-bank-api/responses"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "c78a0d9c021c2a8ef918a188b7ddef78ec0e34329750e2b25d9808e0addc8b7a"

func Register(c *fiber.Ctx) (*models.User, error) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return nil, err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if data["role"] == "customer" {
		user.Role = models.CustomerRole
	} else if data["role"] == "merchant" {
		user.Role = models.MerchantRole
	}

	database.DB.Create(&user)

	if user.Id == 0 {
		return nil, responses.ErrResponse(fiber.StatusBadRequest, "Email Already Taken!")
	}

	return &user, nil
}

func Login(c *fiber.Ctx) (*string, error) {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return nil, err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		return nil, responses.ErrResponse(fiber.StatusNotFound, "User Not Found")
	}

	loginHistory := models.History{
		UserID: user.Id,
		Type:   models.LoginType,
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		loginHistory.Status = models.StatusFailed
		database.DB.Create(&loginHistory)
		return nil, responses.ErrResponse(fiber.StatusBadRequest, "Incorrect Password")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.Id)),
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		return nil, responses.ErrResponse(fiber.StatusInternalServerError, "Could not login")
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	loginHistory.Status = models.StatusSuccess
	database.DB.Create(&loginHistory)

	return &token, nil
}

func Logout(c *fiber.Ctx) error {
	user, _ := GetUser(c)

	database.DB.Create(&models.History{
		UserID: user.Id,
		Type:   models.LogoutType,
		Status: models.StatusSuccess,
	})

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return nil
}

func GetUser(c *fiber.Ctx) (*models.User, error) {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return &user, nil
}
