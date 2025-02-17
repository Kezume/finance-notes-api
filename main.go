package main

import (
	"fmt"
	"log"

	"github.com/Kezume/finance-notes/database"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/Kezume/finance-notes/database/migration"
	"github.com/Kezume/finance-notes/routes"
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

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	routes.RouteInit(app)

	log.Panic(app.Listen(":8000"))
}