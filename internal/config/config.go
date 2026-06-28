package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server       ServerConfig
	Database     DatabaseConfig
	JWTConfig    JWTConfig
	UploadConfig UploadConfig
}
type ServerConfig struct {
	Port    string
	GinMode string
}

type DatabaseConfig struct {
	Host     string
	Port     int64
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWTConfig struct {
	Secret              string
	ExpiresIn           time.Duration
	RefreshTokenExpires time.Duration
}

type UploadConfig struct {
	Path        string
	MaxFileSize int64
}

func Load() (*Config, error) {
	_ = godotenv.Load()
	jwtExpiresIn, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))
	refreshTokenExpiresIn, _ := time.ParseDuration(getEnv("REFRESH_TOKEN_EXPIRES_IN", "72h"))
	maxUploadSize, _ := strconv.ParseInt(getEnv("MAX_UPLOAD_SIZE", "10485760"), 10, 64)
	dbPort, _ := strconv.ParseInt(getEnv("DB_PORT", "5432"), 10, 64)
	return &Config{
		Server: ServerConfig{
			Port:    getEnv("SERVER_PORT", "8080"),
			GinMode: getEnv("GIN_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     dbPort,
			User:     getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASSWORD", "A secretpass"),
			Name:     getEnv("DB_NAME", "postgres"),
			SSLMode:  getEnv("DB_SSLMODE", "disable"),
		},
		JWTConfig: JWTConfig{
			Secret:              getEnv("JWT_SECRET", "secret_key"),
			ExpiresIn:           jwtExpiresIn,
			RefreshTokenExpires: refreshTokenExpiresIn,
		},
		UploadConfig: UploadConfig{
			Path:        getEnv("UPLOAD_PATH", "./uploads"),
			MaxFileSize: maxUploadSize,
		},
	}, nil
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
