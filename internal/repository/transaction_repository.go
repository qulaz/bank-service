package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/qulaz/bank-service/internal/models"
)

// TransactionRepository представляет репозиторий для работы с транзакциями
type TransactionRepository struct {
	db *DB
}

// NewTransactionRepository создает новый репозиторий для работы с транзакциями
func NewTransactionRepository(db *DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

// Create создает новую транзакцию
func (r *TransactionRepository) Create(tx *models.Transaction) error {
	query := `
		INSERT INTO transactions (from_account_id, to_account_id, amount, type, status, reference_number, description, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	tx.CreatedAt = now
	tx.UpdatedAt = now

	result, err := r.db.Exec(
		query,
		tx.FromAccountID,
		tx.ToAccountID,
		tx.Amount,
		tx.Type,
		tx.Status,
		tx.ReferenceNumber,
		tx.Description,
		tx.CreatedAt,
		tx.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	tx.ID = id
	return nil
}

// GetByID получает транзакцию по ID
func (r *TransactionRepository) GetByID(id int64) (*models.Transaction, error) {
	query := `
		SELECT id, from_account_id, to_account_id, amount, type, status, reference_number, description, created_at, updated_at
		FROM transactions
		WHERE id = ?
	`

	var tx models.Transaction
	var toAccountID sql.NullInt64

	err := r.db.QueryRow(query, id).Scan(
		&tx.ID,
		&tx.FromAccountID,
		&toAccountID,
		&tx.Amount,
		&tx.Type,
		&tx.Status,
		&tx.ReferenceNumber,
		&tx.Description,
		&tx.CreatedAt,
		&tx.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}
		return nil, fmt.Errorf("failed to get transaction: %w", err)
	}

	if toAccountID.Valid {
		tx.ToAccountID = toAccountID.Int64
	}

	return &tx, nil
}

// GetByAccountID получает все транзакции счета
func (r *TransactionRepository) GetByAccountID(accountID int64) ([]*models.Transaction, error) {
	query := `
		SELECT id, from_account_id, to_account_id, amount, type, status, reference_number, description, created_at, updated_at
		FROM transactions
		WHERE from_account_id = ? OR to_account_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, accountID, accountID)
	if err != nil {
		return nil, fmt.Errorf("failed to get transactions: %w", err)
	}
	defer rows.Close()

	var transactions []*models.Transaction
	for rows.Next() {
		var tx models.Transaction
		var toAccountID sql.NullInt64

		err := rows.Scan(
			&tx.ID,
			&tx.FromAccountID,
			&toAccountID,
			&tx.Amount,
			&tx.Type,
			&tx.Status,
			&tx.ReferenceNumber,
			&tx.Description,
			&tx.CreatedAt,
			&tx.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %w", err)
		}

		if toAccountID.Valid {
			tx.ToAccountID = toAccountID.Int64
		}

		transactions = append(transactions, &tx)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transactions: %w", err)
	}

	return transactions, nil
}
