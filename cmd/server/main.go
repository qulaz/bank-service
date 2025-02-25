package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/qulaz/bank-service/internal/api"
	"github.com/qulaz/bank-service/internal/api/handlers"
	"github.com/qulaz/bank-service/internal/config"
	customMiddleware "github.com/qulaz/bank-service/internal/middleware"
	"github.com/qulaz/bank-service/internal/repository"
	"github.com/qulaz/bank-service/internal/service"

	_ "github.com/qulaz/bank-service/docs"
)

// @title Bank Service API
// @version 1.0
// @description API для онлайн-банкинга

// @host localhost:8080
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Инициализация конфигурации
	cfg := config.NewConfig()

	// Инициализация базы данных
	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Запуск миграций
	if err := repository.RunMigrations(db.DB, "./migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Инициализация репозиториев
	userRepo := repository.NewUserRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	txRepo := repository.NewTransactionRepository(db)
	otpRepo := repository.NewOTPRepository(db)

	// Инициализация сервисов
	userService := service.NewUserService(userRepo, cfg)
	accountService := service.NewAccountService(accountRepo, userRepo)
	txService := service.NewTransactionService(txRepo, accountRepo)
	otpService := service.NewOTPService(otpRepo, userRepo)

	// Инициализация обработчиков
	userHandler := handlers.NewUserHandler(userService)
	accountHandler := handlers.NewAccountHandler(accountService)
	txHandler := handlers.NewTransactionHandler(txService, accountService)
	otpHandler := handlers.NewOTPHandler(otpService)

	// Инициализация Echo
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Инициализация валидатора
	e.Validator = customMiddleware.NewValidator()

	// Настройка маршрутов
	api.SetupRoutes(e, cfg, userHandler, accountHandler, txHandler, otpHandler)

	// Запуск сервера
	go func() {
		if err := e.Start(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
}
