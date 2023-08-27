package http

import (
	"github.com/asynched/golang-payment-api/internal/app/http/controllers"
	"github.com/asynched/golang-payment-api/internal/database"
	"github.com/asynched/golang-payment-api/internal/database/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := database.CreateClient()
	userRepository := repositories.NewUserRepository(db)

	statusController := controllers.NewStatusController()
	app.Get("/status", statusController.Status)

	authController := controllers.NewAuthController(userRepository)
	app.Post("/auth/register", authController.Register)
	app.Post("/auth/sign-in", authController.SignIn)
}
