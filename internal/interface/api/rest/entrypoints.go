package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) RegisterRoutes() http.Handler {
	// Create a new Gorilla Mux router
	mainRouter := mux.NewRouter()

	// Register API routes
	apiRouter := mainRouter.PathPrefix("/api").Subrouter()

	// Register user routes
	s.UserRouter(apiRouter)

	// Register post routes
	s.PostRouter(apiRouter)

	return mainRouter
}
