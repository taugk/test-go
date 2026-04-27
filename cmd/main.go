package main

import (
	"test-oldo/config"
	"test-oldo/internal/model"
	"test-oldo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.ConnDB()

	config.DB.AutoMigrate(
		&model.User{},
		&model.Paket{},
		&model.Transaksi{},
	)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API Running 🚀")
	})

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
