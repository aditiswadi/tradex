package provider

import (
	"errors"
	"time"
	"tradex/internal/modules/auth/domain"
	user "tradex/internal/modules/user/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenGenerator struct {
	jwtSecret string
}

func NewJWTTokenGenerator(jwtSecret string) *JWTTokenGenerator {
	return &JWTTokenGenerator{jwtSecret: jwtSecret}
}

func (j *JWTTokenGenerator) Generate(user *user.User) (string, error) {
	claims := domain.TokenClaims {
		UserID: user.ID,
		Email: user.Email,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.jwtSecret))
}

func (j *JWTTokenGenerator) Validate(tokenString string) (*domain.TokenClaims, error) {
	claims := &domain.TokenClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("invalid signing method")
			}

			return []byte(j.jwtSecret), nil
		},
	)

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
