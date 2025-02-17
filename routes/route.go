package routes

import (
	"github.com/Kezume/finance-notes/controllers"
	"github.com/Kezume/finance-notes/repositories"
	"github.com/Kezume/finance-notes/services"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App) {
	// repositories
	transactionRepository := repositories.NewTransactionRepository()

	// services
	transactionService := services.NewTransactionService(transactionRepository)

	// controllers
	transactionController := controllers.NewTransactionController(transactionService)

	// routes
	app.Post("/api/transaction", transactionController.CreateTransaction)
}