package database

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}
