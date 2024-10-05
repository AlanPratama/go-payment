package routes

import (
	"merchant-bank-api/controllers"
	"merchant-bank-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(authGroup fiber.Router) {
	authGroup.Post("/register", controllers.Register)
	authGroup.Post("/login", controllers.Login)

	authGroup.Post("/logout", middleware.AuthRequired, controllers.Logout)
	authGroup.Get("/user", middleware.AuthRequired, controllers.User)
}
