package controllers

import (
	"github.com/Kezume/finance-notes/models"
	"github.com/Kezume/finance-notes/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionControllerImpl struct {
	service services.TransactionService
}

func NewTransactionController(service services.TransactionService) TransactionController {
	return &TransactionControllerImpl{service: service}
}

// GetAllTransaction implements TransactionController.
func (t *TransactionControllerImpl) GetAllTransaction(ctx *fiber.Ctx) error {
	transaction, err := t.service.GetAllTransaction()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get transactions",
			"details": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success get transactions",
		"data": transaction,
	})
}

// CreateTransaction implements TransactionController.
func (t *TransactionControllerImpl) CreateTransaction(ctx *fiber.Ctx) error {
	var request models.TransactionRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		errors := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, err.Field()+" is "+err.Tag())
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": errors,
		})
	}

	transaction, err := t.service.CreateTransaction(request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create transaction",
			"details": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "data success created",
		"data":    transaction,
	})
}
