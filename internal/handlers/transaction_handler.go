package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Happy1353/Avito/internal/domain"
	"github.com/Happy1353/Avito/internal/service"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (h *TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	var transactionInfo domain.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&transactionInfo); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.transactionService.Transaction(r.Context(), transactionInfo.ToUser, transactionInfo.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("transaction successful")); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
