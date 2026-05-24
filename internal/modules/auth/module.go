package auth

import (
	"tradex/internal/config"
	"tradex/internal/modules/auth/adapter/inbound/rest"
	"tradex/internal/modules/auth/adapter/outbound/persistence"
	"tradex/internal/modules/auth/adapter/outbound/provider"
	"tradex/internal/modules/auth/application/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Init(r *gin.RouterGroup, db *pgxpool.Pool, cfg config.Config) {
	handler := buildHandler(db, cfg)
	rest.RegisterRoutes(r, handler)
}

func buildHandler(db *pgxpool.Pool, cfg config.Config) *rest.Handler {
	repository := persistence.NewPostgresRepository(db)
	passwordHasher := provider.NewBcryptPasswordHasher()
	tokenGenerator := provider.NewJWTTokenGenerator(cfg.JWTSecret)

	registerService := services.NewRegister(repository, passwordHasher)
	loginService := services.NewLogin(repository, passwordHasher, tokenGenerator)

	return rest.NewHandler(registerService, loginService)
}
