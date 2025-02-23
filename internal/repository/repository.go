package repository

import (
	"gateway_service/internal/repository/queries"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// DB repo implements for test
type repo struct {
	*queries.Queries
	pool   *pgxpool.Pool
	logger logrus.FieldLogger
}

func NewRepository(pgxPool *pgxpool.Pool, logger logrus.FieldLogger) Repository {
	return &repo{
		Queries: queries.New(pgxPool),
		pool:    pgxPool,
		logger:  logger,
	}
}

// TODO if used DB here
type Repository interface {
	// Auth(ctx context.Context, email, password string) error
	// Register(ctx context.Context, email, password string) error
}
