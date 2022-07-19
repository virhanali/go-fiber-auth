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

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.DB
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})

	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})

	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": newUser})
}