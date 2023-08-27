package main

import (
	"log"

	"github.com/asynched/golang-payment-api/internal/app/http"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
