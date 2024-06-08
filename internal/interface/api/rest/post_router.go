package rest

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func PostRouter(apiRouter *mux.Router, db *sqlx.DB) {

	postRouter := apiRouter.PathPrefix("/posts").Subrouter()

	postRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from post!"))
	})

}
