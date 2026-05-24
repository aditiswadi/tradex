package inbound

import (
	"context"
	"tradex/internal/modules/auth/adapter/inbound/rest/dto"
)

type Register interface {
	Register(ctx context.Context, request dto.RegisterRequest) (*dto.RegisterResponse, error)
}
