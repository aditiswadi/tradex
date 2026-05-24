package startup

import "tradex/internal/config"

type App struct {
	server *HTTPServer
}

func NewApp() *App {
	cfg := config.Load()
	db := InitDB(cfg)
	server := NewHTTPServer(cfg, db)

	return &App{
		server: server,
	}
}

func (a *App) Run() {
	a.server.Run()
}
