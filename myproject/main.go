package main

import (
	"log"
	"myproject/config"
	"myproject/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	config.InitDB()

	// Create a new Gin router
	router := gin.Default()

	// Define authentication routes
	routes.AuthRoutes(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
