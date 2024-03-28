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
	s.UserRoutes(apiRouter) // Modify the existing router instead of creating a new one

	return mainRouter
}
