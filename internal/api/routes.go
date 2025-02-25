package api

import (
	"github.com/labstack/echo/v4"
	"github.com/qulaz/bank-service/internal/api/handlers"
	"github.com/qulaz/bank-service/internal/config"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// SetupRoutes настраивает маршруты API
func SetupRoutes(e *echo.Echo, cfg *config.Config, userHandler *handlers.UserHandler, accountHandler *handlers.AccountHandler, txHandler *handlers.TransactionHandler, otpHandler *handlers.OTPHandler) {
	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// API группа
	api := e.Group("/api")

	// Все маршруты доступны без аутентификации для упрощения тестирования
	// Профиль пользователя
	api.GET("/users/profile", userHandler.GetProfile)

	// Счета
	accounts := api.Group("/accounts")
	accounts.POST("", accountHandler.Create)
	accounts.GET("", accountHandler.GetAll)
	accounts.GET("/:id", accountHandler.GetByID)
	accounts.GET("/:account_id/transactions", txHandler.GetByAccountID)

	// Транзакции
	transactions := api.Group("/transactions")
	transactions.POST("/transfer", txHandler.Transfer)

	// OTP
	otp := api.Group("/otp")
	otp.POST("/generate", otpHandler.Generate)
	otp.POST("/verify", otpHandler.Verify)
}
