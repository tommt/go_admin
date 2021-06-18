package main

import (
	"go_admin/src/database"
	"go_admin/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	// database.AutoMigrateDB()
	app := fiber.New()

	routes.Setup(app)
	app.Listen(":8000")
}
