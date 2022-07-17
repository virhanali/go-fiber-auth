package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/virhanali/go-fiber-auth/database"
	"github.com/virhanali/go-fiber-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)

	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	if uid != n {
		return false
	}

	return true
}

func validUser(id string, p string) bool {
	user := models.User{}

	if err := database.DB.Debug().First(&user, id).Error; err != nil {
		return false
	}
	if !CheckPasswordHash(p, user.Password) {
		return false
	}
	return false
}

func GetUser(c *fiber.Ctx) error {
	user := models.User{}
	id := c.Params("id")

	if err := database.DB.Debug().First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"status":  "success",
			"message": "No user found with ID",
			"err":     err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "users found",
		"data":    user,
	})
}
