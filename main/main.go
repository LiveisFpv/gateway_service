package main

import (
	"context"
	"fmt"
	"gateway_service/internal/app"
	"gateway_service/internal/ports/grpc"
	"gateway_service/internal/ports/httpgin"
	"gateway_service/internal/repository"
	"os"
	"strconv"
	"time"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	log "github.com/sirupsen/logrus"
)

// Unpack env and if not seted use standart
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func main() {
	//TODO START GIN WITH GRPC

	//Парсим для подключения к БД
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "0000")
	dbName := getEnv("DB_NAME", "University_DB")

	//Create connection string to db from env
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	//Create logger
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetFormatter(&log.TextFormatter{})

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.WithError(err).Fatalf("can't parse pgxpool config")
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxLogrus.NewLogger(logger),
		LogLevel: tracelog.LogLevelDebug,
	}
	//Create connection pool to db
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		logger.WithError(err).Fatalf("can't create new pool")
	}
	defer pool.Close()

	repo := repository.NewRepository(pool, logger)

	//Create connection to GRPC

	address := getEnv("AUTH_SERVICE_ADDRESS", "localhost:15432")
	timeout := getEnv("AUTH_SERVICE_TIMEOUT", "15s")
	TTL, err := time.ParseDuration(timeout)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse timeout")
	}
	retries := getEnv("AUTH_SERVICE_RETRIES", "3")
	retriesCount, err := strconv.Atoi(retries)
	if err != nil {
		logger.WithError(err).Fatalf("cant parse retries")
	}

	client_auth, err := grpc.New(logger, address, TTL, retriesCount)

	if err != nil {
		logger.WithError(err).Fatalf("cant connect to grpc_auth service")
	}
	usecase := app.NewApp(repo, client_auth)
	port := getEnv("GATEWAY_PORT", "18000")
	port = fmt.Sprintf(":%s", port)
	server := httpgin.NewHTTPServer(port, usecase)
	logger.Infof("Start Gateway Service")
	err = server.Listen()
	if err != nil {
		panic(err)
	}
}
