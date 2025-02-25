package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/service"
)

// AccountHandler представляет обработчик для работы со счетами
type AccountHandler struct {
	accountService *service.AccountService
}

// NewAccountHandler создает новый обработчик для работы со счетами
func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
	}
}

// Create godoc
// @Summary Создание нового счета
// @Description Создает новый счет для тестового пользователя
// @Tags accounts
// @Accept json
// @Produce json
// @Param request body models.AccountCreateRequest true "Данные для создания счета"
// @Success 201 {object} models.AccountResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/accounts [post]
func (h *AccountHandler) Create(c echo.Context) error {
	var req models.AccountCreateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	// Для упрощения всегда используем тестового пользователя с ID 1
	userID := int64(1)

	account, err := h.accountService.Create(userID, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, models.AccountResponse{
		ID:       account.ID,
		Number:   account.Number,
		Balance:  account.Balance,
		Currency: account.Currency,
	})
}

// GetAll godoc
// @Summary Получение всех счетов пользователя
// @Description Возвращает все счета тестового пользователя
// @Tags accounts
// @Accept json
// @Produce json
// @Success 200 {object} models.AccountsResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/accounts [get]
func (h *AccountHandler) GetAll(c echo.Context) error {
	// Для упрощения всегда используем тестового пользователя с ID 1
	userID := int64(1)

	accounts, err := h.accountService.GetByUserID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	var accountResponses []models.AccountResponse
	for _, account := range accounts {
		accountResponses = append(accountResponses, models.AccountResponse{
			ID:       account.ID,
			Number:   account.Number,
			Balance:  account.Balance,
			Currency: account.Currency,
		})
	}

	return c.JSON(http.StatusOK, models.AccountsResponse{
		Accounts: accountResponses,
	})
}

// GetByID godoc
// @Summary Получение счета по ID
// @Description Возвращает счет по ID
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path int true "ID счета"
// @Success 200 {object} models.AccountResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/accounts/{id} [get]
func (h *AccountHandler) GetByID(c echo.Context) error {
	accountID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid account ID"})
	}

	account, err := h.accountService.GetByID(accountID)
	if err != nil {
		return c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	}

	// Для упрощения не проверяем принадлежность счета пользователю

	return c.JSON(http.StatusOK, models.AccountResponse{
		ID:       account.ID,
		Number:   account.Number,
		Balance:  account.Balance,
		Currency: account.Currency,
	})
}
