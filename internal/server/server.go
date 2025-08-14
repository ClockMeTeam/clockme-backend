package server

import (
	"github.com/clockme/clockme-backend/internal/config"
	"github.com/clockme/clockme-backend/internal/db"
	"github.com/clockme/clockme-backend/internal/middleware"
	"github.com/clockme/clockme-backend/internal/users"
	_ "github.com/lib/pq"
	"net/http"
)

type ClockmeServer struct {
	Address     string
	router      *http.ServeMux
	userHandler *users.UserHandler
	cfg         *config.Config
	db          *db.Queries
}

func NewClockmeServer(cfg *config.Config, db *db.Queries) *ClockmeServer {
	userHandler := users.NewUserHandler(db)

	c := &ClockmeServer{
		Address:     ":" + cfg.BackendPort,
		router:      http.NewServeMux(),
		cfg:         cfg,
		db:          db,
		userHandler: userHandler,
	}

	c.routes()
	return c
}
func (c *ClockmeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.router.ServeHTTP(w, r)
}
func (c *ClockmeServer) routes() {

	// Users
	getUsersHandler := http.HandlerFunc(c.userHandler.GetUsersHandler)
	c.router.Handle("GET /users", middleware.EnableCORS(getUsersHandler))
}
