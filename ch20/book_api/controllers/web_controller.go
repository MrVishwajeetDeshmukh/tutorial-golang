package controllers

import (
	"book_api/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebController struct {
	Repository repositories.BookRepository
}

func (wc *WebController) ShowIndexPage(c *gin.Context) {
	books, err := wc.Repository.FetchBooks()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Unable to load data."})
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"Books": books})
}
