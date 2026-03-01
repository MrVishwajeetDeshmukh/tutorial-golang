package routes

import (
	"book_api/controllers"
	"book_api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	webController *controllers.WebController,
	bookController *controllers.BookController,
) {
	r.LoadHTMLGlob("templates/*") // Load HTML templates

	// Default web page routes
	r.GET("/", webController.ShowIndexPage)

	// API route group
	api := r.Group("/api")
	{
		api.GET("/books", bookController.GetBooks)
		api.GET("/books/:id", middlewares.BookLoader(bookController.Repository), bookController.GetBookByID)
		api.POST("/books", bookController.CreateBook)
		api.PUT("/books/:id", middlewares.BookLoader(bookController.Repository), bookController.UpdateBook)
		api.DELETE("/books/:id", middlewares.BookLoader(bookController.Repository), bookController.DeleteBook)
	}
}
