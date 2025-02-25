package app

import (
	"context"
	"fmt"
	"gateway_service/internal/ports/grpc/authorization"
	"gateway_service/internal/repository"

	"github.com/LiveisFPV/sso_v1/gen/go/sso"
)

type App struct {
	repo        repository.Repository
	client_auth *authorization.Client
}

func NewApp(repo repository.Repository, client_auth *authorization.Client) *App {
	return &App{
		repo:        repo,
		client_auth: client_auth,
	}
}

func (a *App) Auth(ctx context.Context, email, password string) (string, error) {
	req := &sso.LoginRequest{
		Email:    email,
		Password: password,
		AppId:    1,
	}

	resp, err := a.client_auth.Login(ctx, req)
	if err != nil {
		return "", fmt.Errorf("login failed: %w", err)
	}

	return resp.Token, nil
}

func (a *App) Register(ctx context.Context, email, password string) error {
	req := &sso.RegisterRequest{
		Email:    email,
		Password: password,
	}

	_, err := a.client_auth.Register(ctx, req)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}
	return nil
}
