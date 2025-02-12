package repository

import (
	"context"
	"database/sql"
	"errors"
)

type SessionRepository struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db: db}
}

func (r *SessionRepository) CreateSession(ctx context.Context, userID int, token string) error {
	query := `INSERT INTO sessions (user_id, token) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, userID, token)
	return err
}

func (r *SessionRepository) GetUserIDByToken(ctx context.Context, token string) (int, error) {
	var userID int
	query := `SELECT user_id FROM sessions WHERE token = $1`
	err := r.db.QueryRowContext(ctx, query, token).Scan(&userID)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, errors.New("session not found")
	}
	return userID, err
}

func (r *SessionRepository) DeleteSession(ctx context.Context, token string) error {
	query := `DELETE FROM sessions WHERE token = $1`
	_, err := r.db.ExecContext(ctx, query, token)
	return err
}