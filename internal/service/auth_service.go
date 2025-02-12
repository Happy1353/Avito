package service

import (
	"context"
	"errors"
	"time"

	"github.com/Happy1353/Avito/internal/repository"

	"github.com/golang-jwt/jwt"
)

type AuthService struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

const LIVE_TOKEN_TIME = time.Hour * 24

func NewAuthService(userRepo *repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.CreateUser(ctx, username, password)
	if err != nil {
		return "", errors.New("Username already exist")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(LIVE_TOKEN_TIME).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return tokenString, nil
}
