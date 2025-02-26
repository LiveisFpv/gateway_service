package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Unpack env and if not seted use standart
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func LoggerSetup(debug bool) *logrus.Logger {
	logger := logrus.New()
	switch debug {
	case true:
		logger.SetLevel(logrus.DebugLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
