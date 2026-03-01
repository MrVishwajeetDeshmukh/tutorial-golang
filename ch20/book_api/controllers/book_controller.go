package controllers

import (
	"book_api/models"
	"book_api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	Repository repositories.BookRepository
}

func (bc *BookController) GetBooks(c *gin.Context) {
	books, err := bc.Repository.FetchBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (bc *BookController) GetBookByID(c *gin.Context) {
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load book data"})
		return
	}

	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid book data format"})
		return
	}

	c.JSON(http.StatusOK, bookModel)
}

func (bc *BookController) DeleteBook(c *gin.Context) {
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load book data"})
		return
	}

	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid book data format"})
		return
	}

	if err := bc.Repository.DeleteBook(bookModel.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted."})
}

func (bc *BookController) CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	if err := bc.Repository.CreateBook(newBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, newBook)
}

func (bc *BookController) UpdateBook(c *gin.Context) {
	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}
	book, exists := c.Get("book")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load book data"})
		return
	}
	bookModel, ok := book.(models.Book)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid book data format"})
		return
	}
	updatedBook.ID = bookModel.ID
	if err := bc.Repository.UpdateBook(updatedBook); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
		return
	}
	c.JSON(http.StatusOK, updatedBook)
}
