package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) UpdateBalaces(ctx context.Context, receiver, sender string, amount int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			fmt.Printf("failed to rollback transaction: %v", err)
		}
	}()

	// Update balance sender
	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE username = $2", amount, sender)
	if err != nil {
		return fmt.Errorf("failed to update sender balance: %w", err)
	}

	// Update balance receiver
	_, err = tx.ExecContext(ctx, "UPDATE users SET balance = balance + $1 WHERE username = $2", amount, receiver)
	if err != nil {
		return fmt.Errorf("failed to update receiver balance: %w", err)
	}

	_, err = tx.ExecContext(ctx, `
    INSERT INTO transactions (sender, receiver, amount)
    VALUES ($1, $2, $3)
`, sender, receiver, amount)
	if err != nil {
		return fmt.Errorf("failed to record transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
