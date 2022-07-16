package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/go-fiber-auth/database"
	"github.com/virhanali/go-fiber-auth/router"
)

func main() {
	database.ConnectDB()

	app := fiber.New()
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
