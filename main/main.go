package main

import (
	"context"
	"gateway_service/internal/app"
	"gateway_service/internal/config"
	"gateway_service/internal/config/configrpc"
	"gateway_service/internal/ports/grpc/authorization"
	"gateway_service/internal/ports/grpc/country"
	"gateway_service/internal/ports/httpgin"
	"gateway_service/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

// HERE ONLY START GATEWAY
func main() {

	//Create logger
	logger := config.LoggerSetup(true)

	//Read config
	gincfg := config.GinMustLoad(logger)

	//Create connection pool to db
	pool, err := pgxpool.NewWithConfig(context.Background(), gincfg.Pgxcnf)
	if err != nil {
		logger.WithError(err).Fatalf("can't create new pool")
	}
	defer pool.Close()

	repo := repository.NewRepository(pool, logger)

	//Create connection to auth service
	authcfg := configrpc.AuthSetup(logger)
	client_auth, err := authorization.New(authcfg.Logger, authcfg.Address, authcfg.Timeout, authcfg.RetriesCount)
	if err != nil {
		logger.WithError(err).Fatalf("cant connect to grpc_auth service")
	}
	//Create connection to country service
	councfg := configrpc.CountrySetup(logger)
	client_country, err := country.New(councfg.Logger, councfg.Address, councfg.Timeout, councfg.RetriesCount)

	if err != nil {
		logger.WithError(err).Fatalf("cant connect to grpc_country service")
	}

	usecase := app.NewApp(repo, client_auth, client_country)

	//Start service
	server := httpgin.NewHTTPServer(gincfg.Port, usecase)
	logger.Info("Start Gateway Service")
	err = server.Listen()
	if err != nil {
		panic(err)
	}
	//TODO! GRASEFUL SHUTDOWN
}
