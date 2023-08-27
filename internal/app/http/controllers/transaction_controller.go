package controllers

import (
	"strconv"

	"github.com/asynched/golang-payment-api/internal/app/schemas"
	"github.com/asynched/golang-payment-api/internal/database/repositories"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	userRepository        *repositories.UserRepository
	transactionRepository *repositories.TransactionRepository
}

func (controller *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int)

	data := schemas.CreateTransactionSchema{}

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

	user, err := controller.userRepository.GetUserById(repositories.GetUserByIdInput{
		Id: userId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if user.Balance < data.Value {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Insufficient funds",
		})
	}

	err = controller.transactionRepository.CreateTransaction(repositories.CreateTransactionInput{
		PayerId: userId,
		PayeeId: data.PayeeId,
		Value:   data.Value,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Transaction created successfully",
	})
}

func (controller *TransactionController) GetTransactions(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int)

	transactions, err := controller.transactionRepository.GetTransactions(repositories.GetTransactionsInput{
		UserId: userId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(transactions)
}

func (controller *TransactionController) GetTransactionById(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(int)
	transactionId, err := strconv.Atoi(ctx.Params("id", "1"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid transaction id",
		})
	}

	transaction, err := controller.transactionRepository.GetTransactionById(repositories.GetTransactionByIdInput{
		Id:     transactionId,
		UserId: userId,
	})

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(transaction)
}

func NewTransactionController(userRepository *repositories.UserRepository, transactionRepository *repositories.TransactionRepository) *TransactionController {
	return &TransactionController{userRepository, transactionRepository}
}
