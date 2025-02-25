package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/service"
)

// OTPHandler представляет обработчик для работы с OTP-кодами
type OTPHandler struct {
	otpService *service.OTPService
}

// NewOTPHandler создает новый обработчик для работы с OTP-кодами
func NewOTPHandler(otpService *service.OTPService) *OTPHandler {
	return &OTPHandler{
		otpService: otpService,
	}
}

// Generate godoc
// @Summary Генерация OTP-кода
// @Description Генерирует новый OTP-код для тестового пользователя
// @Tags otp
// @Accept json
// @Produce json
// @Success 200 {object} models.OTPResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/otp/generate [post]
func (h *OTPHandler) Generate(c echo.Context) error {
	// Для упрощения всегда используем тестового пользователя с ID 1
	userID := int64(1)

	req := &models.OTPGenerateRequest{
		UserID: userID,
	}

	resp, err := h.otpService.Generate(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

// Verify godoc
// @Summary Проверка OTP-кода
// @Description Проверяет OTP-код тестового пользователя
// @Tags otp
// @Accept json
// @Produce json
// @Param request body models.OTPVerifyRequest true "Данные для проверки OTP-кода"
// @Success 200 {object} models.OTPResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/otp/verify [post]
func (h *OTPHandler) Verify(c echo.Context) error {
	var req models.OTPVerifyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	// Для упрощения всегда используем тестового пользователя с ID 1
	req.UserID = int64(1)

	resp, err := h.otpService.Verify(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
