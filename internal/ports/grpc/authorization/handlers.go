package authorization

import (
	"context"

	"github.com/LiveisFPV/sso_v1/gen/go/sso"
)

// Метод для вызова Login у auth-сервиса
func (c *Client) Login(ctx context.Context, req *sso.LoginRequest) (*sso.LoginResponse, error) {
	return c.api.Login(ctx, req)
}

// Метод для вызова Register у auth-сервиса
func (c *Client) Register(ctx context.Context, req *sso.RegisterRequest) (*sso.RegisterResponse, error) {
	return c.api.Register(ctx, req)
}
