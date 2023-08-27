package controllers

import (
	"github.com/asynched/golang-payment-api/internal/database/repositories"
	"github.com/gofiber/fiber/v2"
)

type ProfileController struct {
	userRepository *repositories.UserRepository
}

func (controller *ProfileController) GetProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int)

	user, err := controller.userRepository.GetUserById(repositories.GetUserByIdInput{
		Id: userId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func NewProfileController(userRepository *repositories.UserRepository) *ProfileController {
	return &ProfileController{userRepository}
}
