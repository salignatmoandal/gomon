package config

import (
	"os"
	"time"
)

type Config struct {
	ServerPort  string
	ProfilePort string
	MetricsTTL  time.Duration
}

func Load() *Config {
	return &Config{
		ServerPort:  getEnvOrDefault("GOMON_SERVER_PORT", "8080"),
		ProfilePort: getEnvOrDefault("GOMON_PROFILE_PORT", "6060"),
		MetricsTTL:  getDurationOrDefault("GOMON_METRICS_TTL", 10*time.Second),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvOrDefault(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getDurationOrDefault(key string, defaultValue time.Duration) time.Duration {
	if val := os.Getenv(key); val != "" {
		if duration, err := time.ParseDuration(val); err == nil {
			return duration
		}
	}
	return defaultValue
}
