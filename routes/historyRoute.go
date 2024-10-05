package routes

import (
	"merchant-bank-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func HistoryRoute(historyRouteGroup fiber.Router) {
	historyRouteGroup.Get("/login", controllers.GetLoginHistory)
	historyRouteGroup.Get("/logout", controllers.GetLogoutHistory)
}
