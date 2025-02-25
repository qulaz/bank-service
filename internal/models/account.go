package models

import (
	"time"
)

// Account представляет банковский счет пользователя
type Account struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Number    string    `json:"number"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AccountCreateRequest представляет запрос на создание счета
type AccountCreateRequest struct {
	Currency string `json:"currency" validate:"required,oneof=USD EUR RUB"`
}

// AccountResponse представляет ответ с данными счета
type AccountResponse struct {
	ID       int64  `json:"id"`
	Number   string `json:"number"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

// AccountsResponse представляет ответ со списком счетов
type AccountsResponse struct {
	Accounts []AccountResponse `json:"accounts"`
}
