package service

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/repository"
)

// AccountService представляет сервис для работы со счетами
type AccountService struct {
	accountRepo *repository.AccountRepository
	userRepo    *repository.UserRepository
}

// NewAccountService создает новый сервис для работы со счетами
func NewAccountService(accountRepo *repository.AccountRepository, userRepo *repository.UserRepository) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
		userRepo:    userRepo,
	}
}

// Create создает новый счет для пользователя
func (s *AccountService) Create(userID int64, req *models.AccountCreateRequest) (*models.Account, error) {
	// Проверяем, существует ли пользователь
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Генерируем номер счета
	accountNumber := s.generateAccountNumber()

	// Создаем счет
	account := &models.Account{
		UserID:   userID,
		Number:   accountNumber,
		Balance:  0,
		Currency: req.Currency,
	}

	if err := s.accountRepo.Create(account); err != nil {
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	return account, nil
}

// GetByID получает счет по ID
func (s *AccountService) GetByID(id int64) (*models.Account, error) {
	return s.accountRepo.GetByID(id)
}

// GetByUserID получает все счета пользователя
func (s *AccountService) GetByUserID(userID int64) ([]*models.Account, error) {
	return s.accountRepo.GetByUserID(userID)
}

// UpdateBalance обновляет баланс счета
func (s *AccountService) UpdateBalance(id int64, amount int64) error {
	// Получаем счет
	account, err := s.accountRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("account not found: %w", err)
	}

	// Проверяем, достаточно ли средств для списания
	if amount < 0 && account.Balance+amount < 0 {
		return fmt.Errorf("insufficient funds")
	}

	// Обновляем баланс
	return s.accountRepo.UpdateBalance(id, amount)
}

// generateAccountNumber генерирует случайный номер счета
func (s *AccountService) generateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())
	const charset = "0123456789"
	const length = 16

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
