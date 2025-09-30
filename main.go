package main

import (
	"hospital/database"
	"hospital/router"
	"log"
)

// main is the entry point of the application.
func main() {
	// Initialize database connection
	if err := database.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Setup and run the Gin router
	r := router.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
