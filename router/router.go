package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/go-fiber-auth/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	product := api.Group("/products")
	product.Get("/", handler.GetAllProduct)
	product.Get("/:id", handler.GetProduct)
}
