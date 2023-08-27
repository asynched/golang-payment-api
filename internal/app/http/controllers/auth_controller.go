package controllers

import (
	"log"

	"github.com/asynched/golang-payment-api/internal/app/schemas"
	"github.com/asynched/golang-payment-api/internal/database/repositories"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	userRepository *repositories.UserRepository
}

func (auth *AuthController) Register(ctx *fiber.Ctx) error {
	data := schemas.CreateUserSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	if valid, message := data.IsValid(); !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": message,
		})
	}

	input := repositories.CreateUserInput{
		Name:     data.Name,
		Cpf:      data.Cpf,
		Email:    data.Email,
		Password: data.Password,
	}

	if err := auth.userRepository.CreateUser(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (auth *AuthController) SignIn(ctx *fiber.Ctx) error {
	data := schemas.SignInSchema{}

	if err := ctx.BodyParser(&data); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid data",
		})
	}

	if valid, message := data.IsValid(); !valid {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": message,
		})
	}

	input := repositories.GetUserByEmailInput{
		Email: data.Email,
	}

	user, err := auth.userRepository.GetUserByEmail(input)

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

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": "gibberish",
	})
}

func NewAuthController(userRepository *repositories.UserRepository) *AuthController {
	return &AuthController{userRepository}
}
