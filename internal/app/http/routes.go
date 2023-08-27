package http

import (
	"github.com/asynched/golang-payment-api/internal/app/http/controllers"
	"github.com/asynched/golang-payment-api/internal/app/http/middlewares"
	"github.com/asynched/golang-payment-api/internal/database"
	"github.com/asynched/golang-payment-api/internal/database/repositories"
	"github.com/asynched/golang-payment-api/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	db := database.CreateClient()

	// Services
	jwtService := services.NewJwtService()

	// Repositories
	userRepository := repositories.NewUserRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db)

	// Middlewares
	jwtMiddleware := middlewares.NewJwtMiddleware(jwtService)

	statusController := controllers.NewStatusController()
	app.Get("/status", statusController.Status)

	authController := controllers.NewAuthController(userRepository, jwtService)
	app.Post("/auth/register", authController.Register)
	app.Post("/auth/sign-in", authController.SignIn)

	profileController := controllers.NewProfileController(userRepository)
	app.Get("/profile", jwtMiddleware.Handle, profileController.GetProfile)

	transactionController := controllers.NewTransactionController(userRepository, transactionRepository)
	app.Post("/transactions", jwtMiddleware.Handle, transactionController.CreateTransaction)
	app.Get("/transactions", jwtMiddleware.Handle, transactionController.GetTransactions)
	app.Get("/transactions/:id", jwtMiddleware.Handle, transactionController.GetTransactionById)
}
