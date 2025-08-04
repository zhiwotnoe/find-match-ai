package config

import (
	"log"
	"os"
)

type Config struct {
	GRPCPort    string
	PostgresDSN string
	RedisAddr   string
}

func Load() *Config {
	// Загружаем переменные из .env или окружения
	cfg := &Config{
		GRPCPort:    getEnv("GRPC_PORT", "50051"), // Порт gRPC (значение по умолчанию: 50051)
		PostgresDSN: getEnv("POSTGRES_DSN", ""),   // DSN для Postgres (например: "host=postgres user=postgres dbname=tinder password=123")
		RedisAddr:   getEnv("REDIS_ADDR", "redis:6379"),
	}

	if cfg.PostgresDSN == "" {
		log.Fatal("POSTGRES_DSN не задан в .env")
	}
	return cfg
}

// Вспомогательная функция для чтения переменных окружения
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
