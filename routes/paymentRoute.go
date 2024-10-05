package routes

import (
	"merchant-bank-api/controllers"
	"merchant-bank-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoute(paymentRouteGroup fiber.Router) {
	paymentRouteGroup.Get("/payments", controllers.GetMyPayments)
	paymentRouteGroup.Get("/payments/:paymentID", controllers.GetPaymentById)

	// CUSTOMER ONLY ROUTE
	paymentRouteGroup.Post("/payments", middleware.CustomerOnly, controllers.CreatePayment)

	// MERCHANT ONLY ROUTE
	paymentRouteGroup.Patch("/payments/:paymentID", middleware.MerchantOnly, controllers.ChangePaymentStatus)
}
