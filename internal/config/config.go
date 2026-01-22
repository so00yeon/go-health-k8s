package config

import "os"

type Config struct {
	Port       string
	AppEnv     string
	AppVersion string
}

func Load() *Config {
	return &Config{
		Port:       getEnv("PORT", "8080"),
		AppEnv:     getEnv("APP_ENV", "development"),
		AppVersion: getEnv("APP_VERSION", "0.1.0"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
