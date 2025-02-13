package service

import (
	"context"
	"errors"

	"github.com/Happy1353/Avito/internal/middleware"
	"github.com/Happy1353/Avito/internal/repository"
)

type TransactionService struct {
	transactionRepo *repository.TransactionRepository
	userRepo        *repository.UserRepository
}

func NewTransactionService(transactionRepo *repository.TransactionRepository, userRepo *repository.UserRepository) *TransactionService {
	return &TransactionService{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}

func (s *TransactionService) Transaction(ctx context.Context, receiver string, amount int) error {
	if amount <= 0 {
		return errors.New("Amount must be more then 0")
	}

	userID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return errors.New("failed to retrieve user ID from context")
	}

	sender, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return errors.New("failed to get user")
	}

	if sender.Username == receiver {
		return errors.New("Unable to send transaction to yourself")
	}

	if sender.Balance < amount {
		return errors.New("insufficient funds")
	}

	err = s.transactionRepo.UpdateBalaces(ctx, receiver, sender.Username, amount)
	if err != nil {
		return errors.New("failed to update balances")
	}

	return nil
}
