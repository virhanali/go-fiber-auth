package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/go-fiber-auth/handler"
	"github.com/virhanali/go-fiber-auth/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	product := api.Group("/products")
	product.Get("/", handler.GetAllProduct)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Put("/:id", middleware.Protected(), handler.UpdateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)

	users := api.Group("/users")
	users.Get("/:id", handler.GetUser)
	users.Post("/", handler.CreateUser)
}
