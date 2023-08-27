package http

import (
	"github.com/asynched/golang-payment-api/internal/app/http/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	statusController := controllers.NewStatusController()

	app.Get("/status", statusController.Status)

}
