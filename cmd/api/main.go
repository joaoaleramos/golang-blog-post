package main

import (
	"blog-post/internal/infrastructure/db/postgres"
	server "blog-post/internal/interface/api/rest"
	"fmt"
)

func main() {
	db, err := postgres.GetDBInstance()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	server := server.NewServer(db)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
