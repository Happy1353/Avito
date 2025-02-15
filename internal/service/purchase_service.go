package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Happy1353/Avito/internal/repository"
)

type PurchaseService struct {
	purchesRepo     *repository.PurchesRepository
	userRepo        *repository.UserRepository
	murchandiseRepo *repository.MurchandiseRepository
	transactionRepo *repository.TransactionRepository
}

func NewPurchaseService(purchesRepo *repository.PurchesRepository, userRepo *repository.UserRepository, murchandiseRepo *repository.MurchandiseRepository, transactionRepo *repository.TransactionRepository) *PurchaseService {
	return &PurchaseService{
		purchesRepo:     purchesRepo,
		userRepo:        userRepo,
		murchandiseRepo: murchandiseRepo,
		transactionRepo: transactionRepo,
	}
}

func (s *PurchaseService) BuyItem(ctx context.Context, userID int, itemName string) error {
	item, err := s.murchandiseRepo.GetItem(ctx, itemName)
	if err != nil {
		return errors.New("invalid item")
	}

	user, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return errors.New("failed to get user")
	}

	fmt.Println(user.Balance, item.Price)

	if user.Balance < item.Price {
		return errors.New("insufficient funds")
	}

	err = s.transactionRepo.UpdateBalanceUser(ctx, user.Username, item.Price)
	if err != nil {
		return errors.New("failed to update balance")
	}

	err = s.purchesRepo.AddItem(ctx, userID, item.Id)
	if err != nil {
		return errors.New("failed to buy item")
	}

	return nil
}
