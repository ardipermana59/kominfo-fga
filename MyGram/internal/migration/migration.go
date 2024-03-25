package main

import (
	"log"

	"github.com/ardipermana59/mygram/internal/infrastructure"
	"github.com/ardipermana59/mygram/internal/model"
)

func main() {
	// Koneksi ke database
	db, err := infrastructure.Database()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrasi model User
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatalf("Failed to migrate User model: %v", err)
	}
}
