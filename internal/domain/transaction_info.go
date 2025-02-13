package domain

type TransactionRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}
