package database

import (
	"book_api/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Execute migration
	DB.AutoMigrate(&models.Book{})
	log.Println("Database initialized successfully")
}
