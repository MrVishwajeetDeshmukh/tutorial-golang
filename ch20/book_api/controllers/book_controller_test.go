package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"book_api/models"
	"book_api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetBooks(t *testing.T) {
	// 1. Define mock data
	mockBooks := []models.Book{
		{
			Model:  gorm.Model{ID: 1},
			Title:  "Test Book 1",
			Author: "Test Author 1",
			Year:   2023,
		},
		{
			Model:  gorm.Model{ID: 2},
			Title:  "Test Book 2",
			Author: "Test Author 2",
			Year:   2022,
		},
	}

	// 2. Initialize Mock Repository and Controller
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
	controller := &BookController{Repository: mockRepo}

	// 3. Create Mock Gin Context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 4. Execute handler
	controller.GetBooks(c)

	// 5. Verify results
	assert.Equal(t, http.StatusOK, w.Code)

	// 6. Compare response data
	var response []models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, len(mockBooks), len(response)) // Length comparison

	for i, mockBook := range mockBooks {
		assert.Equal(t, mockBook.ID, response[i].ID, "ID does not match.")
		assert.Equal(t, mockBook.Title, response[i].Title, "Title does not match.")
		assert.Equal(t, mockBook.Author, response[i].Author, "Author does not match.")
		assert.Equal(t, mockBook.Year, response[i].Year, "Year does not match.")
	}
}

func TestDeleteBook(t *testing.T) {
	// 1. Create mock data and Repository
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "Test Book", Author: "Test Author", Year: 2023},
	}
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
	bookController := &BookController{Repository: mockRepo}

	// 2. Create Mock Gin Context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 3. Set required data in Mock Context
	c.Set("book", mockBooks[0]) // Replaces the BookLoader middleware behavior

	// 4. Execute handler
	bookController.DeleteBook(c)

	// 5. Verify results
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "deleted")
	assert.Equal(t, 0, len(mockRepo.MockBooks)) // Verify MockBooks list is empty
}

func TestGetBookByID_Success(t *testing.T) {
	// 1. Define mock data
	mockBook := models.Book{Model: gorm.Model{ID: 1}, Title: "Test Book", Author: "Test Author", Year: 2023}

	// 2. Initialize Mock Repository and Controller
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{mockBook}}
	controller := &BookController{Repository: mockRepo}

	// 3. Create Gin Context and set 'book' data (replaces BookLoader role)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("book", mockBook)

	// 4. Execute handler
	controller.GetBookByID(c)

	// 5. Verify results
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Book
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, mockBook.ID, response.ID)
	assert.Equal(t, mockBook.Title, response.Title)
}

func TestGetBookByID_NoBook(t *testing.T) {
	// 1. Initialize Controller (Mock Repository: empty data)
	controller := &BookController{Repository: &repositories.MockBookRepository{}}
	// 2. Create Gin Context (book data not set)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 3. Execute handler
	controller.GetBookByID(c)

	// 4. Verify results
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Failed to load book data")
}

func TestGetBookByID_InvalidType(t *testing.T) {
	// 1. Initialize Controller (Mock Repository: empty data)
	controller := &BookController{Repository: &repositories.MockBookRepository{}}
	// 2. Create Gin Context and set invalid type for 'book' data
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Set("book", "invalid type")

	// 3. Execute handler
	controller.GetBookByID(c)

	// 4. Verify results
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid book data format")
}

func TestCreateBook(t *testing.T) {
	// 1. Initialize Mock Repository and Controller
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{}}
	controller := &BookController{Repository: mockRepo}

	// 2. Define new book data and serialize to JSON
	newBook := models.Book{Title: "New Book", Author: "New Author", Year: 2023}
	jsonData, err := json.Marshal(newBook)
	assert.NoError(t, err)

	// 3. Create Gin Context and set POST request
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/books", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	// 4. Execute handler
	controller.CreateBook(c)

	// 5. Verify results
	assert.Equal(t, http.StatusCreated, w.Code)
	var response models.Book
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, newBook.Title, response.Title)
	assert.Equal(t, newBook.Author, response.Author)
	assert.Equal(t, newBook.Year, response.Year)
	assert.Equal(t, 1, len(mockRepo.MockBooks))
}

func TestUpdateBook(t *testing.T) {
	// 1. Define initial data and initialize Mock Repository/Controller
	initialBook := models.Book{Model: gorm.Model{ID: 1}, Title: "Original Book", Author: "Original Author", Year: 2020}
	mockRepo := &repositories.MockBookRepository{MockBooks: []models.Book{initialBook}}
	controller := &BookController{Repository: mockRepo}

	// 2. Define update data and serialize to JSON
	updatedPayload := models.Book{Title: "Updated Book", Author: "Updated Author", Year: 2023}
	jsonData, err := json.Marshal(updatedPayload)
	assert.NoError(t, err)

	// 3. Create Gin Context, set PUT request and 'book' data
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("PUT", "/books/1", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("book", initialBook)

	// 4. Execute handler
	controller.UpdateBook(c)

	// 5. Verify results
	assert.Equal(t, http.StatusOK, w.Code)
	var response models.Book
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, initialBook.ID, response.ID)
	assert.Equal(t, updatedPayload.Title, response.Title)
	assert.Equal(t, updatedPayload.Author, response.Author)
	assert.Equal(t, updatedPayload.Year, response.Year)
	assert.Equal(t, response, mockRepo.MockBooks[0])
}
