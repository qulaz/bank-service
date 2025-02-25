package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/service"
)

// TransactionHandler представляет обработчик для работы с транзакциями
type TransactionHandler struct {
	txService      *service.TransactionService
	accountService *service.AccountService
}

// NewTransactionHandler создает новый обработчик для работы с транзакциями
func NewTransactionHandler(txService *service.TransactionService, accountService *service.AccountService) *TransactionHandler {
	return &TransactionHandler{
		txService:      txService,
		accountService: accountService,
	}
}

// Transfer godoc
// @Summary Перевод средств между счетами
// @Description Выполняет перевод средств между счетами
// @Tags transactions
// @Accept json
// @Produce json
// @Param request body models.TransferRequest true "Данные для перевода"
// @Success 201 {object} models.TransactionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/transactions/transfer [post]
func (h *TransactionHandler) Transfer(c echo.Context) error {
	var req models.TransferRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request body"})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
	}

	// Для упрощения не проверяем принадлежность счета пользователю

	// Выполняем перевод
	tx, err := h.txService.Transfer(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, models.TransactionResponse{
		ID:              tx.ID,
		FromAccountID:   tx.FromAccountID,
		ToAccountID:     tx.ToAccountID,
		Amount:          tx.Amount,
		Type:            tx.Type,
		Status:          tx.Status,
		ReferenceNumber: tx.ReferenceNumber,
		Description:     tx.Description,
		CreatedAt:       tx.CreatedAt,
	})
}

// GetByAccountID godoc
// @Summary Получение транзакций счета
// @Description Возвращает все транзакции счета
// @Tags transactions
// @Accept json
// @Produce json
// @Param account_id path int true "ID счета"
// @Success 200 {object} models.TransactionsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/accounts/{account_id}/transactions [get]
func (h *TransactionHandler) GetByAccountID(c echo.Context) error {
	accountID, err := strconv.ParseInt(c.Param("account_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid account ID"})
	}

	// Для упрощения не проверяем принадлежность счета пользователю

	// Получаем транзакции
	transactions, err := h.txService.GetByAccountID(accountID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	var txResponses []models.TransactionResponse
	for _, tx := range transactions {
		txResponses = append(txResponses, models.TransactionResponse{
			ID:              tx.ID,
			FromAccountID:   tx.FromAccountID,
			ToAccountID:     tx.ToAccountID,
			Amount:          tx.Amount,
			Type:            tx.Type,
			Status:          tx.Status,
			ReferenceNumber: tx.ReferenceNumber,
			Description:     tx.Description,
			CreatedAt:       tx.CreatedAt,
		})
	}

	return c.JSON(http.StatusOK, models.TransactionsResponse{
		Transactions: txResponses,
	})
}
