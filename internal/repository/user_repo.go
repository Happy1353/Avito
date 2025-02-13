package repository

import (
	"context"
	"database/sql"

	"github.com/Happy1353/Avito/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, username, password string) (*domain.User, error) {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username, password, balance`
	var user domain.User
	err := r.db.QueryRowContext(ctx, query, username, password).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, userId int) (domain.User, error) {
	query := `SELECT id, username, password, balance FROM users WHERE id = $1`
	var user domain.User
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
