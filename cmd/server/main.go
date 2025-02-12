package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Happy1353/Avito/config"
	"github.com/Happy1353/Avito/internal/database"
	"github.com/Happy1353/Avito/internal/repository"
	"github.com/Happy1353/Avito/internal/router"
)

func main() {
	dbConfig := config.LoadDatabaseConfig()

	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	r := router.NewRouter(userRepo, jwtSecret)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
