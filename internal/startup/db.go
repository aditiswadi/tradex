package startup

import (
	"context"
	"log"
	"tradex/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDB(cfg config.Config) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), cfg.DBUrl)
	if err != nil {
		log.Fatal("failed to connectdb", err)
	}

	return db
}
