package handlers

import (
	"net/http"

	"github.com/Happy1353/Avito/internal/middleware"
	"github.com/Happy1353/Avito/internal/service"
	"github.com/go-chi/chi/v5"
)

type PurchaseHandler struct {
	purchaseService *service.PurchaseService
}

func NewPurchaseHandler(purchaseService *service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{purchaseService: purchaseService}
}

func (h *PurchaseHandler) BuyItem(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(int)
	if !ok {
		http.Error(w, "failed unauthorized", http.StatusUnauthorized)
		return
	}

	itemName := chi.URLParam(r, "item")
	if itemName == "" {
		http.Error(w, "missing item parameter item", http.StatusBadRequest)
		return
	}

	err := h.purchaseService.BuyItem(r.Context(), userID, itemName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("purchase successful")); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
