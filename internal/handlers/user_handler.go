package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Happy1353/Avito/internal/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) Info(w http.ResponseWriter, r *http.Request) {
	data, err := h.userService.Info(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(response); err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
