package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Happy1353/Avito/internal/domain"
	"github.com/Happy1353/Avito/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials domain.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.authService.Login(r.Context(), credentials.Username, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(token)); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
