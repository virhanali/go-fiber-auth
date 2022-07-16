package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/go-fiber-auth/database"
	"github.com/virhanali/go-fiber-auth/models"
)

func GetAllProduct(c *fiber.Ctx) error {
	products := []models.Product{}

	if err := database.DB.Debug().Find(&products).Error; err != nil {
		return c.JSON(fiber.Map{
			"status": "failed",
			"error":  err.Error(),
		})
	}

	if len(products) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":   "failed",
			"messages": "products not found",
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "all products",
		"data":    products,
	})
}

func GetProduct(c *fiber.Ctx) error {
	products := models.Product{}

	id := c.Params("id")
	if err := database.DB.Debug().First(&products, id).Error; err != nil {
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": "No product found with ID",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Product found",
		"data":    products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	products := models.Product{}

	if err := c.BodyParser(&products); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "bad request",
			"error":   err,
		})
	}

	if err := database.DB.Create(&products).Error; err != nil {
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": "create product failed",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   "success",
		"messages": "create product success",
		"data":     products,
	})
}
