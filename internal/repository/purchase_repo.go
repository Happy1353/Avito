package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Happy1353/Avito/internal/domain"
)

type Purchase struct {
	ID            int
	UserID        int
	MerchandiseID int
	CreatedAt     string
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type PurchesRepository struct {
	db *sql.DB
}

func NewPurchesRepository(db *sql.DB) *PurchesRepository {
	return &PurchesRepository{db: db}
}

func (r *PurchesRepository) AddItem(ctx context.Context, userID int, itemID int) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO purchases (user_id, merchandise_id) VALUES ($1, $2)`, userID, itemID)
	if err != nil {
		return fmt.Errorf("failed to insert purchase: %w", err)
	}

	return nil
}

func (r *PurchesRepository) GetUserInventory(ctx context.Context, userID int) ([]domain.InventoryItem, error) {
	query := `
		SELECT m.name, COUNT(p.id) as quantity
		FROM purchases p
		JOIN merchandise m ON p.merchandise_id = m.id
		WHERE p.user_id = $1
		GROUP BY m.name
	`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query user inventory: %w", err)
	}
	defer rows.Close()

	var inventory []domain.InventoryItem
	for rows.Next() {
		var item domain.InventoryItem
		if err := rows.Scan(&item.Type, &item.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan inventory item: %w", err)
		}
		inventory = append(inventory, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over inventory rows: %w", err)
	}

	return inventory, nil
}
