package main

import (
	"fmt"
	"log"

	"github.com/Kezume/finance-notes/database"
	"github.com/Kezume/finance-notes/database/migration"
	"github.com/gofiber/fiber/v2"
)

func main() {
	defer func()  {
		if r := recover(); r != nil {
			fmt.Println("Sistem masih dalam perbaikan")
		}
	}()
	
	database.DBConnect()
	migration.RunMigration()

	app := fiber.New()

	log.Panic(app.Listen(":8000"))
}