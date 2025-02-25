package models

import (
	"time"
)

// OTP представляет одноразовый пароль для подтверждения операций
type OTP struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Code      int       `json:"code"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"created_at"`
}

// OTPGenerateRequest представляет запрос на генерацию OTP-кода
type OTPGenerateRequest struct {
	UserID int64 `json:"user_id" validate:"required"`
}

// OTPVerifyRequest представляет запрос на проверку OTP-кода
type OTPVerifyRequest struct {
	UserID int64  `json:"user_id" validate:"required"`
	Code   string `json:"code" validate:"required,len=6"`
}

// OTPResponse представляет ответ с данными OTP-кода
type OTPResponse struct {
	Message string `json:"message"`
}
