package config

import "os"

type Config struct {
	GRPCPort     string
	DATABASE_URL string
	JWTSecret    string
}

func Load() *Config {
	return &Config{
		GRPCPort:     getEnv("GRPC_PORT", "50051"),
		DATABASE_URL: getEnv("DATABASE_URL", ""),
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"), // Change in production
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
