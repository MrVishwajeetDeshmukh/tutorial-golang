package controllers

import (
	"errors"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"book_api/models"
	"book_api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestShowIndexPage_Success(t *testing.T) {
	// 1. Define mock data for testing
	mockBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "Test Book 1", Author: "Test Author 1", Year: 2023},
		{Model: gorm.Model{ID: 2}, Title: "Test Book 2", Author: "Test Author 2", Year: 2022},
	}

	// 2. Initialize Mock Repository and Controller
	mockRepo := &repositories.MockBookRepository{MockBooks: mockBooks}
	webController := &WebController{Repository: mockRepo}

	// 3. Create test context
	w := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(w)

	// 4. Set up HTML template
	gin.SetMode(gin.TestMode)
	tmpl := template.Must(template.New("tmpl").Parse(`
{{define "index.html"}}Index Page: {{len .Books}}{{end}}
{{define "error.html"}}Error Page: {{.error}}{{end}}
`))
	engine.SetHTMLTemplate(tmpl)

	// 5. Execute handler
	webController.ShowIndexPage(c)

	// 6. Verify results
	assert.Equal(t, http.StatusOK, w.Code)
	expected := "Index Page: 2"
	assert.Contains(t, w.Body.String(), expected)
}

func TestShowIndexPage_Error(t *testing.T) {
	// 1. Define mock data for testing
	mockRepo := &repositories.MockBookRepository{
		MockErr: errors.New("fetch error"),
	}

	// 2. Initialize Controller
	webController := &WebController{Repository: mockRepo}

	// 3. Create test context
	w := httptest.NewRecorder()
	c, engine := gin.CreateTestContext(w)

	// 4. Set up HTML template
	gin.SetMode(gin.TestMode)
	tmpl := template.Must(template.New("tmpl").Parse(`
{{define "index.html"}}Index Page: {{len .Books}}{{end}}
{{define "error.html"}}Error Page: {{.error}}{{end}}
`))
	engine.SetHTMLTemplate(tmpl)

	// 5. Execute handler
	webController.ShowIndexPage(c)

	// 6. Verify results
	assert.Equal(t, http.StatusInternalServerError, w.Code)
	expected := "Error Page: Unable to load data."
	assert.Contains(t, w.Body.String(), expected)
}
