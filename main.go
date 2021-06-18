package main

import (
	"go_admin/src/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	// database.AutoMigrateDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Daju sanchai ho")
	})

	app.Listen(":8000")
}
