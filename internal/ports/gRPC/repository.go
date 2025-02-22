package grpc

import (
	"gateway_service/internal/domain"

	"github.com/LiveisFPV/sso_v1/gen/go/sso"
)

// Я пока не поняв как соединить без швов
type GRPC_auth interface {
	Register(client sso.AuthClient) error
	Login(client sso.AuthClient) (*domain.Token, error)
}
