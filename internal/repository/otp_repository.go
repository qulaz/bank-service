package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/qulaz/bank-service/internal/models"
)

// OTPRepository представляет репозиторий для работы с OTP-кодами
type OTPRepository struct {
	db *DB
}

// NewOTPRepository создает новый репозиторий для работы с OTP-кодами
func NewOTPRepository(db *DB) *OTPRepository {
	return &OTPRepository{db: db}
}

// Create создает новый OTP-код
func (r *OTPRepository) Create(otp *models.OTP) error {
	query := `
		INSERT INTO otps (user_id, code, expires_at, used, created_at)
		VALUES (?, ?, ?, ?, ?)
	`

	now := time.Now()
	otp.CreatedAt = now

	result, err := r.db.Exec(query, otp.UserID, otp.Code, otp.ExpiresAt, otp.Used, otp.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create OTP: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	otp.ID = id
	return nil
}

// GetLatestByUserID получает последний активный OTP-код пользователя
func (r *OTPRepository) GetLatestByUserID(userID int64) (*models.OTP, error) {
	query := `
		SELECT id, user_id, code, expires_at, used, created_at
		FROM otps
		WHERE user_id = ? AND used = 0 AND expires_at > ?
		ORDER BY created_at DESC
		LIMIT 1
	`

	var otp models.OTP
	err := r.db.QueryRow(query, userID, time.Now()).Scan(
		&otp.ID,
		&otp.UserID,
		&otp.Code,
		&otp.ExpiresAt,
		&otp.Used,
		&otp.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("OTP not found")
		}
		return nil, fmt.Errorf("failed to get OTP: %w", err)
	}

	return &otp, nil
}

// MarkAsUsed помечает OTP-код как использованный
func (r *OTPRepository) MarkAsUsed(id int64) error {
	query := `
		UPDATE otps
		SET used = 1
		WHERE id = ?
	`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to mark OTP as used: %w", err)
	}

	return nil
}

// VerifyOTP проверяет OTP-код
func (r *OTPRepository) VerifyOTP(userID int64, code int) (*models.OTP, error) {
	otp, err := r.GetLatestByUserID(userID)
	if err != nil {
		return nil, err
	}

	if otp.Code != code {
		return nil, fmt.Errorf("invalid OTP code")
	}

	if otp.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("OTP code expired")
	}

	if otp.Used {
		return nil, fmt.Errorf("OTP code already used")
	}

	return otp, nil
}
