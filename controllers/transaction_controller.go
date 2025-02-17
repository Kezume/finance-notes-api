package controllers

import "github.com/gofiber/fiber/v2"

type TransactionController interface {
	GetAllTransaction(ctx *fiber.Ctx) error
	CreateTransaction(ctx *fiber.Ctx) error
}

