package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	var err error
	dsn := "host=localhost user=todo_user password=password dbname=todo_app2 port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error of opening database")
	}

	fmt.Println("Successful opening DB")

	return DB
}

func MigrateDB(db *gorm.DB) {
	if err := db.AutoMigrate(&Task{}); err != nil {
		log.Fatal("Error migrating")
	}
	fmt.Println("Successful migrating")
}
