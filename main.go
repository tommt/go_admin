package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=db user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Kathmandu"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.

	if err != nil {
		panic("Could not connect with the database")
	}
	sqlDB.SetMaxIdleConns(10)
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Daju!!")
	})

	app.Listen(":8000")
}
