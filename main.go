package main

import (
	"log"

	"github.com/Kezume/finance-notes/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBConnect()

	app := fiber.New()

	log.Panic(app.Listen(":8000"))
}