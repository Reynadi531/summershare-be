package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type AppConfig struct {
	App struct {
		PORT      string `json:"port"`
		Name      string `json:"name"`
		JwtSecret string `json:"jwt_secret"`
	}
	DB struct {
		DSN      string `json:"dsn"`
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
		SslMode  string `json:"sslmode"`
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

var Config AppConfig

func InitConfig(DevMode bool) *AppConfig {
	if DevMode {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
			return nil
		}
	}

	Config.App.Name = getEnv("APP_NAME", "")
	Config.App.PORT = getEnv("PORT", "")
	Config.App.JwtSecret = getEnv("JWT_SECRET", "")

	Config.DB.DSN = getEnv("DB_DSN", "")
	Config.DB.Host = getEnv("DB_HOST", "")
	Config.DB.Port = getEnv("DB_PORT", "")
	Config.DB.User = getEnv("DB_USER", "")
	Config.DB.Password = getEnv("DB_PASSWORD", "")
	Config.DB.Name = getEnv("DB_NAME", "")
	Config.DB.SslMode = getEnv("DB_SSL_MODE", "disable")

	return &Config
}
