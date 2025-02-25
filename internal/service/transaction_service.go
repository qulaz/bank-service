package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/repository"
)

// TransactionService представляет сервис для работы с транзакциями
type TransactionService struct {
	txRepo      *repository.TransactionRepository
	accountRepo *repository.AccountRepository
}

// NewTransactionService создает новый сервис для работы с транзакциями
func NewTransactionService(txRepo *repository.TransactionRepository, accountRepo *repository.AccountRepository) *TransactionService {
	return &TransactionService{
		txRepo:      txRepo,
		accountRepo: accountRepo,
	}
}

// Transfer выполняет перевод средств между счетами
func (s *TransactionService) Transfer(req *models.TransferRequest) (*models.Transaction, error) {
	// Получаем счета
	fromAccount, err := s.accountRepo.GetByID(req.FromAccountID)
	if err != nil {
		return nil, fmt.Errorf("from account not found: %w", err)
	}

	toAccount, err := s.accountRepo.GetByID(req.ToAccountID)
	if err != nil {
		return nil, fmt.Errorf("to account not found: %w", err)
	}

	// Проверяем, достаточно ли средств
	if fromAccount.Balance < req.Amount {
		return nil, fmt.Errorf("insufficient funds")
	}

	// Проверяем, совпадают ли валюты
	if fromAccount.Currency != toAccount.Currency {
		return nil, fmt.Errorf("currency mismatch")
	}

	// Создаем транзакцию
	tx := &models.Transaction{
		FromAccountID:   req.FromAccountID,
		ToAccountID:     req.ToAccountID,
		Amount:          req.Amount,
		Type:            models.TransactionTypeTransfer,
		Status:          "completed",
		ReferenceNumber: s.generateReferenceNumber(),
		Description:     req.Description,
	}

	// Обновляем балансы счетов
	if err := s.accountRepo.UpdateBalance(req.FromAccountID, -req.Amount); err != nil {
		return nil, fmt.Errorf("failed to update from account balance: %w", err)
	}

	// Симулируем задержку базы/сети/внешних сервисов
	time.Sleep(time.Duration(rand.Intn(290)+100) * time.Millisecond)

	if err := s.accountRepo.UpdateBalance(req.ToAccountID, req.Amount); err != nil {
		return nil, fmt.Errorf("failed to update to account balance: %w", err)
	}

	// Сохраняем транзакцию
	if err := s.txRepo.Create(tx); err != nil {
		return nil, fmt.Errorf("failed to create transaction: %w", err)
	}

	return tx, nil
}

// GetByID получает транзакцию по ID
func (s *TransactionService) GetByID(id int64) (*models.Transaction, error) {
	return s.txRepo.GetByID(id)
}

// GetByAccountID получает все транзакции счета
func (s *TransactionService) GetByAccountID(accountID int64) ([]*models.Transaction, error) {
	return s.txRepo.GetByAccountID(accountID)
}

// generateReferenceNumber генерирует случайный номер транзакции
func (s *TransactionService) generateReferenceNumber() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 10

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
