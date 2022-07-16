package database

import (
	"fmt"
	"strconv"

	"github.com/virhanali/go-fiber-auth/config"
	"github.com/virhanali/go-fiber-auth/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var err error
var DB *gorm.DB

func ConnectDB() {

	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.Product{}, &models.User{})
}
