package config

import (
	"os"
)

// Unpack env and if not seted use standart
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
