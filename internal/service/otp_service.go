package service

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/repository"
)

// OTPService представляет сервис для работы с OTP-кодами
type OTPService struct {
	otpRepo  *repository.OTPRepository
	userRepo *repository.UserRepository
}

// NewOTPService создает новый сервис для работы с OTP-кодами
func NewOTPService(otpRepo *repository.OTPRepository, userRepo *repository.UserRepository) *OTPService {
	return &OTPService{
		otpRepo:  otpRepo,
		userRepo: userRepo,
	}
}

// Generate генерирует новый OTP-код для пользователя
func (s *OTPService) Generate(req *models.OTPGenerateRequest) (*models.OTPResponse, error) {
	// Проверяем, существует ли пользователь
	_, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Генерируем 6-значный код
	code, err := s.generateOTPCode(6)
	if err != nil {
		return nil, fmt.Errorf("failed to generate OTP code: %w", err)
	}

	// Создаем OTP-код
	otp := &models.OTP{
		UserID:    req.UserID,
		Code:      code,
		ExpiresAt: time.Now().Add(time.Minute * 5),
		Used:      false,
	}

	if err := s.otpRepo.Create(otp); err != nil {
		return nil, fmt.Errorf("failed to create OTP: %w", err)
	}

	// Имитируем отправку SMS
	s.sendSMS(req.UserID, code)

	return &models.OTPResponse{
		Message: "OTP code has been sent to your phone",
	}, nil
}

// Verify проверяет OTP-код
func (s *OTPService) Verify(req *models.OTPVerifyRequest) (*models.OTPResponse, error) {
	// Проверяем, существует ли пользователь
	_, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	code, err := strconv.Atoi(req.Code)
	if err != nil {
		return nil, fmt.Errorf("failed to convert code to int: %w", err)
	}

	// Проверяем OTP-код
	otp, err := s.otpRepo.VerifyOTP(req.UserID, code)
	if err != nil {
		return nil, fmt.Errorf("invalid OTP: %w", err)
	}

	// Помечаем код как использованный
	if err := s.otpRepo.MarkAsUsed(otp.ID); err != nil {
		return nil, fmt.Errorf("failed to mark OTP as used: %w", err)
	}

	return &models.OTPResponse{
		Message: "OTP code verified successfully",
	}, nil
}

// generateOTPCode генерирует случайный 6-значный OTP-код
func (s *OTPService) generateOTPCode(length int) (int, error) {
	var code strings.Builder
	code.Grow(length)

	for i := 0; i < length; i++ {
		code.WriteString(strconv.Itoa(rand.Intn(10)))
	}

	codeInt, err := strconv.Atoi(code.String())
	if err != nil {
		return 0, fmt.Errorf("failed to convert code to int: %w", err)
	}

	return codeInt, nil
}

// sendSMS имитирует отправку SMS с OTP-кодом
func (s *OTPService) sendSMS(userID int64, code int) {
	log.Printf("Sending SMS with OTP code %d to user %d", code, userID)
}
