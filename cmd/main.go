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
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// @title country API
// @version 2.0
// @description API с документацией Swagger
// @host hackatonsite.jumpingcrab.com:18000
// @BasePath /
func main() {
	// HERE ONLY START GATEWAY
	//Create logger
	logger := config.LoggerSetup(true)

	//Read config
	gincfg := config.GinMustLoad(logger)

	//Create connection pool to db
	pool, err := pgxpool.NewWithConfig(context.Background(), gincfg.Pgxcnf)
	if err != nil {
		logger.WithError(err).Fatalf("can't create new pool")
	}
	defer func() {
		logger.Info("Closing db connection")
		pool.Close()
	}()

	repo := repository.NewRepository(pool, logger)

	//Create connection to auth service
	authcfg := configrpc.AuthSetup(logger)
	client_auth, err := authorization.New(authcfg.Logger, authcfg.Address, authcfg.Timeout, authcfg.RetriesCount)

	if err != nil {
		logger.WithError(err).Fatalf("cant connect to grpc_auth service")
	}

	defer func() {
		logger.Info("Closing gRPC auth client...")
		client_auth.Close()
	}()

	//Create connection to country service
	councfg := configrpc.CountrySetup(logger)
	client_country, err := country.New(councfg.Logger, councfg.Address, councfg.Timeout, councfg.RetriesCount)

	if err != nil {
		logger.WithError(err).Fatalf("cant connect to grpc_country service")
	}

	defer func() {
		logger.Info("Closing gRPC country client...")
		client_country.Close()
	}()

	usecase := app.NewApp(repo, client_auth, client_country)

	//Start service
	server := httpgin.NewHTTPServer(gincfg, usecase)
	logger.Info("Start Gateway Service")

	go func() {
		err = server.Listen()
		if err != nil {
			logger.Fatal(err)
		}
	}()
	//Wait for interrupt signal to shutdown server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	//Wait 30 secconds and then stoped force
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = server.Stop(ctx)
	if err != nil {
		logger.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		logger.Info("Timeout stop server")
	default:
		logger.Info("Server exiting")
	}
}
