package app

import (
	"context"
	"gateway_service/internal/domain"
	grpc "gateway_service/internal/ports/gRPC"
	"gateway_service/internal/repository"

	"github.com/LiveisFPV/sso_v1/gen/go/sso"
)

type App struct {
	repo        repository.Repository
	client_auth grpc.GRPC_auth
}

func NewApp(repo repository.Repository, client_auth grpc.GRPC_auth) *App {
	return &App{
		repo:        repo,
		client_auth: client_auth,
	}
}

// Send to Auth-service request
// TODO
func (a *App) Auth(ctx context.Context, email, password string) (*domain.Token, error) {
	req := sso.LoginRequest{
		Email:    email,
		Password: password,
	}
	a.client_auth.Login(req)
}

// Send to Auth-service request
// TODO
func (a *App) Register(ctx context.Context, email, password string) error {
	req := sso.RegisterRequest{
		Email:    email,
		Password: password,
	}
	a.client_auth.Register(req)
	return nil
}
