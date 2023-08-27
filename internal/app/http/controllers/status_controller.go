package controllers

import "github.com/gofiber/fiber/v2"

type StatusController struct {
}

func (status *StatusController) Status(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"status": "ok",
	})
}

func NewStatusController() *StatusController {
	return &StatusController{}
}
