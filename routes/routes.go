package routes

import (
	"merchant-bank-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	prefix := "/api/v1"

	authRouteGroup := app.Group(prefix + "/auth")
	AuthRoute(authRouteGroup)

	paymentRouteGroup := app.Group(prefix, middleware.AuthRequired)
	PaymentRoute(paymentRouteGroup)

	historyRouteGroup := app.Group(prefix+"/history", middleware.AuthRequired)
	HistoryRoute(historyRouteGroup)

}
