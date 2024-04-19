package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) PostRouter(apiRouter *mux.Router) {

	postRouter := apiRouter.PathPrefix("/posts").Subrouter()

	postRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from post!"))
	})

}
