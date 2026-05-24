package outbound

import (
	"context"
	"tradex/internal/modules/user/domain"
)

type Repository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
}
