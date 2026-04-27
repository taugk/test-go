package routes

import (
	"test-oldo/config"
	"test-oldo/internal/handler"
	"test-oldo/internal/repository"
	"test-oldo/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	userRepo := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	paketRepo := repository.NewPaketRepository(config.DB)
	paketService := service.NewPaketService(paketRepo)
	paketHandler := handler.NewPaketHandler(paketService)

	transRepo := repository.NewTransaksiRepository(config.DB)
	transService := service.NewTransaksiService(transRepo, paketRepo, userRepo)
	transHandler := handler.NewTransaksiHandler(transService)

	api := app.Group("/api")

	users := api.Group("/users")
	users.Post("/", userHandler.Create)
	users.Get("/", userHandler.GetAll)
	users.Get("/:id", userHandler.GetByID)
	users.Put("/:id", userHandler.Update)
	users.Delete("/:id", userHandler.Delete)

	pakets := api.Group("/paket")
	pakets.Post("/", paketHandler.Create)
	pakets.Get("/", paketHandler.GetAll)
	pakets.Get("/:id", paketHandler.GetByID)
	pakets.Put("/:id", paketHandler.Update)
	pakets.Delete("/:id", paketHandler.Delete)

	trans := api.Group("/transaksis")
	trans.Post("/", transHandler.Create)
	trans.Get("/", transHandler.GetAll)
	trans.Get("/:id", transHandler.GetByID)
}
