package config

import (
    "os"
)

// Config holds application-wide configurations
type Config struct {
    ServerPort string
}

// LoadConfig initializes configuration settings
func LoadConfig() Config {
    return Config{
        ServerPort: getEnv("SERVER_PORT", "8080"),
    }
}

// getEnv fetches an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}
