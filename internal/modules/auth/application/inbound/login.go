package inbound

import (
	"context"
	"tradex/internal/modules/auth/adapter/inbound/rest/dto"
)

type Login interface {
	Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error)
}
