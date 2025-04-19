package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	LogLevel string
}

// Load reads from .env and populates Config
func Load() (*Config, error) {
	_ = godotenv.Load() // ignore error if .env does not exist

	cfg := &Config{
		Port:     getEnv("PORT", "8080"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}

	return cfg, nil
}

// getEnv gets an environment variable or returns default
func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists {
		return strings.TrimSpace(val)
	}
	return defaultVal
}

// Validate performs basic checks on config
func (c *Config) Validate() error {
	if c.Port == "" {
		return fmt.Errorf("PORT cannot be empty")
	}
	if c.LogLevel == "" {
		return fmt.Errorf("LOG_LEVEL cannot be empty")
	}
	return nil
}
