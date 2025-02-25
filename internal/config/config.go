package config

import (
	"os"
	"time"
)

// Config представляет конфигурацию приложения
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
}

// ServerConfig представляет конфигурацию сервера
type ServerConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// DatabaseConfig представляет конфигурацию базы данных
type DatabaseConfig struct {
	Path string
}

// JWTConfig представляет конфигурацию JWT
type JWTConfig struct {
	Secret    string
	ExpiresIn time.Duration
}

// NewConfig создает новую конфигурацию с значениями по умолчанию
func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:         getEnv("SERVER_PORT", "8080"),
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		},
		Database: DatabaseConfig{
			Path: getEnv("DB_PATH", "./bank.db"),
		},
		JWT: JWTConfig{
			Secret:    getEnv("JWT_SECRET", "super-secret-key"),
			ExpiresIn: time.Hour * 24,
		},
	}
}

// getEnv получает значение переменной окружения или возвращает значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
