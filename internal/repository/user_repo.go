package repository

import (
	"context"
	"database/sql"
)

type User struct {
	ID       int
	Username string
	Password string
	Balance  int
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, username, password string) (*User, error) {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id, username, password, balance`
	var user User
	err := r.db.QueryRowContext(ctx, query, username, password).Scan(&user.ID, &user.Username, &user.Password, &user.Balance)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
