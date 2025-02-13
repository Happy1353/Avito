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
