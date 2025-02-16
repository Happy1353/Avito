package domain

type User struct {
	ID       int
	Username string
	Password string
	Balance  int
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type CoinTransaction struct {
	User   string `json:"user"`
	Amount int    `json:"amount"`
}

type CoinHistory struct {
	Received []CoinTransaction `json:"received"`
	Sent     []CoinTransaction `json:"sent"`
}

type UserInfo struct {
	Coins       int             `json:"coins"`
	Inventory   []InventoryItem `json:"inventory"`
	CoinHistory CoinHistory     `json:"coinHistory"`
}
