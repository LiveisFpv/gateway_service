package configrpc

import (
	"gateway_service/internal/config"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Authconfig struct {
	Address      string
	Timeout      time.Duration
	RetriesCount int
	Logger       *logrus.Logger
}

// Configure connection to auth service
func AuthSetup(logger *logrus.Logger) *Authconfig {
	address := config.GetEnv("AUTH_SERVICE_ADDRESS", "localhost:15432")
	timeout := config.GetEnv("AUTH_SERVICE_TIMEOUT", "15s")
	TTL, err := time.ParseDuration(timeout)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse timeout")
	}
	retries := config.GetEnv("AUTH_SERVICE_RETRIES", "3")
	retriesCount, err := strconv.Atoi(retries)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse retries")
	}
	authconfig := Authconfig{
		Address:      address,
		Timeout:      TTL,
		RetriesCount: retriesCount,
		Logger:       logger,
	}
	return &authconfig
}
