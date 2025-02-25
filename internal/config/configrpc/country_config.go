package configrpc

import (
	"gateway_service/internal/config"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Countryconfig struct {
	Address      string
	Timeout      time.Duration
	RetriesCount int
	Logger       *logrus.Logger
}

// Configure connection to country service
func CountrySetup(logger *logrus.Logger) *Countryconfig {
	address := config.GetEnv("COUNTRY_SERVICE_ADDRESS", "localhost:14000")
	timeout := config.GetEnv("COUNTRY_SERVICE_TIMEOUT", "15s")
	TTL, err := time.ParseDuration(timeout)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse timeout")
	}
	retries := config.GetEnv("COUNTRY_SERVICE_RETRIES", "3")
	retriesCount, err := strconv.Atoi(retries)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse retries")
	}
	countryconfig := Countryconfig{
		Address:      address,
		Timeout:      TTL,
		RetriesCount: retriesCount,
		Logger:       logger,
	}
	return &countryconfig
}
