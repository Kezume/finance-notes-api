package controllers

import "github.com/gofiber/fiber/v2"

type TransactionController interface {
	CreateTransaction(ctx *fiber.Ctx) error
}

