package main

import (
	"book_api/controllers"
	"book_api/database"
	"book_api/repositories"
	"book_api/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
)

func init() {
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
	log.Println("Server started")
}

func main() {
	// Initialize database
	db, err := database.InitDB(sqlite.Open("books.db"))
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	// Create Gin engine
	r := gin.New()

	// Add logging and error recovery middleware
	r.Use(gin.LoggerWithWriter(log.Writer()))
	r.Use(gin.Recovery())

	// Dependency injection
	repo := &repositories.GormBookRepository{DB: db}
	webController := &controllers.WebController{Repository: repo}
	bookController := &controllers.BookController{Repository: repo}

	// Setup routes
	routes.SetupRoutes(r, webController, bookController)

	// Start server
	r.Run(":8080")
}
