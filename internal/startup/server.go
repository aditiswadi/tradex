package startup

import (
	"tradex/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type HTTPServer struct {
	engine *gin.Engine
	cfg    config.Config
}

func NewHTTPServer(cfg config.Config, db *pgxpool.Pool) *HTTPServer {
	engine := SetupRouter(db, cfg)

	return &HTTPServer{
		engine: engine,
		cfg:    cfg,
	}
}

func (s *HTTPServer) Run() {
	s.engine.Run(":" + s.cfg.ServerPort)
}
