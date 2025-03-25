package config

import "os"

type Config struct {
	ServerPort  string
	ProfilePort string
}

func Load() *Config {
	return &Config{
		ServerPort:  getEnv("GOMON_SERVER_PORT", "8080"),
		ProfilePort: getEnv("GOMON_PROFILE_PORT", "6060"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
