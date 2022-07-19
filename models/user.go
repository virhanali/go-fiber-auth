package models

import "gorm.io/gorm"

// User struct
type User struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"unique_index;not null" json:"username"`
	Email      string `gorm:"unique_index;not null" json:"email"`
	Password   string `gorm:"not null" json:"password"`
	Name       string `json:"name"`
}

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}

type UserData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
