package models

import (
	"time"
)

// TransactionType представляет тип транзакции
type TransactionType string

const (
	TransactionTypeDeposit  TransactionType = "deposit"
	TransactionTypeWithdraw TransactionType = "withdraw"
	TransactionTypeTransfer TransactionType = "transfer"
)

// Transaction представляет транзакцию между счетами
type Transaction struct {
	ID              int64           `json:"id"`
	FromAccountID   int64           `json:"from_account_id"`
	ToAccountID     int64           `json:"to_account_id,omitempty"`
	Amount          int64           `json:"amount"`
	Type            TransactionType `json:"type"`
	Status          string          `json:"status"`
	ReferenceNumber string          `json:"reference_number"`
	Description     string          `json:"description"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

// TransferRequest представляет запрос на перевод средств
type TransferRequest struct {
	FromAccountID int64  `json:"from_account_id" validate:"required"`
	ToAccountID   int64  `json:"to_account_id" validate:"required,nefield=FromAccountID"`
	Amount        int64  `json:"amount" validate:"required,gt=0"`
	Description   string `json:"description"`
}

// TransactionResponse представляет ответ с данными транзакции
type TransactionResponse struct {
	ID              int64           `json:"id"`
	FromAccountID   int64           `json:"from_account_id"`
	ToAccountID     int64           `json:"to_account_id,omitempty"`
	Amount          int64           `json:"amount"`
	Type            TransactionType `json:"type"`
	Status          string          `json:"status"`
	ReferenceNumber string          `json:"reference_number"`
	Description     string          `json:"description"`
	CreatedAt       time.Time       `json:"created_at"`
}

// TransactionsResponse представляет ответ со списком транзакций
type TransactionsResponse struct {
	Transactions []TransactionResponse `json:"transactions"`
}
