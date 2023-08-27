package main

import (
	"log"

	"github.com/asynched/golang-payment-api/internal/app/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	http.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
