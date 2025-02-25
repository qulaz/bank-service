package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/qulaz/bank-service/internal/models"
	"github.com/qulaz/bank-service/internal/service"
)

// UserHandler представляет обработчик для работы с пользователями
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler создает новый обработчик для работы с пользователями
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetProfile godoc
// @Summary Получение профиля пользователя
// @Description Возвращает профиль тестового пользователя
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} models.UserResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/users/profile [get]
func (h *UserHandler) GetProfile(c echo.Context) error {
	// Для упрощения всегда возвращаем профиль тестового пользователя с ID 1
	user, err := h.userService.GetByID(1)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		FullName: user.FullName,
	})
}
