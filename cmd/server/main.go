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
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Error closing the database: %v", err)
		}
	}()

	userRepo := repository.NewUserRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	purchesRepo := repository.NewPurchesRepository(db)
	murchandiseRepo := repository.NewMurchandiseRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Panic("JWT_SECRET is not set")
	}

	r := router.NewRouter(userRepo, transactionRepo, purchesRepo, murchandiseRepo, jwtSecret)

	log.Println("Starting server on :8080")
	log.Panic(http.ListenAndServe(":8080", r))
}
