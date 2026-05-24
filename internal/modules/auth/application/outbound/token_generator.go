package outbound

import (
	"tradex/internal/modules/auth/domain"
	user "tradex/internal/modules/user/domain"
)

type TokenGenerator interface {
	Generate(user *user.User) (string, error)
	Validate(tokenString string) (*domain.TokenClaims, error)
}
