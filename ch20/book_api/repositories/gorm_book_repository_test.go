package repositories

import (
	"testing"

	"book_api/database"
	"book_api/models"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGormBookRepository(t *testing.T) {
	// Initialize SQLite in-memory database
	db, err := database.InitDB(sqlite.Open(":memory:"))
	assert.NoError(t, err, "Database connection failed")

	// Create table
	err = db.AutoMigrate(&models.Book{})
	assert.NoError(t, err, "Table creation failed")

	// Initialize GormBookRepository
	repo := &GormBookRepository{DB: db}

	// Test CreateBook
	book := models.Book{Title: "Test Book", Author: "Test Author", Year: 2023}
	err = repo.CreateBook(book)
	assert.NoError(t, err, "No error should occur in CreateBook.")

	// Test FetchBooks
	books, err := repo.FetchBooks()
	assert.NoError(t, err, "No error should occur in FetchBooks.")
	assert.Equal(t, 1, len(books), "The number of books should be 1.")

	// Test FetchBookByID
	fetchedBook, err := repo.FetchBookByID(1)
	assert.NoError(t, err, "No error should occur in FetchBookByID.")
	assert.Equal(t, "Test Book", fetchedBook.Title, "The book title should match.")

	// Test UpdateBook
	updatedBook := models.Book{
		Model:  gorm.Model{ID: 1},
		Title:  "Updated Book",
		Author: "Updated Author",
		Year:   2024,
	}
	err = repo.UpdateBook(updatedBook)
	assert.NoError(t, err, "No error should occur in UpdateBook.")

	// Test DeleteBook
	err = repo.DeleteBook(1)
	assert.NoError(t, err, "No error should occur in DeleteBook.")
	books, err = repo.FetchBooks()
	assert.NoError(t, err, "No error should occur in FetchBooks.")
	assert.Equal(t, 0, len(books), "The number of books should be 0.")
}
