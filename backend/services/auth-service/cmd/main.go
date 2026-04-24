package main

import (
	"auth-service/database" // Import database connection package
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database connection
	database.Connect()

	// Create Gin router with default middleware (logger + recovery)
	router := gin.Default()

	// Health check route (to verify server is running)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Auth Service Running",
		})
	})

	// Get port from environment variable
	port := os.Getenv("PORT")

	// Fallback to default port if not defined
	if port == "" {
		port = "8080"
	}

	// Start server on given port
	router.Run(":" + port)
}