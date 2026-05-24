package startup

import (
	"tradex/internal/config"
	"tradex/internal/modules/auth"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func registerModules(api *gin.RouterGroup, db *pgxpool.Pool, cfg config.Config) {
	auth.Init(api, db, cfg)
}

func SetupRouter(db *pgxpool.Pool, cfg config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	registerModules(api, db, cfg)

	return r
}
