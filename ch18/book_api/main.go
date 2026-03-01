package main

import (
	"log"
	"os"

	"book_api/database"
	"book_api/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	// Log file configuration
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)
	log.Println("Server started")
}

func main() {
	database.InitDB()

	r := gin.New()

	// Add logging middleware
	r.Use(gin.LoggerWithWriter(log.Writer()))
	r.Use(gin.Recovery())

	// Setup routes
	routes.SetupRoutes(r)

	r.Run(":8080")
}
