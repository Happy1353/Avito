package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type Murchandise struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type MurchandiseRepository struct {
	db *sql.DB
}

func NewMurchandiseRepository(db *sql.DB) *MurchandiseRepository {
	return &MurchandiseRepository{db: db}
}

func (r *MurchandiseRepository) GetItem(ctx context.Context, name string) (*Murchandise, error) {
	var merchandise Murchandise
	err := r.db.QueryRowContext(ctx, `SELECT id, name, price FROM merchandise WHERE name = $1`, name).
		Scan(&merchandise.Id, &merchandise.Name, &merchandise.Price)
	if err != nil {
		return nil, fmt.Errorf("failed to get merchandise: %w", err)
	}
	return &merchandise, nil
}
