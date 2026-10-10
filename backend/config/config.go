package config

import (
	"fmt"
	"os"
)

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

func Load() (*Config, error) {
	cfg := &Config{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},

		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
	}

	if cfg.Database.Host == "" {
		return nil, fmt.Errorf("DB_HOST is required")
	}

	if cfg.Database.Port == "" {
		return nil, fmt.Errorf("DB_PORT is required")
	}

	if cfg.Database.User == "" {
		return nil, fmt.Errorf("DB_USER is required")
	}

	if cfg.Database.Password == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	if cfg.Database.Name == "" {
		return nil, fmt.Errorf("DB_NAME is required")
	}

	return cfg, nil
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
