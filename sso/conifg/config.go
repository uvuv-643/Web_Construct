package config

import (
	"os"
	"strconv"
	"strings"
)

type PostgresConfig struct {
	Username string
	Password string
	Database string
}

type Config struct {
	PG              PostgresConfig
	ApplicationUUID string
}

func New() *Config {
	return &Config{
		PG: PostgresConfig{
			Username: getEnv("PG_USERNAME", ""),
			Password: getEnv("PG_PASSWORD", ""),
			Database: getEnv("PG_DATABASE", ""),
		},
		ApplicationUUID: getEnv("APPLICATION_UUID", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
