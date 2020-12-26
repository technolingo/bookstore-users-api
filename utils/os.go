package utils

import "os"

// Getenv gets the environment variable or falls back to the default value
func Getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
