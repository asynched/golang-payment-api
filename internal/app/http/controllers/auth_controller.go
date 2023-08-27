package controllers

import (
	"log"

	"github.com/asynched/golang-payment-api/internal/app/schemas"
	"github.com/asynched/golang-payment-api/internal/database/repositories"
	"github.com/asynched/golang-payment-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	userRepository *repositories.UserRepository
	jwtService     *services.JwtService
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	data := schemas.CreateUserSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	if ok, errors := data.IsValid(); !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
			"errors":  errors,
		})
	}

	input := repositories.CreateUserInput{
		Name:     data.Name,
		Cpf:      data.Cpf,
		Email:    data.Email,
		Password: data.Password,
	}

	if err := controller.userRepository.CreateUser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (controller *AuthController) SignIn(ctx *fiber.Ctx) error {
	data := schemas.SignInSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	if valid, errors := data.IsValid(); !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
			"errors":  errors,
		})
	}

	input := repositories.GetUserByEmailInput{
		Email: data.Email,
	}

	user, err := controller.userRepository.GetUserByEmail(input)

	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if data.Password != user.Password {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	log.Printf("Authenticating user: %s\n", user)
	token, err := controller.jwtService.Sign(jwt.MapClaims{
		"id": user.Id,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error signing token",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func NewAuthController(userRepository *repositories.UserRepository, jwtService *services.JwtService) *AuthController {
	return &AuthController{userRepository, jwtService}
}
