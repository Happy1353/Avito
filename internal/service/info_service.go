package service

import (
	"context"

	"github.com/Happy1353/Avito/internal/repository"
)

type InfoService struct {
	userRepo *repository.UserRepository
}

func NewInfoService(userRepo *repository.UserRepository) *InfoService {
	return &InfoService{
		userRepo: userRepo,
	}
}

func (s *InfoService) Info(ctx context.Context, userId string) (string, error) {
	return "test", nil
}
