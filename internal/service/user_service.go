package service

import (
	"context"
	"errors"

	"github.com/Happy1353/Avito/internal/domain"
	"github.com/Happy1353/Avito/internal/middleware"
	"github.com/Happy1353/Avito/internal/repository"
)

type UserService struct {
	userRepo        *repository.UserRepository
	purchesRepo     *repository.PurchesRepository
	transactionRepo *repository.TransactionRepository
}

func NewUserService(userRepo *repository.UserRepository, purchesRepo *repository.PurchesRepository, transactionRepo *repository.TransactionRepository) *UserService {
	return &UserService{
		userRepo:        userRepo,
		purchesRepo:     purchesRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *UserService) Info(ctx context.Context) (*domain.UserInfo, error) {
	userID, ok := ctx.Value(middleware.UserIDKey).(int)
	if !ok {
		return nil, errors.New("failed to retrieve user ID from context")
	}

	userInfo := &domain.UserInfo{}

	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to get user")
	}

	userInfo.Coins = user.Balance

	inventory, err := s.purchesRepo.GetUserInventory(ctx, userID)
	if err != nil {
		return nil, errors.New("failed to get user inventory")
	}

	userInfo.Inventory = inventory

	transactionHistory, err := s.transactionRepo.GetTransactionHistory(ctx, user.Username)
	if err != nil {
		return nil, errors.New("failed to get user transaction history")
	}

	userInfo.CoinHistory = transactionHistory

	return userInfo, nil
}
