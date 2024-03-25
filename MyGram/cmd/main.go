package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ardipermana59/mygram/internal/config"
	"github.com/ardipermana59/mygram/internal/router"
)

func main() {
	// Load environment variables
	if err := config.LoadEnvVariables(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	// Initialize router
	r := router.SetupRouter()

	// Start server
	port := os.Getenv("PORT")
	fmt.Println(os.Getenv("DB_HOST"))
	if port == "" {
		port = "8000" // Default port
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
