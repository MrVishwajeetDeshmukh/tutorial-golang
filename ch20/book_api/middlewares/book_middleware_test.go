package middlewares

import (
	"book_api/models"
	"book_api/repositories"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// Normal case: When requesting with a valid id, 'book' data should be stored in the context.
func TestBookLoader(t *testing.T) {
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "Test Book", Author: "Test Author", Year: 2023},
	}
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/:id", BookLoader(mockRepo), func(c *gin.Context) {
		book, exists := c.Get("book")
		assert.True(t, exists, "'book' data should be stored in the context.")

		bookModel, ok := book.(models.Book)
		assert.True(t, ok, "'book' data in the context should be of type models.Book.")

		c.JSON(http.StatusOK, bookModel)
	})

	req := httptest.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Book")
}

// Invalid ID case: When passing a non-numeric id, 400 Bad Request should be returned.
func TestBookLoader_InvalidID(t *testing.T) {
	mockRepo := &repositories.MockBookRepository{}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/:id", BookLoader(mockRepo), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "should not reach"})
	})

	req := httptest.NewRequest("GET", "/books/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid ID")
}

// Non-existent book: If the book with the given id does not exist in the Repository, 404 Not Found should be returned.
func TestBookLoader_NotFound(t *testing.T) {
	// Use empty MockBookRepository: no book with id=1.
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{}}
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/books/:id", BookLoader(mockRepo), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "should not reach"})
	})

	req := httptest.NewRequest("GET", "/books/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Book not found")
}
