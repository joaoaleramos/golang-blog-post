package main

import (
	"blog-post/internal/infrastructure/db/postgres"
	server "blog-post/internal/interface"
	"fmt"
)

func main() {
	db, err := postgres.GetDBInstance()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %s", err))
	}

	server := server.NewServer(db)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
