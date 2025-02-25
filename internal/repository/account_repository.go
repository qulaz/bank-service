package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/qulaz/bank-service/internal/models"
)

// AccountRepository представляет репозиторий для работы со счетами
type AccountRepository struct {
	db *DB
}

// NewAccountRepository создает новый репозиторий для работы со счетами
func NewAccountRepository(db *DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// Create создает новый счет
func (r *AccountRepository) Create(account *models.Account) error {
	query := `
		INSERT INTO accounts (user_id, number, balance, currency, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	account.CreatedAt = now
	account.UpdatedAt = now

	result, err := r.db.Exec(query, account.UserID, account.Number, account.Balance, account.Currency, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	account.ID = id
	return nil
}

// GetByID получает счет по ID
func (r *AccountRepository) GetByID(id int64) (*models.Account, error) {
	query := `
		SELECT id, user_id, number, balance, currency, created_at, updated_at
		FROM accounts
		WHERE id = ?
	`

	var account models.Account
	err := r.db.QueryRow(query, id).Scan(
		&account.ID,
		&account.UserID,
		&account.Number,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("account not found")
		}
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return &account, nil
}

// GetByUserID получает все счета пользователя
func (r *AccountRepository) GetByUserID(userID int64) ([]*models.Account, error) {
	query := `
		SELECT id, user_id, number, balance, currency, created_at, updated_at
		FROM accounts
		WHERE user_id = ?
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}
	defer rows.Close()

	var accounts []*models.Account
	for rows.Next() {
		var account models.Account
		err := rows.Scan(
			&account.ID,
			&account.UserID,
			&account.Number,
			&account.Balance,
			&account.Currency,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan account: %w", err)
		}
		accounts = append(accounts, &account)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating accounts: %w", err)
	}

	return accounts, nil
}

// UpdateBalance обновляет баланс счета
func (r *AccountRepository) UpdateBalance(id int64, amount int64) error {
	query := `
		UPDATE accounts
		SET balance = balance + ?, updated_at = ?
		WHERE id = ?
	`

	now := time.Now()

	_, err := r.db.Exec(query, amount, now, id)
	if err != nil {
		return fmt.Errorf("failed to update balance: %w", err)
	}

	return nil
}
