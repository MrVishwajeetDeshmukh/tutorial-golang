package repositories

import (
	"errors"
	"testing"

	"book_api/models"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestMockBookRepository(t *testing.T) {
	// 1. Initialize mock data
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "Mock Book 1", Author: "Mock Author 1", Year: 2023},
		{Model: gorm.Model{ID: 2}, Title: "Mock Book 2", Author: "Mock Author 2", Year: 2022},
	}

	// 2. Create MockBookRepository
	repo := &MockBookRepository{MockBooks: mockBooks}

	// 3. Test FetchBooks
	books, err := repo.FetchBooks()
	assert.NoError(t, err, "No error should occur in FetchBooks.")
	assert.Equal(t, len(mockBooks), len(books), "The number of books should match.")

	// 4. Test FetchBookByID
	book, err := repo.FetchBookByID(1)
	assert.NoError(t, err, "No error should occur in FetchBookByID.")
	assert.Equal(t, "Mock Book 1", book.Title, "The book title should match.")

	// 5. Test CreateBook
	newBook := models.Book{Model: gorm.Model{ID: 3}, Title: "New Book", Author: "New Author", Year: 2024}
	err = repo.CreateBook(newBook)
	assert.NoError(t, err, "No error should occur in CreateBook.")
	assert.Equal(t, 3, len(repo.MockBooks), "MockBooks length should be 3.")

	// 6. Test UpdateBook
	updatedBook := models.Book{Model: gorm.Model{ID: 1}, Title: "Modified Book", Author: "Modified Author", Year: 2025}
	err = repo.UpdateBook(updatedBook)
	assert.NoError(t, err, "No error should occur in UpdateBook.")
	book, _ = repo.FetchBookByID(1)
	assert.Equal(t, "Modified Book", book.Title, "The book title was not modified.")

	// 7. Test DeleteBook
	err = repo.DeleteBook(1)
	assert.NoError(t, err, "No error should occur in DeleteBook.")
	assert.Equal(t, 2, len(repo.MockBooks), "MockBooks length should be 2.")

	// 8. Test MockErr behavior
	expectedErr := errors.New("mock error")
	repo.MockErr = expectedErr

	// 9. Verify FetchBooks error
	_, err = repo.FetchBooks()
	assert.EqualError(t, err, expectedErr.Error(), "MockErr should be returned from FetchBooks.")

	// 10. Verify FetchBookByID error
	_, err = repo.FetchBookByID(2)
	assert.EqualError(t, err, expectedErr.Error(), "MockErr should be returned from FetchBookByID.")

	// 11. Verify CreateBook error
	err = repo.CreateBook(models.Book{Model: gorm.Model{ID: 4}, Title: "Error Book", Author: "Error Author", Year: 2025})
	assert.EqualError(t, err, expectedErr.Error(), "MockErr should be returned from CreateBook.")

	// 12. Verify UpdateBook error
	err = repo.UpdateBook(models.Book{Model: gorm.Model{ID: 2}, Title: "Error Update", Author: "Error Author", Year: 2026})
	assert.EqualError(t, err, expectedErr.Error(), "MockErr should be returned from UpdateBook.")

	// 13. Verify DeleteBook error
	err = repo.DeleteBook(2)
	assert.EqualError(t, err, expectedErr.Error(), "MockErr should be returned from DeleteBook.")
}
