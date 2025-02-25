package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qulaz/bank-service/internal/config"
	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/repository"
	"github.com/qulaz/bank-service/internal/utils"
)

// UserService представляет сервис для работы с пользователями
type UserService struct {
	userRepo *repository.UserRepository
	config   *config.Config
}

// NewUserService создает новый сервис для работы с пользователями
func NewUserService(userRepo *repository.UserRepository, config *config.Config) *UserService {
	return &UserService{
		userRepo: userRepo,
		config:   config,
	}
}

// Register регистрирует нового пользователя
func (s *UserService) Register(req *models.UserRegisterRequest) (*models.User, error) {
	// Проверяем, существует ли пользователь с таким именем
	_, err := s.userRepo.GetByUsername(req.Username)
	if err == nil {
		return nil, fmt.Errorf("user with username %s already exists", req.Username)
	}

	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Создаем пользователя
	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
		FullName: req.FullName,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// Login аутентифицирует пользователя и возвращает JWT токен
func (s *UserService) Login(req *models.UserLoginRequest) (*models.UserResponse, error) {
	// Получаем пользователя по имени
	user, err := s.userRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	// Проверяем пароль
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, fmt.Errorf("invalid username or password")
	}

	// Генерируем JWT токен
	token, err := s.generateJWT(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
		Token:    token,
	}, nil
}

// GetByID получает пользователя по ID
func (s *UserService) GetByID(id int64) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

// generateJWT генерирует JWT токен для пользователя
func (s *UserService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(s.config.JWT.ExpiresIn).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.config.JWT.Secret))
}
