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
	product.Post("/", handler.CreateProduct)
	product.Put("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)

	users := api.Group("/users")
	users.Get("/:id", handler.GetUser)
}
