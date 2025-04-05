package configrpc

import (
	"gateway_service/internal/config"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Fitnessconfig struct {
	Address      string
	Timeout      time.Duration
	RetriesCount int
	Logger       *logrus.Logger
}

// Configure connection to fitness service
func FitnessSetup(logger *logrus.Logger) *Fitnessconfig {
	address := config.GetEnv("FITNESS_SERVICE_ADDRESS", "localhost:14500")
	timeout := config.GetEnv("FITNESS_SERVICE_TIMEOUT", "15s")
	TTL, err := time.ParseDuration(timeout)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse timeout")
	}
	retries := config.GetEnv("FITNESS_SERVICE_RETRIES", "3")
	retriesCount, err := strconv.Atoi(retries)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse retries")
	}
	fitnessconfig := Fitnessconfig{
		Address:      address,
		Timeout:      TTL,
		RetriesCount: retriesCount,
		Logger:       logger,
	}
	return &fitnessconfig
}
