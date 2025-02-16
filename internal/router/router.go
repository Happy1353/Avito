package router

import (
	"github.com/Happy1353/Avito/internal/handlers"
	"github.com/Happy1353/Avito/internal/middleware"
	"github.com/Happy1353/Avito/internal/repository"
	"github.com/Happy1353/Avito/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRouter(userRepo *repository.UserRepository, transactionRepo *repository.TransactionRepository, purchesRepo *repository.PurchesRepository, murchandiseRepo *repository.MurchandiseRepository, jwtSecret string) *chi.Mux {
	authService := service.NewAuthService(userRepo, jwtSecret)
	transactionService := service.NewTransactionService(transactionRepo, userRepo)
	purchaseService := service.NewPurchaseService(purchesRepo, userRepo, murchandiseRepo, transactionRepo)
	userService := service.NewUserService(userRepo, purchesRepo, transactionRepo)

	authHandler := handlers.NewAuthHandler(authService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseService)
	userHandler := handlers.NewUserHandler(userService)

	authMiddleware := middleware.AuthMiddleware(jwtSecret)

	r := chi.NewRouter()

	// CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "Authorization"},
		ExposedHeaders:   []string{"X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Post("/api/auth", authHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware)

		r.Get("/api/info", userHandler.Info)
		r.Post("/api/sendCoin", transactionHandler.Transaction)
		r.Get("/api/buy/{item}", purchaseHandler.BuyItem)
	})

	return r
}
