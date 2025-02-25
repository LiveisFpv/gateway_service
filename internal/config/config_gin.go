package config

import (
	"fmt"

	pgxLogrus "github.com/jackc/pgx-logrus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/sirupsen/logrus"
)

type GinConfig struct {
	ConnStr string
	Logger  *logrus.Logger
	Port    string
	Pgxcnf  *pgxpool.Config
}

func GinMustLoad(logger *logrus.Logger) *GinConfig {
	//Parse enviroment for db
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "5432")
	dbUser := GetEnv("DB_USER", "postgres")
	dbPassword := GetEnv("DB_PASSWORD", "0000")
	dbName := GetEnv("DB_NAME", "gateway_db")
	port := GetEnv("GATEWAY_PORT", "18000")
	port = fmt.Sprintf(":%s", port)

	//Create connection string to db from env
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.WithError(err).Fatalf("can't parse pgxpool config")
	}

	config.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxLogrus.NewLogger(logger),
		LogLevel: tracelog.LogLevelDebug,
	}

	ginconfig := GinConfig{
		ConnStr: connStr,
		Logger:  logger,
		Port:    port,
		Pgxcnf:  config,
	}
	return &ginconfig
}
