package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Happy1353/Avito/internal/domain"
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

func (r *TransactionRepository) UpdateBalanceUser(ctx context.Context, user string, amount int) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET balance = balance - $1 WHERE username = $2", amount, user)
	if err != nil {
		return fmt.Errorf("failed to update user balance: %w", err)
	}

	return nil
}

func (r *TransactionRepository) GetTransactionHistory(ctx context.Context, username string) (domain.CoinHistory, error) {
	query := `
		SELECT sender, receiver, amount, created_at
		FROM transactions
		WHERE sender = $1 OR receiver = $1
		ORDER BY created_at DESC
	`
	rows, err := r.db.QueryContext(ctx, query, username)
	if err != nil {
		return domain.CoinHistory{}, fmt.Errorf("failed to query transaction history: %w", err)
	}
	defer rows.Close()

	var history domain.CoinHistory
	for rows.Next() {
		var sender, receiver string
		var amount int
		var createdAt time.Time

		if err := rows.Scan(&sender, &receiver, &amount, &createdAt); err != nil {
			return domain.CoinHistory{}, fmt.Errorf("failed to scan transaction row: %w", err)
		}

		if sender == username {
			history.Sent = append(history.Sent, domain.CoinTransaction{
				User:   receiver,
				Amount: amount,
			})
		}
		if receiver == username {
			history.Received = append(history.Received, domain.CoinTransaction{
				User:   sender,
				Amount: amount,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return domain.CoinHistory{}, fmt.Errorf("error iterating over transaction rows: %w", err)
	}

	return history, nil
}
