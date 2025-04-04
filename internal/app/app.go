package app

import (
	"context"
	"fmt"
	"gateway_service/internal/ports/grpc/authorization"
	"gateway_service/internal/ports/grpc/country"
	"gateway_service/internal/ports/grpc/fitness"
	"gateway_service/internal/repository"

	sso "github.com/LiveisFPV/sso_v1/gen/go/sso"
)

type App struct {
	repo           repository.Repository
	client_auth    *authorization.Client
	client_country *country.Client
	client_fitness *fitness.Client
}

func NewApp(repo repository.Repository, client_auth *authorization.Client, client_country *country.Client, client_fitness *fitness.Client) *App {
	return &App{
		repo:           repo,
		client_auth:    client_auth,
		client_country: client_country,
		client_fitness: client_fitness,
	}
}

// ? Кто ты воин
func (a *App) Auth(ctx context.Context, email string, password, yandex_token *string) (string, error) {
	req := &sso.LoginRequest{
		Email:       email,
		Password:    password,
		AppId:       1,
		YandexToken: yandex_token,
	}

	resp, err := a.client_auth.Login(ctx, req)
	if err != nil {
		return "", fmt.Errorf("login failed: %w", err)
	}

	return resp.Token, nil
}

func (a *App) Register(ctx context.Context, email string, password, yandex_token *string) error {
	req := &sso.RegisterRequest{
		Email:       email,
		Password:    password,
		YandexToken: yandex_token,
	}

	_, err := a.client_auth.Register(ctx, req)
	if err != nil {
		return fmt.Errorf("registration failed: %w", err)
	}
	return nil
}
