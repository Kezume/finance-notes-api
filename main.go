package main

import (
	"log"

	"github.com/Kezume/finance-notes/database"
	"github.com/Kezume/finance-notes/database/migration"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DBConnect()
	migration.RunMigration()

	app := fiber.New()

	log.Panic(app.Listen(":8000"))
}