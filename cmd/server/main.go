package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Happy1353/Avito/config"
	"github.com/Happy1353/Avito/internal/database"
	handler "github.com/Happy1353/Avito/internal/handlers"
	"github.com/Happy1353/Avito/internal/repository"
	"github.com/Happy1353/Avito/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	dbConfig := config.LoadDatabaseConfig()

	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	sessionRepo := repository.NewSessionRepository(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	authService := service.NewAuthService(userRepo, sessionRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", authHandler.Login)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))}