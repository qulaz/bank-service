package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/qulaz/bank-service/internal/models"
)

// UserRepository представляет репозиторий для работы с пользователями
type UserRepository struct {
	db *DB
}

// NewUserRepository создает новый репозиторий для работы с пользователями
func NewUserRepository(db *DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create создает нового пользователя
func (r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (username, password, email, full_name, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	result, err := r.db.Exec(query, user.Username, user.Password, user.Email, user.FullName, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	user.ID = id
	return nil
}

// GetByID получает пользователя по ID
func (r *UserRepository) GetByID(id int64) (*models.User, error) {
	query := `
		SELECT id, username, password, email, full_name, created_at, updated_at
		FROM users
		WHERE id = ?
	`

	var user models.User
	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// GetByUsername получает пользователя по имени пользователя
func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	query := `
		SELECT id, username, password, email, full_name, created_at, updated_at
		FROM users
		WHERE username = ?
	`

	var user models.User
	err := r.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FullName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// Update обновляет пользователя
func (r *UserRepository) Update(user *models.User) error {
	query := `
		UPDATE users
		SET username = ?, password = ?, email = ?, full_name = ?, updated_at = ?
		WHERE id = ?
	`

	user.UpdatedAt = time.Now()

	_, err := r.db.Exec(query, user.Username, user.Password, user.Email, user.FullName, user.UpdatedAt, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
