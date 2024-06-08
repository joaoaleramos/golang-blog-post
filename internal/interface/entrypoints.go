package server

import (
	"blog-post/internal/interface/api/rest"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func (s *Server) RegisterRoutes(db *sqlx.DB) http.Handler {
	// Create a new Gorilla Mux router
	mainRouter := mux.NewRouter()

	// Register API routes
	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	// Register user routes
	rest.UserRouter(apiRouter, db)

	// Register post routes
	rest.PostRouter(apiRouter, db)

	return mainRouter
}
