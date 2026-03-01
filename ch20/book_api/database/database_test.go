package database

import (
	"book_api/models"
	"testing"

	"github.com/glebarez/sqlite"
)

func TestInitDBWithMemory(t *testing.T) {
	// Call InitDB using an in-memory database
	db, err := InitDB(sqlite.Open(":memory:"))
	if err != nil {
		t.Fatalf("InitDB failed: %v", err)
	}

	// Test data insertion
	testBook := models.Book{Title: "Test Book", Author: "Tester", Year: 2023}
	result := db.Create(&testBook)
	if result.Error != nil {
		t.Fatalf("Data insertion failed: %v", result.Error)
	}

	// Test data retrieval
	var retrievedBook models.Book
	if err := db.First(&retrievedBook, testBook.ID).Error; err != nil {
		t.Fatalf("Data retrieval failed: %v", err)
	}

	// Verify data
	if retrievedBook.Title != testBook.Title {
		t.Errorf("Stored data and retrieved data do not match. (Stored: %v, Retrieved: %v)", testBook.Title, retrievedBook.Title)
	}
}
