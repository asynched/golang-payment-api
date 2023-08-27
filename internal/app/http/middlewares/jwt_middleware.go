package middlewares

import (
	"strings"

	"github.com/asynched/golang-payment-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JwtMiddleware struct {
	jwtService *services.JwtService
}

func (middleware *JwtMiddleware) Handle(ctx *fiber.Ctx) error {
	auth := strings.Replace(ctx.Get("Authorization"), "Bearer ", "", 1)
	claims := jwt.MapClaims{}

	_, err := middleware.jwtService.ValidateToken(auth, claims)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	id := int(claims["id"].(float64))

	ctx.Locals("userId", id)

	return ctx.Next()
}

func NewJwtMiddleware(jwtService *services.JwtService) *JwtMiddleware {
	return &JwtMiddleware{jwtService}
}
