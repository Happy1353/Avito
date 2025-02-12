package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Happy1353/Avito/internal/repository"
	"github.com/Happy1353/Avito/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}


func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials repository.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Выводим полученные данные
	fmt.Println("user:", credentials.Username, "password:", credentials.Password)

	token, err := h.authService.Login(r.Context(), credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))
}