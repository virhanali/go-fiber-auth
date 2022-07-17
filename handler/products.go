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
	productReq := models.ProductRequest{}

	if err := c.BodyParser(&productReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "bad request",
			"error":   err,
		})
	}

	products := models.Product{
		Title:       productReq.Title,
		Description: productReq.Description,
		Amount:      productReq.Amount,
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

func UpdateProduct(c *fiber.Ctx) error {
	productReq := models.ProductRequest{}

	if err := c.BodyParser(&productReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "bad request",
			"error":   err,
		})
	}

	id := c.Params("id")

	var products models.Product

	if err := database.DB.Debug().First(&products, "id = ?", id).Error; err != nil {
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": "update product failed",
			"error":   err.Error(),
		})
	}

	products.Title = productReq.Title
	products.Description = productReq.Description
	products.Amount = productReq.Amount

	if err := database.DB.Debug().Save(&products).Error; err != nil {
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": "save product failed",
			"error":   err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"status":   "success",
		"messages": "update product success",
		"data":     products,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	products := models.Product{}
	id := c.Params("id")

	if err := database.DB.Debug().First(&products, id).Error; err != nil {
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": "No product found with ID",
			"error":   err.Error(),
		})
	}

	if err := database.DB.Debug().Delete(&products).Error; err != nil {
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": "delete product failed",
			"error":   err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "delete product success",
	})
}
