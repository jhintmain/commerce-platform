package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	AppEnv       string
	AppPort      string
	MySQLDSN     string
	RedisAddr    string
	RedisPass    string
	RedisDB      string
	KafkaBrokers string
	JWTSecret    string
}

func Load() *Config {
	_ = godotenv.Load(".env.local")

	cfg := &Config{

		AppName:      getEnv("APP_NAME", "commerce-platform"),
		AppEnv:       getEnv("APP_ENV", "local"),
		AppPort:      getEnv("APP_PORT", "8080"),
		RedisAddr:    getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPass:    getEnv("REDIS_PASSWORD", ""),
		RedisDB:      getEnv("REDIS_DB", "0"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "kafka:9092"),
		JWTSecret:    getEnv("JWT_SECRET", "change-me-secret"),
	}

	mysqlHost := getEnv("MYSQL_HOST", "mysql")
	mysqlPort := getEnv("MYSQL_PORT", "3306")
	mysqlDB := getEnv("MYSQL_DATABASE", "commerce")
	mysqlUser := getEnv("MYSQL_USERNAME", "commerce")
	mysqlPass := getEnv("MYSQL_PASSWORD", "commerce")

	cfg.MySQLDSN = mysqlUser + ":" + mysqlPass + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB + "?parseTime=true&charset=utf8mb4&loc=Asia%2FSeoul"

	log.Printf("[config] env=%s port=%s", cfg.AppEnv, cfg.AppPort)

	return cfg
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
